package service

import (
	"member-service/config"
	"member-service/model"
	"member-service/producer"
	"member-service/repository"
	"member-service/validation"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewAddressService(addressRepository *repository.AddressRepository, provinsiRepository *repository.ProvinsiRepository, kotakabupatenRepository *repository.KotakabupatenRepository, kecamatanRepository *repository.KecamatanRepository, kelurahanRepository *repository.KelurahanRepository, postcodeRepository *repository.PostcodeRepository, registerProducer *producer.RegisterProducer, config config.Config, bucket *oss.Bucket) AddressService {
	return &addressServiceImpl{
		AddressRepository:       *addressRepository,
		ProvinsiRepository:      *provinsiRepository,
		KotakabupatenRepository: *kotakabupatenRepository,
		KecamatanRepository:     *kecamatanRepository,
		KelurahanRepository:     *kelurahanRepository,
		PostcodeRepository:      *postcodeRepository,
		RegisterProducer:        *registerProducer,
		Config:                  config,
		Bucketprofile:           bucket,
	}
}

type addressServiceImpl struct {
	AddressRepository       repository.AddressRepository
	ProvinsiRepository      repository.ProvinsiRepository
	KotakabupatenRepository repository.KotakabupatenRepository
	KecamatanRepository     repository.KecamatanRepository
	KelurahanRepository     repository.KelurahanRepository
	PostcodeRepository      repository.PostcodeRepository
	RegisterProducer        producer.RegisterProducer
	Config                  config.Config
	Bucketprofile           *oss.Bucket
}

func (service *addressServiceImpl) GetProvinsi(request model.GetProvinsiRequest) (response []model.GetProvinsiResponse) {
	validation.ValidateProvinsi(request)

	response = service.AddressRepository.FindAllProvinsi(request)
	if response == nil {
		response = []model.GetProvinsiResponse{}
	}

	return response
}

func (service *addressServiceImpl) GetColProvinsi(request model.GetColProvinsiRequest) (response []model.GetColProvinsiResponse) {
	validation.ValidateColProvinsi(request)

	response = service.ProvinsiRepository.FindAllColProvinsi(request)
	if response == nil {
		response = []model.GetColProvinsiResponse{}
	}

	return response
}

func (service *addressServiceImpl) GetKotaKabupaten(request model.GetKotaKabupatenRequest) (response []model.GetKotaKabupatenResponse) {
	validation.ValidateKotakabupaten(request)

	response = service.AddressRepository.FindAllKotaKabupaten(request)
	if response == nil {
		response = []model.GetKotaKabupatenResponse{}
	}

	return response
}

func (service *addressServiceImpl) GetColKotaKabupaten(request model.GetKotaKabupatenRequest) (response []model.GetKotaKabupatenResponse) {
	validation.ValidateKotakabupaten(request)

	response = service.KotakabupatenRepository.FindAllColKotaKabupaten(request)
	if response == nil {
		response = []model.GetKotaKabupatenResponse{}
	}

	return response
}

func (service *addressServiceImpl) GetKecamatan(request model.GetKecamatanRequest) (response []model.GetKecamatanResponse) {
	validation.ValidateKecamatan(request)

	response = service.AddressRepository.FindAllKecamatan(request)
	if response == nil {
		response = []model.GetKecamatanResponse{}
	}

	return response
}

func (service *addressServiceImpl) GetColKecamatan(request model.GetKecamatanRequest) (response []model.GetKecamatanResponse) {
	validation.ValidateKecamatan(request)

	response = service.KecamatanRepository.FindAllColKecamatan(request)

	if response == nil {
		response = []model.GetKecamatanResponse{}
	}

	return response
}

func (service *addressServiceImpl) GetKelurahan(request model.GetKelurahanRequest) (response []model.GetKelurahanResponse) {
	validation.ValidateKelurahan(request)

	response = service.AddressRepository.FindAllKelurahan(request)
	if response == nil {
		response = []model.GetKelurahanResponse{}
	}

	return response
}

func (service *addressServiceImpl) GetColKelurahan(request model.GetKelurahanRequest) (response []model.GetKelurahanResponse) {
	validation.ValidateKelurahan(request)

	response = service.KelurahanRepository.FindAllColKelurahan(request)
	if response == nil {
		response = []model.GetKelurahanResponse{}
	}

	return response
}

func (service *addressServiceImpl) GetKodepos(request model.GetAddressByIdKecamatanRequest) (response []model.GetAddressResponse) {
	validation.ValidateAddressByIdKecamatan(request)

	response = service.AddressRepository.FindAllPostCode(request)
	if response == nil {
		response = []model.GetAddressResponse{}
	}

	return response
}

func (service *addressServiceImpl) GetAllAddress(request model.GetAddressRequest) (response []model.GetAddressResponse) {
	validation.ValidateAddress(request)

	response = service.AddressRepository.FindAllAddress(request)
	if response == nil {
		response = []model.GetAddressResponse{}
	}

	return response
}

func (service *addressServiceImpl) GetKodeposByKode(request model.GetPostcodeRequest) (response []model.GetPostcodeResponse) {
	validation.Validatekodepos(request)

	response = service.PostcodeRepository.FindAllColPostcode(request)
	if response == nil {
		response = []model.GetPostcodeResponse{}
	}

	return response
}
