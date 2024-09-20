package business

import (
	"context"
	"fx-golang-server/module/core/dto"
	"fx-golang-server/module/core/storage"

	"github.com/rs/zerolog/log"
)

type IElasticBiz interface {
	SearchUsers(ctx context.Context, data dto.SearchUsersRequest) (dto.SearchUsersResponse, error)
}

type elasticBiz struct {
	elasticStorage storage.IElasticStorage
}

func NewElasticBiz(
	elasticStorage storage.IElasticStorage,
) IElasticBiz {
	return &elasticBiz{
		elasticStorage: elasticStorage,
	}
}

func (e *elasticBiz) SearchUsers(ctx context.Context, data dto.SearchUsersRequest) (dto.SearchUsersResponse, error) {
	log.Info().Ctx(ctx).Interface("data", data).Msg("elasticBiz SearchUsers")
	var result dto.SearchUsersResponse
	return result, nil
}
