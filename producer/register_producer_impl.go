package producer

import (
	"encoding/json"
	"member-service/model"

	"github.com/nsqio/go-nsq"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func NewRegisterProducer(producer *nsq.Producer, rabbitmq *amqp.Channel) RegisterProducer {
	return &registerProducerImpl{
		Producer: producer,
		Rabbitmq: rabbitmq,
	}
}

type registerProducerImpl struct {
	Producer *nsq.Producer
	Rabbitmq *amqp.Channel
}

func (producer *registerProducerImpl) CreateRegisterEmail(regis model.RegisterUserEvent) {
	msg, err := json.Marshal(regis)

	err = producer.Producer.Publish("register", []byte(string(msg)))
	if err != nil {
		log.Error(err, "Failed to publish a message")
	}
}
