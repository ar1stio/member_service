package repository

import (
	"member-service/entity"
	"member-service/model"
)

type MemberAddressRepository interface {
	CreateAddress(address entity.Address)
	UpdateAddress(address entity.Address)
	DeleteAddress(address model.DeleteAddressRequest)
	DeleteAll()
	UpdateNonDefaultAddress(address model.UpdateIsDefaultRequest)
	UpdateIsDefaultAddress(address model.UpdateIsDefaultRequest)
	FindAddress(request model.GetAddressRequest) (response []model.GetAddressFindResponse)
	FindAddressOne(request model.GetAddressOneRequest) (memberaddress model.GetCekAddressFindResponse)
}
