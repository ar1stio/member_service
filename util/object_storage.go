package util

import (
	"member-service/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func ObjectStorageClient() (config.Config, *oss.Client, error) {
	conf := config.New()
	client, err := oss.New(conf.Get("OSS_ENDPOINT"), conf.Get("OSS_KEY_ID"), conf.Get("OSS_KEY_SECRET"))
	return conf, client, err
}
