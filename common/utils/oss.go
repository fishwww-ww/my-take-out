package utils

import (
	"context"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"io"
	"log"
	"mime/multipart"
	"my-take-out/global"
)

var (
	AccessKeyId     = global.Config.AliOss.AccessKeyId
	AccessKeySecret = global.Config.AliOss.AccessKeySecret
	EndPoint        = global.Config.AliOss.EndPoint
	BucketName      = global.Config.AliOss.BucketName
)

func NewOssClient() *oss.Client {
	config := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewStaticCredentialsProvider(AccessKeyId, AccessKeySecret)).
		WithRegion(EndPoint)
	client := oss.NewClient(config)
	return client
}

func AliyunOss(fileName string, file *multipart.FileHeader) (string, error) {
	client := NewOssClient()
	// 上传文件
	fileData, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileData.Close()

	putObjectInput := &oss.PutObjectRequest{
		Bucket: oss.Ptr(BucketName),
		Key:    oss.Ptr(fileName),
		Body:   fileData,
	}

	_, err = client.PutObject(context.TODO(), putObjectInput)
	if err != nil {
		return "", err
	}

	imagePath := "https://" + BucketName + ".oss-" + EndPoint + "/" + fileName
	return imagePath, nil
}

func Download(fileName string) ([]byte, error) {
	client := NewOssClient()
	// 创建获取对象的请求
	request := &oss.GetObjectRequest{
		Bucket: oss.Ptr(BucketName), // 存储空间名称
		Key:    oss.Ptr(fileName),   // 对象名称
	}
	// 执行获取对象的操作并处理结果
	result, err := client.GetObject(context.TODO(), request)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close() // 确保在函数结束时关闭响应体

	log.Printf("get object result:%#v\n", result)

	// 读取对象的内容
	data, _ := io.ReadAll(result.Body)
	global.Log.Info(string(data))
	return data, nil
}
