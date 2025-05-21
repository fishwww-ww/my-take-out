package admin

import (
	"github.com/gin-gonic/gin"
	"my-take-out/global"
	"my-take-out/internal/api/controller"
	"my-take-out/internal/repository/dao"
	"my-take-out/internal/service"
	"my-take-out/middle"
)

type EmployeeRouter struct {
	service service.IEmployeeService
}

func (er *EmployeeRouter) InitApiRouter(router *gin.RouterGroup) {
	publicRouter := router.Group("employee")
	privateRouter := router.Group("employee")

	privateRouter.Use(middle.VerifyJWTAdmin())

	er.service = service.NewEmployeeService(
		dao.NewEmployeeDao(global.DB, global.Redis),
	)
	employeeCtl := controller.NewEmployeeController(er.service)
	{
		publicRouter.POST("/login", employeeCtl.Login)
		privateRouter.POST("", employeeCtl.AddEmployee)
		//privateRouter.GET("/page", employeeCtl.PageQuery)
		publicRouter.GET("/page", employeeCtl.PageQuery)
	}
}
