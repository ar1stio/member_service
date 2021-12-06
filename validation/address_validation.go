package validation

import (
	"member-service/exception"
	"member-service/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateIsDefault(request model.UpdateIsDefaultRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.UserId, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidateProvinsi(request model.GetProvinsiRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidateColProvinsi(request model.GetColProvinsiRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidateKotakabupaten(request model.GetKotaKabupatenRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidateKecamatan(request model.GetKecamatanRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidateKelurahan(request model.GetKelurahanRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidateAddressByIdKecamatan(request model.GetAddressByIdKecamatanRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func Validatekodepos(request model.GetPostcodeRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Kodepos, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidateAddress(request model.GetAddressRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Filter, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
