package token

import (
	"context"
	"crypto/rsa"
	"fx-golang-server/config"
	"fx-golang-server/module/core/dto"
	"fx-golang-server/pkg/e"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
)

type IJWTMaker interface {
	CreateToken(ctx context.Context, data interface{}) (string, error)
	VerifyToken(ctx context.Context, token string) (interface{}, error)
}

type jwtMaker struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewJWTMaker(cfg *config.Config) (IJWTMaker, error) {
	cnf := cfg.Token
	privateKeyFile, err := os.ReadFile(cnf.PrivateKeyPath)
	if err != nil {
		log.Fatal().Err(err).Msg("get private key error")
	}
	priKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyFile)
	if err != nil {
		log.Fatal().Err(err).Msg("parse private key error")
	}

	publicKeyFile, err := os.ReadFile(cnf.PublicKeyPath)
	if err != nil {
		log.Fatal().Err(err).Msg("get public key error")
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyFile)
	if err != nil {
		log.Fatal().Err(err).Msg("parse public key error")
	}
	return jwtMaker{
		privateKey: priKey,
		publicKey:  pubKey,
	}, nil
}

// CreateToken implements IJWTMaker.
func (j jwtMaker) CreateToken(ctx context.Context, payload interface{}) (string, error) {
	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, payload.(jwt.Claims))

	// Sign the token with the RSA private key
	signedToken, err := token.SignedString(j.privateKey)
	if err != nil {
		log.Error().Ctx(ctx).Err(err).Msg("Failed to sign JWT token")
		return "", err
	}
	return signedToken, nil
}

// VerifyToken implements IJWTMaker.
func (j jwtMaker) VerifyToken(ctx context.Context, tokenString string) (interface{}, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &dto.UserPayload{}, func(token *jwt.Token) (interface{}, error) {
		return j.publicKey, nil
	})

	if err != nil {
		log.Fatal().Ctx(ctx).Err(err).Msg("Failed to parse JWT token:")
		return nil, err
	}

	// Validate the token
	if claims, ok := token.Claims.(*dto.UserPayload); ok && token.Valid {
		return claims, nil
	} else {
		log.Info().Ctx(ctx).Msg("Token is invalid")
		return nil, e.ErrUnauthorized
	}
}
