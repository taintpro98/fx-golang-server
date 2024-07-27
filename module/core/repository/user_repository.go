package repository

import (
	"fx-golang-server/config"

	"gorm.io/gorm"
)

type IUserRepository interface {
}

type userRepository struct {
	commonRepository
}

func NewUserRepository(cfg *config.Config, db *gorm.DB) IUserRepository {
	return userRepository{
		commonRepository: commonRepository{
			db:       db,
			configDb: cfg.Database,
		},
	}
}
