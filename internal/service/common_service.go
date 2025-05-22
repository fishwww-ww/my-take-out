package service

import (
	"context"
	"my-take-out/internal/model"
	"my-take-out/internal/repository/dao"
)

type ICommonService interface {
	Insert(ctx context.Context, entity model.File) error
}

type CommonImpl struct {
	repo *dao.CommonDao
}

func (ci *CommonImpl) Insert(ctx context.Context, entity model.File) error {
	return ci.repo.Insert(ctx, entity)
}

func NewCommonService(repo *dao.CommonDao) ICommonService {
	return &CommonImpl{repo: repo}
}
