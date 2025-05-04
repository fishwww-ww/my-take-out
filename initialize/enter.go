package initialize

import "github.com/gin-gonic/gin"

func GlobalInit() *gin.Engine {
	router := routerInit()
	return router
}
