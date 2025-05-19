package utils

import (
	"context"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"mime/multipart"
	"my-take-out/global"
)

func AliyunOss(fileName string, file *multipart.FileHeader) (string, error) {
	config := global.Config.AliOss
	AliyunConfig := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewStaticCredentialsProvider(config.AccessKeyId, config.AccessKeySecret)).
		WithRegion(config.EndPoint)
	client := oss.NewClient(AliyunConfig)
	// 上传文件
	fileData, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileData.Close()

	putObjectInput := &oss.PutObjectRequest{
		Bucket: oss.Ptr(config.BucketName),
		Key:    oss.Ptr(fileName),
		Body:   fileData,
	}

	_, err = client.PutObject(context.TODO(), putObjectInput)
	if err != nil {
		return "", err
	}

	imagePath := "https://" + config.BucketName + ".oss-" + config.EndPoint + "/" + fileName
	return imagePath, nil
}
