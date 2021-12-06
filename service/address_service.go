package service

import (
	"member-service/model"
)

type AddressService interface {
	// List() (responses []model.GetMemberResponse)
	GetProvinsi(request model.GetProvinsiRequest) (response []model.GetProvinsiResponse)
	GetColProvinsi(request model.GetColProvinsiRequest) (response []model.GetColProvinsiResponse)

	GetKotaKabupaten(request model.GetKotaKabupatenRequest) (response []model.GetKotaKabupatenResponse)
	GetColKotaKabupaten(request model.GetKotaKabupatenRequest) (response []model.GetKotaKabupatenResponse)

	GetKecamatan(request model.GetKecamatanRequest) (response []model.GetKecamatanResponse)
	GetColKecamatan(request model.GetKecamatanRequest) (response []model.GetKecamatanResponse)

	GetKelurahan(request model.GetKelurahanRequest) (response []model.GetKelurahanResponse)
	GetColKelurahan(request model.GetKelurahanRequest) (response []model.GetKelurahanResponse)

	GetKodepos(request model.GetAddressByIdKecamatanRequest) (response []model.GetAddressResponse)
	GetAllAddress(request model.GetAddressRequest) (response []model.GetAddressResponse)

	GetKodeposByKode(request model.GetPostcodeRequest) (response []model.GetPostcodeResponse)
}
