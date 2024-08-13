package repository

import (
	"context"
	"fx-golang-server/config"
	"fx-golang-server/module/core/dto"
	"fx-golang-server/module/core/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindOne(ctx context.Context, filter dto.FilterUser) (model.UserModel, error)
	Insert(ctx context.Context, data *model.UserModel) error
}

type userRepository struct {
	commonRepository
}

func NewUserRepository(cfg *config.Config, db *gorm.DB) IUserRepository {
	return &userRepository{
		commonRepository: commonRepository{
			db:       db,
			configDb: cfg.Database,
		},
	}
}

func (u *userRepository) tableName() string {
	return model.UserModel{}.TableName()
}

func (s *userRepository) BuildQuery(ctx context.Context, filter dto.FilterUser) *gorm.DB {
	query := s.table(ctx, s.tableName())
	if filter.ID != "" {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.Phone != "" {
		query = query.Where("phone = ?", filter.Phone)
	}
	if filter.Email != "" {
		query = query.Where("email = ?", filter.Email)
	}
	return query
}

func (u *userRepository) FindOne(ctx context.Context, filter dto.FilterUser) (model.UserModel, error) {
	var result model.UserModel
	err := u.CFindOne(ctx, CommonRepositoryParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(ctx, filter),
		Data:         &result,
	})
	return result, err
}

func (u *userRepository) Insert(ctx context.Context, data *model.UserModel) error {
	return u.CInsert(ctx, CommonRepositoryParams{
		TableName: u.tableName(),
		Data:      data,
	})
}
