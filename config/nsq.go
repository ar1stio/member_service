package config

import (
	"github.com/nsqio/go-nsq"
	log "github.com/sirupsen/logrus"
)

func NewNsq(configuration Config) *nsq.Config {
	config := nsq.NewConfig()
	return config
}

func NewProducer(config *nsq.Config, configuration Config) *nsq.Producer {
	// p, err := nsq.NewProducer("127.0.0.1:4150", config)
	p, err := nsq.NewProducer(configuration.Get("NSQ_URI"), config)

	if err != nil {
		log.Error(err, "Failed to create producer")
	}
	return p
}

// func NewProducer(configuration Config) *nsq.Producer {
// 	config := nsq.NewConfig()
// 	// p, err := nsq.NewProducer("127.0.0.1:4150", config)
// 	p, err := nsq.NewProducer(configuration.Get("NSQ_URI"), config)

// 	if err != nil {
// 		log.Error(err, "Failed to create producer")
// 	}
// 	return p
// }
