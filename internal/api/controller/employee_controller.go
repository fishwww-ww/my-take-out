package controller

import (
	"github.com/gin-gonic/gin"
	"my-take-out/common/retcode"
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
		ctx.JSON(400, gin.H{
			"message": "失败1",
		})
		return
	}
	resp, err := ec.service.Login(ctx, employeeLogin)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "失败2",
			"error":   err.Error(),
		})
		return
	}
	retcode.OK(ctx, resp)
}

func (ec *EmployeeController) AddEmployee(ctx *gin.Context) {
	employee := request.EmployeeDTO{}
	err := ctx.Bind(&employee)
	if err != nil {
		retcode.Fatal(ctx, err, "失败1")
		return
	}
	err = ec.service.CreateEmployee(ctx, employee)
	if err != nil {
		retcode.Fatal(ctx, err, "失败2")
		return
	}
	retcode.OK(ctx, "")
}
