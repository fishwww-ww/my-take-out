package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"my-take-out/common/retcode"
	"my-take-out/common/utils"
	"my-take-out/global"
	"my-take-out/internal/model"
	"my-take-out/internal/service"
	"os"
)

type CommonController struct {
	service service.ICommonService
}

func (cc *CommonController) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		return
	}
	uuid := uuid.New()
	imageName := uuid.String() + file.Filename
	fileInfo := model.File{
		Uuid: uuid.String(),
		Name: file.Filename,
	}
	err = cc.service.Insert(ctx, fileInfo)
	if err != nil {
		global.Log.Error(err)
		retcode.Fatal(ctx, err, "")
		return
	}
	imagePath, err := utils.AliyunOss(imageName, file)
	if err != nil {
		global.Log.Warn("AliyunOss upload failed", "err", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, imagePath)
}

func (cc *CommonController) Download(ctx *gin.Context) {
	// 获取文件名称
	fileName := ctx.Param("file")
	if fileName == "" {
		return
	}
	outputPath := "./file/" + fileName + ".txt"

	// 获取uuid,拼接oss里的文件名称
	uuid, err := cc.service.QueryUuid(ctx, fileName)
	if err != nil {
		global.Log.Error(err)
		retcode.Fatal(ctx, err, "")
		return
	}
	ossFileName := uuid + fileName
	// 下载文件

	data, err := utils.Download(ossFileName)
	if err != nil {
		global.Log.Fatal("failed to get object: ", err)
		retcode.Fatal(ctx, err, "")
		return
	}
	err = os.WriteFile(outputPath, data, 0666)
	if err != nil {
		global.Log.Error()
		return
	}
}
