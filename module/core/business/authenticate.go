package business

import (
	"context"
	"fx-golang-server/module/core/dto"
	"fx-golang-server/module/core/model"
	"fx-golang-server/module/core/repository"
	"fx-golang-server/token"

	"github.com/rs/zerolog/log"
)

type IAuthenticateBiz interface {
	Register(ctx context.Context, data dto.CreateUserRequest) (dto.CreateUserResponse, error)
}

type authenticateBiz struct {
	jwtMaker token.IJWTMaker
	userRepo repository.IUserRepository
}

func NewAuthenticateBiz(
	jwtMaker token.IJWTMaker,
	userRepo repository.IUserRepository,
) IAuthenticateBiz {
	return authenticateBiz{
		jwtMaker: jwtMaker,
		userRepo: userRepo,
	}
}

func (t authenticateBiz) Register(ctx context.Context, data dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	log.Info().Ctx(ctx).Msg("authenticateBiz Register")
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
		UserID: dataInsert.ID,
	})
	if err != nil {
		log.Error().Ctx(ctx).Err(err).Msg("create token error")
		return response, err
	}
	response.Token = tokenString
	return response, nil
}
