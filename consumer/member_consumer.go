package consumer

import (
	"member-service/exception"
	"member-service/service"

	"github.com/streadway/amqp"
)

func NewMemberConsumer(amqpCon *amqp.Connection, memberService service.MemberService) MemberConsumer {
	return MemberConsumer{
		amqpCon:       amqpCon,
		memberService: memberService,
	}
}

type MemberConsumer struct {
	amqpCon       *amqp.Connection
	memberService service.MemberService
}

func (consumer *MemberConsumer) getQueue(queueName string) <-chan amqp.Delivery {
	ch, err := consumer.amqpCon.Channel()
	exception.PanicIfNeeded(err)

	queue, err := ch.Consume(queueName, "", false, false, false, false, nil)
	exception.PanicIfNeeded(err)

	return queue
}

// region ACTION 1: Auth register verification Code

func (consumer *MemberConsumer) AddNewMember() {
	queueName := "member-create-new-member"
	queue := consumer.getQueue(queueName)

	for msg := range queue {
		consumer.memberService.CreateConsumer(msg.Body, queueName)
		msg.Ack(false)
	}
}
