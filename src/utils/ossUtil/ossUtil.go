package ossUtil

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"path/filepath"
	"tiktok/src/config"
	"time"
)

var bucket *oss.Bucket

//
// InitBucket
// @Description: OSS初始化返回Bucket
// @return *ossUtil.Bucket
//
func InitBucket() {
	var (
		endpoint, accessKeyId, accessKeySecret, bucketName string
	)

	endpoint = config.AppConfig.OSS.EndPoint
	accessKeyId = config.AppConfig.OSS.AccessKeyId
	accessKeySecret = config.AppConfig.OSS.AccessKeySecret
	bucketName = config.AppConfig.OSS.BucketName
	//根据配置生成OSS Client
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	//根据配置创建Bucket
	bucket, err = client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		panic("ossUtil init error!")
	}
}

func UploadFile(fileName string, fileByte []byte) (url string, err error) {
	folderName := time.Now().Format("2006-01-02")
	fileTmpPath := filepath.Join("uploads"+folderName) + "/" + fileName
	if err := bucket.PutObject(fileTmpPath, bytes.NewReader([]byte(fileByte))); err != nil {
		return "", err
	}
	return config.AppConfig.OSS.Domain + "/" + fileTmpPath, nil
}
