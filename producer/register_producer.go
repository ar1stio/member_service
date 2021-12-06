package producer

import (
	"member-service/model"
)

type RegisterProducer interface {
	CreateRegisterEmail(regis model.RegisterUserEvent)
}
