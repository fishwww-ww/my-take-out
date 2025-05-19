package initialize

import (
	"github.com/gin-gonic/gin"
	"my-take-out/internal/router"
)

func routerInit() *gin.Engine {
	r := gin.Default()
	allRouter := router.AllRouter
	admin := r.Group("admin")
	{
		allRouter.EmployeeRouter.InitApiRouter(admin)
		allRouter.CommonRouter.InitApiRouter(admin)
	}
	return r
}
