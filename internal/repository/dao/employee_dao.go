package dao

import (
	"context"
	"gorm.io/gorm"
	"my-take-out/common"
	"my-take-out/common/e"
	"my-take-out/common/retcode"
	"my-take-out/internal/api/request"
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

func (d *EmployeeDao) PageQuery(ctx context.Context, dto request.EmployeePageQueryDTO) (*common.PageResult, error) {
	var result common.PageResult      // 存储查询结果
	var employeeList []model.Employee // 存储员工记录
	var err error
	query := d.db.WithContext(ctx).Model(&model.Employee{})
	if dto.Name != "" {
		// 模糊查询
		query = query.Where("name LIKE ?", "%"+dto.Name+"%")
	}
	// 查询记录总数
	if err = query.Count(&result.Total).Error; err != nil {
		return nil, retcode.NewError(e.MysqlERR, "Get employee count failed")
	}
	err = query.Scopes(result.Paginate(&dto.Page, &dto.PageSize)).Find(&employeeList).Error
	if err != nil {
		return nil, retcode.NewError(e.MysqlERR, "Get employee list failed")
	}
	result.Records = employeeList
	return &result, nil
}
