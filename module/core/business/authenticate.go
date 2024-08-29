package business

import (
	"context"
	"fx-golang-server/module/core/dto"
	"fx-golang-server/module/core/model"
	"fx-golang-server/module/core/repository"
	"fx-golang-server/pkg/e"
	"fx-golang-server/token"
	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
	"time"
)

type IAuthenticateBiz interface {
	Register(ctx context.Context, data dto.CreateUserRequest) (dto.CreateUserResponse, error)
	Login(ctx context.Context, data dto.LoginRequest) (dto.CreateUserResponse, error)
}

type authenticateBiz struct {
	jwtMaker token.IJWTMaker
	userRepo repository.IUserRepository
}

func NewAuthenticateBiz(
	jwtMaker token.IJWTMaker,
	userRepo repository.IUserRepository,
) IAuthenticateBiz {
	return &authenticateBiz{
		jwtMaker: jwtMaker,
		userRepo: userRepo,
	}
}

func (t *authenticateBiz) Register(ctx context.Context, data dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	log.Info().Ctx(ctx).Interface("data", data).Msg("authenticateBiz Register")
	var response dto.CreateUserResponse
	dataInsert := model.UserModel{
		Phone: &data.Phone,
		Email: &data.Email,
	}
	err := t.userRepo.Insert(ctx, &dataInsert)
	if err != nil {
		return response, err
	}

	tokenString, err := t.jwtMaker.CreateToken(ctx, dto.UserPayload{
		StandardClaims: jwt.StandardClaims{
			Subject:   dataInsert.ID,
			ExpiresAt: 0,
		},
	})
	if err != nil {
		log.Error().Ctx(ctx).Err(err).Msg("create token error")
		return response, err
	}
	response.Token = tokenString
	return response, nil
}

func (b *authenticateBiz) Login(ctx context.Context, data dto.LoginRequest) (dto.CreateUserResponse, error) {
	var response dto.CreateUserResponse
	user, err := b.userRepo.FindOne(ctx, dto.FilterUser{
		Phone: data.Phone,
	})
	if err != nil {
		return response, err
	}
	if user.ID == "" {
		return response, e.ErrDataNotFound("user")
	}
	tokenString, err := b.jwtMaker.CreateToken(ctx, dto.UserPayload{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.ID,
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	})
	if err != nil {
		log.Error().Ctx(ctx).Err(err).Msg("create token error")
		return response, err
	}
	response.Token = tokenString
	return response, nil
}
