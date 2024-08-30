package token

import (
	"context"
	"crypto/rsa"
	"errors"
	"fx-golang-server/config"
	"fx-golang-server/pkg/e"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
)

type IJWTMaker interface {
	CreateToken(ctx context.Context, data *Payload) (string, error)
	CreateTokenPair(ctx context.Context, data *Payload) (string, string, error)
	VerifyToken(ctx context.Context, tokenString string) (*Payload, error)
}

type jwtMaker struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewJWTMaker(cfg *config.Config) (IJWTMaker, error) {
	cnf := cfg.Token
	privateKeyFile, err := os.ReadFile(cnf.PrivateKeyPath)
	if err != nil {
		log.Error().Err(err).Msg("get private key error")
	}
	priKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyFile)
	if err != nil {
		log.Error().Err(err).Msg("parse private key error")
	}

	publicKeyFile, err := os.ReadFile(cnf.PublicKeyPath)
	if err != nil {
		log.Error().Err(err).Msg("get public key error")
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyFile)
	if err != nil {
		log.Error().Err(err).Msg("parse public key error")
	}
	return &jwtMaker{
		privateKey: priKey,
		publicKey:  pubKey,
	}, nil
}

func (j *jwtMaker) getTokenDuration() (time.Duration, time.Duration) {
	return time.Hour, 24 * time.Hour
}

func (j *jwtMaker) CreateToken(ctx context.Context, payload *Payload) (string, error) {
	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, payload)

	// Sign the token with the RSA private key
	signedToken, err := token.SignedString(j.privateKey)
	if err != nil {
		log.Error().Ctx(ctx).Err(err).Msg("Failed to sign JWT token")
		return "", err
	}
	return signedToken, nil
}

func (j *jwtMaker) CreateTokenPair(ctx context.Context, payload *Payload) (string, string, error) {
	accessDuration, refreshDuration := j.getTokenDuration()

	// Generate access token
	accessPayload := payload
	accessPayload.SetExpires(accessDuration)
	accessPayload.SetRefresh(nil, false)

	jwtAccessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, accessPayload)
	token, err := jwtAccessToken.SignedString(j.privateKey)
	if err != nil {
		return "", "", err
	}

	// Generate refresh token
	refreshPayload := payload
	refreshPayload.SetExpires(refreshDuration)
	refreshPayload.SetRefresh(&accessPayload.Id, true)

	jwtRefreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshPayload)
	refreshToken, err := jwtRefreshToken.SignedString(j.privateKey)
	if err != nil {
		return "", "", err
	}
	return token, refreshToken, nil
}

func (j *jwtMaker) VerifyToken(ctx context.Context, tokenString string) (*Payload, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		return j.publicKey, nil
	})

	if err != nil {
		log.Error().Ctx(ctx).Err(err).Msg("Failed to parse JWT token:")
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			return nil, ve.Inner
		}
		return nil, err
	}

	// Validate the token
	if claims, ok := token.Claims.(*Payload); ok && token.Valid {
		return claims, nil
	} else {
		log.Error().Ctx(ctx).Msg("Token is invalid")
		return nil, e.ErrUnauthorized
	}
}
