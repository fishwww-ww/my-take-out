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

func (d *CommonDao) Query(ctx context.Context, fileName string) (string, error) {
	// 查询数据库中filename字段为fileName的记录
	var file model.File
	err := d.db.WithContext(ctx).Where("name = ?", fileName).First(&file).Error
	if err != nil {
		return "", retcode.NewError(e.MysqlERR, "Query file failed")
	}
	return file.Uuid, nil
}
