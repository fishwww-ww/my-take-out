package admin

import (
	"github.com/gin-gonic/gin"
	"my-take-out/internal/api/controller"
	"my-take-out/middle"
)

type CommonRouter struct{}

func (dr *CommonRouter) InitApiRouter(parent *gin.RouterGroup) {
	privateRouter := parent.Group("common")
	privateRouter.Use(middle.VerifyJWTAdmin())
	commonCtrl := new(controller.CommonController)
	{
		privateRouter.POST("upload", commonCtrl.Upload)
	}
}
