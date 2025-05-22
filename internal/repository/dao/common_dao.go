package dao

import (
	"context"
	"gorm.io/gorm"
	"my-take-out/common/e"
	"my-take-out/common/retcode"
	"my-take-out/internal/model"
)

type CommonDao struct {
	db *gorm.DB
}

func NewCommonDao(db *gorm.DB) *CommonDao {
	return &CommonDao{db: db}
}

func (d *CommonDao) Insert(ctx context.Context, entity model.File) error {
	err := d.db.WithContext(ctx).Create(&entity).Error
	if err != nil {
		return retcode.NewError(e.MysqlERR, "Insert file failed")
	}
	return nil
}
