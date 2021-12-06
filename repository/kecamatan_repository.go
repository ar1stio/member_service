package repository

import (
	"member-service/model"
)

type KecamatanRepository interface {
	FindAllColKecamatan(request model.GetKecamatanRequest) (response []model.GetKecamatanResponse)
}
