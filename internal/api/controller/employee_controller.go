package controller

import (
	"github.com/gin-gonic/gin"
	"my-take-out/common/retcode"
	"my-take-out/global"
	"my-take-out/internal/api/request"
	"my-take-out/internal/service"
)

type EmployeeController struct {
	service service.IEmployeeService
}

func NewEmployeeController(employeeService service.IEmployeeService) *EmployeeController {
	return &EmployeeController{service: employeeService}
}

func (ec *EmployeeController) Login(ctx *gin.Context) {
	employeeLogin := request.EmployeeLogin{}
	err := ctx.Bind(&employeeLogin)
	if err != nil {
		global.Log.Debug("EmployeeController login 解析失败")
		retcode.Fatal(ctx, err, "")
		return
	}
	resp, err := ec.service.Login(ctx, employeeLogin)
	if err != nil {
		global.Log.Warn("EmployeeController login Error:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, resp)
}

func (ec *EmployeeController) AddEmployee(ctx *gin.Context) {
	employee := request.EmployeeDTO{}
	err := ctx.Bind(&employee)
	if err != nil {
		global.Log.Debug("AddEmployee Error:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	err = ec.service.CreateEmployee(ctx, employee)
	if err != nil {
		global.Log.Warn("AddEmployee Error", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, "")
}

func (ec *EmployeeController) PageQuery(ctx *gin.Context) {
	var employeePageQueryDTO request.EmployeePageQueryDTO
	err := ctx.Bind(&employeePageQueryDTO)
	if err != nil {
		global.Log.Error("AddEmployee  invalid params err:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	// 进行分页查询
	_, err = ec.service.PageQuery(ctx, employeePageQueryDTO)
	if err != nil {
		global.Log.Warn("AddEmployee  Error:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
}
