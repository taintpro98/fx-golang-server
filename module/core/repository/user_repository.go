package repository

import (
	"context"
	"fx-golang-server/config"
	"fx-golang-server/module/core/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Insert(ctx context.Context, data *model.UserModel) error
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

func (u userRepository) tableName() string {
	return model.UserModel{}.TableName()
}

func (u userRepository) Insert(ctx context.Context, data *model.UserModel) error {
	return u.CInsert(ctx, CommonRepositoryParams{
		TableName: u.tableName(),
		Data:      data,
	})
}
