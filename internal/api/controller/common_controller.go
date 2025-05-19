package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"my-take-out/common/retcode"
	"my-take-out/common/utils"
	"my-take-out/global"
)

type CommonController struct{}

func (cc *CommonController) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		return
	}
	uuid := uuid.New()
	imageName := uuid.String() + file.Filename
	imagePath, err := utils.AliyunOss(imageName, file)
	if err != nil {
		global.Log.Warn("AliyunOss upload failed", "err", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, imagePath)
}
