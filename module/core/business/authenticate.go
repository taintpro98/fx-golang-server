package business

import (
	"context"
	"fx-golang-server/module/core/dto"
	"fx-golang-server/module/core/model"
	"fx-golang-server/module/core/repository"
	"fx-golang-server/pkg/e"
	"fx-golang-server/token"
	"time"

	"github.com/rs/zerolog/log"
)

type IAuthenticateBiz interface {
	Register(ctx context.Context, data dto.CreateUserRequest) (dto.CreateUserResponse, error)
	Login(ctx context.Context, data dto.LoginRequest) (dto.CreateUserResponse, error)
	Refresh(ctx context.Context, data dto.RefreshRequest) (dto.CreateUserResponse, error)
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

	payload, err := token.NewPayload(dataInsert.ID, time.Hour, map[string]interface{}{
		"email": dataInsert.Email,
	})
	if err != nil {
		log.Error().Ctx(ctx).Err(err).Msg("new payload error")
		return response, err
	}
	tokenString, refreshToken, err := t.jwtMaker.CreateTokenPair(ctx, payload)
	if err != nil {
		log.Error().Ctx(ctx).Err(err).Msg("create token error")
		return response, err
	}
	response.Token = tokenString
	response.RefreshToken = refreshToken
	return response, nil
}

func (b *authenticateBiz) Login(ctx context.Context, data dto.LoginRequest) (dto.CreateUserResponse, error) {
	log.Info().Ctx(ctx).Interface("data", data).Msg("authenticateBiz Login")
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
	payload, err := token.NewPayload(user.ID, time.Hour, map[string]interface{}{
		"email": user.Email,
	})
	if err != nil {
		log.Error().Ctx(ctx).Err(err).Msg("new payload error")
		return response, err
	}
	accessToken, refreshToken, err := b.jwtMaker.CreateTokenPair(ctx, payload)
	if err != nil {
		log.Error().Ctx(ctx).Err(err).Msg("create token error")
		return response, err
	}
	response.Token = accessToken
	response.RefreshToken = refreshToken
	return response, nil
}

func (t *authenticateBiz) Refresh(ctx context.Context, data dto.RefreshRequest) (dto.CreateUserResponse, error) {
	log.Info().Ctx(ctx).Interface("data", data).Msg("authenticateBiz Refresh")
	var response dto.CreateUserResponse
	payload, err := t.jwtMaker.VerifyToken(ctx, data.RefreshToken)
	if err != nil {
		return response, err
	}
	if !payload.Refresh {
		return response, e.ErrInvalidRefreshToken
	}
	accessToken, refreshToken, err := t.jwtMaker.CreateTokenPair(ctx, payload)
	if err != nil {
		log.Error().Ctx(ctx).Err(err).Msg("create token error")
		return response, err
	}
	response.Token = accessToken
	response.RefreshToken = refreshToken
	return response, nil
}
