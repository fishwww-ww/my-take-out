package admin

import (
	"github.com/gin-gonic/gin"
	"my-take-out/global"
	"my-take-out/internal/api/controller"
	"my-take-out/internal/repository/dao"
	"my-take-out/internal/service"
	"my-take-out/middle"
)

type CommonRouter struct {
	service service.ICommonService
}

func (dr *CommonRouter) InitApiRouter(parent *gin.RouterGroup) {
	privateRouter := parent.Group("common")
	privateRouter.Use(middle.VerifyJWTAdmin())
	dr.service = service.NewCommonService(
		dao.NewCommonDao(global.DB))
	commonCtrl := new(controller.CommonController)
	{
		privateRouter.POST("upload", commonCtrl.Upload)
	}
}
