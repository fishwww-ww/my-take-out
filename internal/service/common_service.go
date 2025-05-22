package service

import (
	"context"
	"my-take-out/internal/model"
	"my-take-out/internal/repository/dao"
)

type ICommonService interface {
	Insert(ctx context.Context, entity model.File) error
	QueryUuid(ctx context.Context, fileName string) (string, error)
}

type CommonImpl struct {
	repo *dao.CommonDao
}

func (ci *CommonImpl) Insert(ctx context.Context, entity model.File) error {
	return ci.repo.Insert(ctx, entity)
}

func (ci *CommonImpl) QueryUuid(ctx context.Context, fileName string) (string, error) {
	return ci.repo.Query(ctx, fileName)
}

func NewCommonService(repo *dao.CommonDao) ICommonService {
	return &CommonImpl{repo: repo}
}
