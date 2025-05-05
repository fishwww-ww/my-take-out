package admin

import (
	"github.com/gin-gonic/gin"
	"my-take-out/global"
	"my-take-out/internal/api/controller"
	"my-take-out/internal/repository/dao"
	"my-take-out/internal/service"
)

type EmployeeRouter struct {
	service service.IEmployeeService
}

func (er *EmployeeRouter) InitApiRouter(router *gin.RouterGroup) {
	publicRouter := router.Group("employee")
	er.service = service.NewEmployeeService(
		dao.NewEmployeeDao(global.DB),
	)
	employeeCtl := controller.NewEmployeeController(er.service)
	{
		publicRouter.POST("/login", employeeCtl.Login)
		publicRouter.POST("", employeeCtl.AddEmployee)
	}
}
