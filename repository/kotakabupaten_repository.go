package repository

import (
	"member-service/model"
)

type KotakabupatenRepository interface {
	FindAllColKotaKabupaten(request model.GetKotaKabupatenRequest) (response []model.GetKotaKabupatenResponse)
}
