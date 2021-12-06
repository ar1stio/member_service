package repository

import (
	"member-service/model"
)

type KelurahanRepository interface {
	FindAllColKelurahan(request model.GetKelurahanRequest) (response []model.GetKelurahanResponse)
}
