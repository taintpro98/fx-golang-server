package business

import (
	"context"
	"fx-golang-server/module/core/dto"
	"fx-golang-server/module/core/model"
	"fx-golang-server/module/core/repository"
)

type ICustomerBiz interface {
	GetCustomerProfile(ctx context.Context, userID string) (model.UserModel, error)
}

type customerBiz struct {
	userRepo repository.IUserRepository
}

func NewCustomerBiz(
	userRepo repository.IUserRepository,
) ICustomerBiz {
	return &customerBiz{
		userRepo: userRepo,
	}
}

func (v *customerBiz) GetCustomerProfile(ctx context.Context, userID string) (model.UserModel, error) {
	userDB, err := v.userRepo.FindOne(ctx, dto.FilterUser{
		ID: userID,
	})
	if err != nil {
		return model.UserModel{}, err
	}
	return userDB, nil
}
