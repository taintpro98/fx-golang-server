package repository

import (
	"context"
	"fx-golang-server/config"
	"fx-golang-server/module/core/dto"
	"fx-golang-server/module/core/model"

	"gorm.io/gorm"
)

type IMovieRepository interface {
	Count(ctx context.Context, filter dto.FilterMovie) (*int64, error)

	FindOne(ctx context.Context, filter dto.FilterMovie) (model.MovieModel, error)

	List(ctx context.Context, filter dto.FilterMovie) ([]model.MovieModel, error)

	Insert(ctx context.Context, data *model.MovieModel) error
}

type movieRepository struct {
	commonRepository
}

func NewMovieRepository(cfg *config.Config, db *gorm.DB) IMovieRepository {
	return &movieRepository{
		commonRepository: commonRepository{
			db:       db,
			configDb: cfg.Database,
		},
	}
}

func (u *movieRepository) tableName() string {
	return model.MovieModel{}.TableName()
}

func (s *movieRepository) BuildQuery(ctx context.Context, filter dto.FilterMovie) *gorm.DB {
	query := s.table(ctx, s.tableName())
	if filter.ID != "" {
		query = query.Where("id = ?", filter.ID)
	}
	return query
}

func (u *movieRepository) FindOne(ctx context.Context, filter dto.FilterMovie) (model.MovieModel, error) {
	var result model.MovieModel
	err := u.CFindOne(ctx, CommonRepositoryParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(ctx, filter),
		Data:         &result,
	})
	return result, err
}

func (u *movieRepository) Count(ctx context.Context, filter dto.FilterMovie) (*int64, error) {
	return u.CCount(ctx, CommonRepositoryParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(ctx, filter),
	})
}

// List implements IPostStorage.
func (u *movieRepository) List(ctx context.Context, filter dto.FilterMovie) ([]model.MovieModel, error) {
	var result []model.MovieModel // khoi tao cho nay ra mang rong
	err := u.CList(ctx, CommonRepositoryParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(ctx, filter),
		Data:         &result,
	})
	return result, err
}

func (u *movieRepository) Insert(ctx context.Context, data *model.MovieModel) error {
	return u.CInsert(ctx, CommonRepositoryParams{
		TableName: u.tableName(),
		Data:      data,
	})
}

func (u *movieRepository) UpdateMany(ctx context.Context, filter dto.FilterMovie, data model.MovieModel) error {
	return u.CUpdateMany(ctx, CommonRepositoryParams{
		TableName: u.tableName(),
		Data:      data,
		Query:     u.BuildQuery(ctx, filter),
	})
}
