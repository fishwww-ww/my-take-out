package service

import (
	"context"
	"my-take-out/common"
	"my-take-out/common/e"
	"my-take-out/common/enum"
	"my-take-out/common/utils"
	"my-take-out/global"
	"my-take-out/internal/api/request"
	"my-take-out/internal/api/response"
	"my-take-out/internal/model"
	"my-take-out/internal/repository/dao"
	"time"
)

type IEmployeeService interface {
	Login(context.Context, request.EmployeeLogin) (*response.EmployeeLogin, error)
	CreateEmployee(context.Context, request.EmployeeDTO) error
	PageQuery(ctx context.Context, query request.EmployeePageQueryDTO) (*common.PageResult, error)
}

type EmployeeImpl struct {
	repo *dao.EmployeeDao
}

func (ei *EmployeeImpl) Login(ctx context.Context, employeeLogin request.EmployeeLogin) (*response.EmployeeLogin, error) {
	employee, err := ei.repo.GetByUserName(ctx, employeeLogin.UserName)
	if err != nil || employee == nil {
		return nil, e.Error_ACCOUNT_NOT_FOUND
	}

	password := utils.MD5V(employeeLogin.Password, "", 0)
	if password != employee.Password {
		return nil, e.Error_PASSWORD_ERROR
	}

	if employee.Status == enum.DISABLE {
		return nil, e.Error_ACCOUNT_LOCKED
	}

	jwtConfig := global.Config.Jwt.Admin
	token, err := utils.GenerateToken(employee.Id, jwtConfig.Name, jwtConfig.Secret)
	if err != nil {
		return nil, err
	}
	resp := response.EmployeeLogin{
		Id:       employee.Id,
		Name:     employee.Name,
		Token:    token,
		UserName: employee.Username,
	}
	return &resp, nil
}

func (ei *EmployeeImpl) CreateEmployee(ctx context.Context, employeeDTO request.EmployeeDTO) error {
	//// 校验用户名是否存在
	//employee, err := ei.repo.GetByUserName(ctx, employeeDTO.UserName)
	//if err != nil {
	//	return err
	//}
	//if employee != nil {
	//	return e.Error_ACCOUNT_EXIST
	//}
	entity := model.Employee{
		Id:         employeeDTO.Id,
		Name:       employeeDTO.Name,
		Username:   employeeDTO.UserName,
		Password:   utils.MD5V("123456", "", 0),
		Phone:      employeeDTO.Phone,
		Sex:        employeeDTO.Sex,
		IdNumber:   employeeDTO.IdNumber,
		Status:     enum.ENABLE,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := ei.repo.Insert(ctx, entity)
	return err
}

func (ei *EmployeeImpl) PageQuery(ctx context.Context, dto request.EmployeePageQueryDTO) (*common.PageResult, error) {
	pageResult, err := ei.repo.PageQuery(ctx, dto)
	// 防止信息泄露
	//if employees, ok := pageResult.Records.([]model.Employee); ok {
	//	for key, _ := range employees {
	//		employees[key].Password = "****"
	//		employees[key].IdNumber = "****"
	//		employees[key].Phone = "****"
	//	}
	//	pageResult.Records = employees
	//}
	return pageResult, err
}

// 为什么返回类型是接口？？？
// EmployeeImpl 是一个结构体类型，它实现了 IEmployeeService 所有接口
func NewEmployeeService(repo *dao.EmployeeDao) IEmployeeService {
	return &EmployeeImpl{repo: repo}
}
