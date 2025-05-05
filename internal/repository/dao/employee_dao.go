package dao

import (
	"context"
	"gorm.io/gorm"
	"my-take-out/common/e"
	"my-take-out/common/retcode"
	"my-take-out/internal/model"
)

type EmployeeDao struct {
	db *gorm.DB
}

func NewEmployeeDao(db *gorm.DB) *EmployeeDao {
	return &EmployeeDao{db: db}
}

func (d *EmployeeDao) GetByUserName(ctx context.Context, username string) (*model.Employee, error) {
	var employee model.Employee
	err := d.db.WithContext(ctx).Where("username=?", username).First(&employee).Error
	if err != nil {
		return nil, retcode.NewError(e.MysqlERR, "Get employee failed")
	}
	return &employee, err
}

func (d *EmployeeDao) Insert(ctx context.Context, entity model.Employee) error {
	err := d.db.WithContext(ctx).Create(&entity).Error
	if err != nil {
		return retcode.NewError(e.MysqlERR, "Insert employee failed")
	}
	return nil
}
