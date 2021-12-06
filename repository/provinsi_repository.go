package repository

import (
	"member-service/model"
)

type ProvinsiRepository interface {
	FindAllColProvinsi(request model.GetColProvinsiRequest) (response []model.GetColProvinsiResponse)
}
