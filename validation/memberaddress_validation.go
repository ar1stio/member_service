package validation

import (
	"member-service/exception"
	"member-service/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateDeleteAddress(request model.DeleteAddressRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		// validation.Field(&request.ProfilePicture, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidateCreateAddress(request model.CreateAddressRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.UserId, validation.Required),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.ReceiverName, validation.Required),
		validation.Field(&request.Phone, validation.Required),
		validation.Field(&request.IdProvinsi, validation.Required),
		validation.Field(&request.Provinsi, validation.Required),
		validation.Field(&request.IdKotaKabupaten, validation.Required),
		validation.Field(&request.KotaKabupaten, validation.Required),
		validation.Field(&request.IdKecamatan, validation.Required),
		validation.Field(&request.Kecamatan, validation.Required),
		validation.Field(&request.IdKelurahan, validation.Required),
		validation.Field(&request.Kelurahan, validation.Required),
		validation.Field(&request.IdKodepos, validation.Required),
		validation.Field(&request.Kodepos, validation.Required),
		validation.Field(&request.Address, validation.Required),
		// validation.Field(&request.ProfilePicture, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidateUpdateAddress(request model.CreateAddressRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.UserId, validation.Required),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.ReceiverName, validation.Required),
		validation.Field(&request.Phone, validation.Required),
		validation.Field(&request.IdProvinsi, validation.Required),
		validation.Field(&request.Provinsi, validation.Required),
		validation.Field(&request.IdKotaKabupaten, validation.Required),
		validation.Field(&request.KotaKabupaten, validation.Required),
		validation.Field(&request.IdKecamatan, validation.Required),
		validation.Field(&request.Kecamatan, validation.Required),
		validation.Field(&request.IdKelurahan, validation.Required),
		validation.Field(&request.Kelurahan, validation.Required),
		validation.Field(&request.IdKodepos, validation.Required),
		validation.Field(&request.Kodepos, validation.Required),
		validation.Field(&request.Address, validation.Required),
		// validation.Field(&request.ProfilePicture, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
