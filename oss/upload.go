package oss

import (
	"Java2GO/pkg/setting"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"os"
)

var bucket *oss.Bucket

//
// InitBucket
// @Description: OSS初始化返回Bucket
// @return *oss.Bucket
//

func InitBucket() *oss.Bucket {
	var (
		endpoint, accessKeyId, accessSecret, bucketName string
	)

	sec, err := setting.Cfg.GetSection("oss")
	if err != nil {
		log.Fatal(2, "Fail to get section 'oss': %v", err)
	}
	endpoint = sec.Key("ENDPOINT").String()
	accessKeyId = sec.Key("ACCESSKEYID").String()
	accessSecret = sec.Key("ACCESSKETSERCERT").String()
	bucketName = sec.Key("BUCKETNAME").String()
	//根据配置生成OSS Client
	client, err := oss.New(endpoint, accessKeyId, accessSecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	//根据配置返回Bucket
	bucket, err = client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return bucket
}
