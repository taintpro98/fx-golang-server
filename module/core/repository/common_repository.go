package repository

import (
	"context"
	"errors"
	"fmt"
	"fx-golang-server/config"
	"fx-golang-server/module/core/dto"

	"github.com/rs/zerolog/log"

	"gorm.io/gorm"
)

type CommonRepositoryParams struct {
	TableName    string
	Query        *gorm.DB
	CommonFilter dto.CommonFilter
	Filter       interface{}
	Data         interface{}
}

type commonRepository struct {
	db       *gorm.DB
	configDb config.DatabaseConfig
}

func (s commonRepository) table(ctx context.Context, tableName string) *gorm.DB {
	return s.db.Table(fmt.Sprintf("%s.%s", s.configDb.Schema, tableName)).WithContext(ctx)
}

func (s commonRepository) log(ctx context.Context, funcName string, param CommonRepositoryParams) {
	log.Info().
		Ctx(ctx).
		Interface("data", param.Data).
		Interface("filter", param.Filter).
		Interface("common_filter", param.CommonFilter).
		Msg(fmt.Sprintf("%s %s table", funcName, param.TableName))
}

func (s commonRepository) CCount(ctx context.Context, param CommonRepositoryParams) (*int64, error) {
	s.log(ctx, "CCount", param)
	var count int64
	tx := param.Query.Count(&count)
	if tx.Error != nil {
		log.Error().Ctx(ctx).Stack().Err(tx.Error).Msg(fmt.Sprintf("count %s error", param.TableName))
		return nil, tx.Error
	}
	return &count, nil
}

func (s commonRepository) CFindOne(ctx context.Context, param CommonRepositoryParams) error {
	s.log(ctx, "CFindOne", param)
	if len(param.CommonFilter.Select) > 0 {
		param.Query = param.Query.Select(param.CommonFilter.Select)
	}
	if len(param.CommonFilter.Preloads) > 0 {
		for _, item := range param.CommonFilter.Preloads {
			param.Query = param.Query.Preload(item)
		}
	}

	tx := param.Query.First(param.Data)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		log.Error().Ctx(ctx).Stack().Err(tx.Error).Msg(fmt.Sprintf("find one %s error", param.TableName))
		return tx.Error
	}
	return nil
}

func (s commonRepository) CList(ctx context.Context, param CommonRepositoryParams) error {
	s.log(ctx, "CList", param)
	if param.CommonFilter.Limit != 0 {
		param.Query = param.Query.Limit(param.CommonFilter.Limit)
	}
	if param.CommonFilter.Offset != nil {
		param.Query = param.Query.Offset(*param.CommonFilter.Offset)
	}
	if param.CommonFilter.Sort != "" {
		param.Query = param.Query.Order(param.CommonFilter.Sort) // age desc hoac age asc hoac age
	}
	if len(param.CommonFilter.Select) > 0 {
		param.Query = param.Query.Select(param.CommonFilter.Select)
	}
	if len(param.CommonFilter.Preloads) > 0 { // khong khuyen khich dung preloads trong list
		for _, item := range param.CommonFilter.Preloads {
			param.Query = param.Query.Preload(item)
		}
	}

	tx := param.Query.Find(param.Data) // day la con tro vao bien ket qua
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		log.Error().Ctx(ctx).Stack().Err(tx.Error).Msg(fmt.Sprintf("list %s error", param.TableName))
		return tx.Error
	}
	return nil
}

func (s commonRepository) CInsert(ctx context.Context, param CommonRepositoryParams) error {
	s.log(ctx, "CInsert", param)

	tx := s.table(ctx, param.TableName).Create(param.Data)
	if tx.Error != nil {
		log.Error().Ctx(ctx).Stack().Err(tx.Error).Msg(fmt.Sprintf("insert %s data error", param.TableName))
	}
	return tx.Error
}

func (s commonRepository) CInsertBatch(ctx context.Context, param CommonRepositoryParams) error {
	s.log(ctx, "CInsertBatch", param)
	// Create a transaction
	tx := s.table(ctx, param.TableName).Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Info().Ctx(ctx).Msg(fmt.Sprintf("rollback CInsertBatch %s", param.TableName))
			tx.Rollback()
		}
	}()
	err := tx.Create(param.Data).Error // param.Data phai la con tro
	if err != nil {
		log.Error().Ctx(ctx).Err(err).Msg(fmt.Sprintf("CInsertBatch %s error", param.TableName))
		return err
	}
	err = tx.Commit().Error
	if err != nil {
		log.Error().Ctx(ctx).Err(err).Msg(fmt.Sprintf("CInsertBatch %s Error when commit transaction", param.TableName))
		return err
	}
	return nil
}

func (s commonRepository) CUpdateMany(ctx context.Context, param CommonRepositoryParams) error {
	s.log(ctx, "CUpdateMany", param)

	tx := param.Query.Updates(param.Data)
	if tx.Error != nil {
		log.Error().Ctx(ctx).Err(tx.Error).Msg(fmt.Sprintf("Failed to update many %s", param.TableName))
	}
	return tx.Error
}

func (s commonRepository) CDelete(ctx context.Context, param CommonRepositoryParams) error {
	s.log(ctx, "CDelete", param)
	tx := param.Query.Delete(param.Data)
	if tx.Error != nil {
		log.Error().Ctx(ctx).Err(tx.Error).Msg(fmt.Sprintf("Failed to delete %s", param.TableName))
	}
	return tx.Error
}
