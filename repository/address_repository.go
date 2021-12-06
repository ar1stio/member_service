package repository

import (
	"member-service/model"
)

type AddressRepository interface {
	FindAllProvinsi(request model.GetProvinsiRequest) (response []model.GetProvinsiResponse)
	FindAllKotaKabupaten(request model.GetKotaKabupatenRequest) (response []model.GetKotaKabupatenResponse)
	FindAllKecamatan(request model.GetKecamatanRequest) (response []model.GetKecamatanResponse)
	FindAllKelurahan(request model.GetKelurahanRequest) (response []model.GetKelurahanResponse)
	FindAllAddress(request model.GetAddressRequest) (response []model.GetAddressResponse)
	FindAllPostCode(request model.GetAddressByIdKecamatanRequest) (response []model.GetAddressResponse)
}
