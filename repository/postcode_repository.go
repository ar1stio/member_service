package repository

import (
	"member-service/model"
)

type PostcodeRepository interface {
	FindAllColPostcode(request model.GetPostcodeRequest) (response []model.GetPostcodeResponse)
}
