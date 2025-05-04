package initialize

import (
	"github.com/gin-gonic/gin"
	"my-take-out/config"
	"my-take-out/global"
)

func GlobalInit() *gin.Engine {
	global.Config = config.InitLoadConfig()
	global.DB = InitDatabase(global.Config.DataSource.Dsn())
	router := routerInit()
	return router
}
