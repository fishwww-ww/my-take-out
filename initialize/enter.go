package initialize

import (
	"github.com/gin-gonic/gin"
	"my-take-out/config"
	"my-take-out/global"
	"my-take-out/logger"
)

func GlobalInit() *gin.Engine {
	global.Config = config.InitLoadConfig()
	global.Log = logger.NewMySLog(global.Config.Log.Level, global.Config.Log.FilePath)
	global.DB = InitDatabase(global.Config.DataSource.Dsn())
	router := routerInit()
	return router
}
