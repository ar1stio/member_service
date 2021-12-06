package config

import (
	"member-service/exception"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewOss(configuration Config) *oss.Client {

	endpoint := configuration.Get("OSS_ENDPOINT")
	keyId := configuration.Get("OSS_KEY_ID")
	keySecret := configuration.Get("OSS_KEY_SECRET")

	client, err := oss.New(endpoint, keyId, keySecret)
	exception.PanicIfNeeded(err)
	return client
}

func NewBucket(buck string, client *oss.Client, configuration Config) *oss.Bucket {
	bucket, err := client.Bucket(buck)
	if err != nil {
		// fmt.Println("failed read bucket")
	}
	return bucket
}
