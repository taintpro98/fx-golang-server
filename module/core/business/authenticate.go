package business

import "fx-golang-server/module/core/repository"

type IAuthenticateBiz interface {
}

type authenticateBiz struct {
	userRepo repository.IUserRepository
}

func NewAuthenticateBiz(
	userRepo repository.IUserRepository,
) IAuthenticateBiz {
	return authenticateBiz{
		userRepo: userRepo,
	}
}
