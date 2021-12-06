package service

import (
	"member-service/model"
)

type MemberAddressService interface {
	// List() (responses []model.GetMemberResponse)
	CreateAddress(request model.CreateAddressRequest, params ...string) (response model.CreateAddressResponse)
	UpdateAddress(request model.CreateAddressRequest, params ...string) (response model.CreateAddressResponse)
	UpdateIsDefaultAddress(request model.UpdateIsDefaultRequest, params ...string) (response string)
	DeleteAddress(request model.DeleteAddressRequest, params ...string) (response string)
	FindAddress(request model.GetAddressRequest, params ...string) (response []model.GetAddressFindResponse)
}
