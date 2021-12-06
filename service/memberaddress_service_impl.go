package service

import (
	"encoding/json"
	"io/ioutil"
	"member-service/config"
	"member-service/entity"
	"member-service/exception"
	"member-service/logger"
	"member-service/model"
	"member-service/producer"
	"member-service/repository"
	"member-service/util"
	"member-service/validation"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
)

func NewMemberAddressService(memberaddressRepository *repository.MemberAddressRepository, addressRepository *repository.AddressRepository, registerProducer *producer.RegisterProducer, config config.Config, bucket *oss.Bucket) MemberAddressService {
	return &memberaddressServiceImpl{
		AddressRepository:       *addressRepository,
		MemberAddressRepository: *memberaddressRepository,
		RegisterProducer:        *registerProducer,
		Config:                  config,
		Bucketprofile:           bucket,
	}
}

type memberaddressServiceImpl struct {
	AddressRepository       repository.AddressRepository
	MemberAddressRepository repository.MemberAddressRepository
	RegisterProducer        producer.RegisterProducer
	Config                  config.Config
	Bucketprofile           *oss.Bucket
}

func (service *memberaddressServiceImpl) CreateAddress(request model.CreateAddressRequest, params ...string) (response model.CreateAddressResponse) {
	validation.ValidateCreateAddress(request)

	reqCekDefault := model.GetAddressOneRequest{
		Type:   "nama",
		Filter: request.Name,
		UserId: request.UserId,
	}
	resCekDefault := service.MemberAddressRepository.FindAddressOne(reqCekDefault)

	if resCekDefault.UserId != "" {
		exception.PanicIfNeeded(exception.ConflictError{Status: "CONFLICT", Message: "Maaf nama alamat sudah ada"})
	}

	nowInt := time.Now().UnixNano() / int64(time.Millisecond)
	loc, _ := time.LoadLocation(service.Config.Get("TIME_ZONE"))
	server := time.Now().Local()
	local := time.Now().In(loc).String()

	created := model.Created{
		By:         request.UserId,
		At:         nowInt,
		DateServer: server,
		DateLocal:  local,
	}

	modified := model.Modified{
		By:         request.UserId,
		At:         nowInt,
		DateServer: server,
		DateLocal:  local,
	}

	Address := strings.ReplaceAll(request.Address, " ", "+")
	resAddress := GetDataLongLat(Address)
	longitude := 0.0
	latitude := 0.0
	if resAddress.Status == "OK" {
		mlng := resAddress.Result[0].GeometryStruct.Location.Lng
		// slng := strconv.FormatFloat(mlng, 'E', -1, 64)
		mlat := resAddress.Result[0].GeometryStruct.Location.Lat
		// slat := strconv.FormatFloat(mlat, 'E', -1, 64)

		// longitude = slng
		longitude = mlng

		// latitude = slat
		latitude = mlat
	} else {
		resAddress := GetDataLongLat(Address)
		if resAddress.Status == "OK" {
			mlng := resAddress.Result[0].GeometryStruct.Location.Lng
			// slng := strconv.FormatFloat(mlng, 'E', -1, 64)
			mlat := resAddress.Result[0].GeometryStruct.Location.Lat
			// slat := strconv.FormatFloat(mlat, 'E', -1, 64)

			// longitude = slng
			longitude = mlng

			// latitude = slat
			latitude = mlat

		}
	}

	idaddress := uuid.New().String()

	memberinput := entity.Address{
		Id:              idaddress,
		UserId:          request.UserId,
		Name:            request.Name,
		IsDefault:       request.IsDefault,
		ReceiverName:    request.ReceiverName,
		Phone:           request.Phone,
		IdProvinsi:      request.IdProvinsi,
		Provinsi:        request.Provinsi,
		IdKotaKabupaten: request.IdKotaKabupaten,
		KotaKabupaten:   request.KotaKabupaten,
		IdKecamatan:     request.IdKecamatan,
		Kecamatan:       request.Kecamatan,
		IdKelurahan:     request.IdKelurahan,
		Kelurahan:       request.Kelurahan,
		IdKodepos:       request.IdKodepos,
		Kodepos:         request.Kodepos,
		Latitude:        latitude,
		Longitude:       longitude,
		Address:         request.Address,
		Created:         created,
		Modified:        modified,
	}

	service.MemberAddressRepository.CreateAddress(memberinput)

	response = model.CreateAddressResponse{
		Id: memberinput.Id,
		// UserId:          request.UserId,
		Name:            request.Name,
		IsDefault:       request.IsDefault,
		ReceiverName:    request.ReceiverName,
		Phone:           request.Phone,
		IdProvinsi:      request.IdProvinsi,
		Provinsi:        request.Provinsi,
		IdKotaKabupaten: request.IdKotaKabupaten,
		KotaKabupaten:   request.KotaKabupaten,
		IdKecamatan:     request.IdKecamatan,
		Kecamatan:       request.Kecamatan,
		IdKelurahan:     request.IdKelurahan,
		Kelurahan:       request.Kelurahan,
		IdKodepos:       request.IdKodepos,
		Kodepos:         request.Kodepos,
		Latitude:        request.Latitude,
		Longitude:       request.Longitude,
		Address:         request.Address,
	}

	zlog := logger.NewLogger()
	zlog.Info().
		Str("endpoint", params[0]).
		Int16("httpResCode", 200).
		// RawJSON("req", util.JsonStruct(request)).
		RawJSON("req", util.JsonStruct("cek")).
		// RawJSON("res", util.JsonStruct(response)).
		RawJSON("res", util.JsonStruct("cek")).
		Msg("new address created")

	return response
}

func (service *memberaddressServiceImpl) UpdateAddress(request model.CreateAddressRequest, params ...string) (response model.CreateAddressResponse) {
	validation.ValidateUpdateAddress(request)

	nowInt := time.Now().UnixNano() / int64(time.Millisecond)
	loc, _ := time.LoadLocation(service.Config.Get("TIME_ZONE"))
	server := time.Now().Local()
	local := time.Now().In(loc).String()

	created := model.Created{
		By:         request.UserId,
		At:         nowInt,
		DateServer: server,
		DateLocal:  local,
	}

	modified := model.Modified{
		By:         request.UserId,
		At:         nowInt,
		DateServer: server,
		DateLocal:  local,
	}

	Address := strings.ReplaceAll(request.Address, " ", "+")
	resAddress := GetDataLongLat(Address)
	longitude := 0.0
	latitude := 0.0
	if resAddress.Status == "OK" {
		mlng := resAddress.Result[0].GeometryStruct.Location.Lng
		// slng := strconv.FormatFloat(mlng, 'E', -1, 64)
		mlat := resAddress.Result[0].GeometryStruct.Location.Lat
		// slat := strconv.FormatFloat(mlat, 'E', -1, 64)

		longitude = mlng
		// longitude = slng

		// latitude = slat
		latitude = mlat

	} else {
		resAddress := GetDataLongLat(Address)
		if resAddress.Status == "OK" {
			mlng := resAddress.Result[0].GeometryStruct.Location.Lng
			// slng := strconv.FormatFloat(mlng, 'E', -1, 64)
			mlat := resAddress.Result[0].GeometryStruct.Location.Lat
			// slat := strconv.FormatFloat(mlat, 'E', -1, 64)

			// longitude = slng
			longitude = mlng

			// latitude = slat
			latitude = mlat
		}
	}

	memberinput := entity.Address{
		Id:              request.Id,
		UserId:          request.UserId,
		Name:            request.Name,
		IsDefault:       request.IsDefault,
		ReceiverName:    request.ReceiverName,
		Phone:           request.Phone,
		IdProvinsi:      request.IdProvinsi,
		Provinsi:        request.Provinsi,
		IdKotaKabupaten: request.IdKotaKabupaten,
		KotaKabupaten:   request.KotaKabupaten,
		IdKecamatan:     request.IdKecamatan,
		Kecamatan:       request.Kecamatan,
		IdKelurahan:     request.IdKelurahan,
		Kelurahan:       request.Kelurahan,
		IdKodepos:       request.IdKodepos,
		Kodepos:         request.Kodepos,
		Latitude:        latitude,
		Longitude:       longitude,
		Address:         request.Address,
		Created:         created,
		Modified:        modified,
	}

	service.MemberAddressRepository.UpdateAddress(memberinput)

	response = model.CreateAddressResponse{
		Id: request.Id,
		// UserId:          request.UserId,
		Name:            request.Name,
		IsDefault:       request.IsDefault,
		ReceiverName:    request.ReceiverName,
		Phone:           request.Phone,
		IdProvinsi:      request.IdProvinsi,
		Provinsi:        request.Provinsi,
		IdKotaKabupaten: request.IdKotaKabupaten,
		KotaKabupaten:   request.KotaKabupaten,
		IdKecamatan:     request.IdKecamatan,
		Kecamatan:       request.Kecamatan,
		IdKelurahan:     request.IdKelurahan,
		Kelurahan:       request.Kelurahan,
		IdKodepos:       request.IdKodepos,
		Kodepos:         request.Kodepos,
		Latitude:        request.Latitude,
		Longitude:       request.Longitude,
		Address:         request.Address,
	}

	zlog := logger.NewLogger()
	zlog.Info().
		Str("endpoint", params[0]).
		Int16("httpResCode", 200).
		// RawJSON("req", util.JsonStruct(request)).
		RawJSON("req", util.JsonStruct("cek")).
		// RawJSON("res", util.JsonStruct(response)).
		RawJSON("res", util.JsonStruct("cek")).
		Msg("new address updated")

	return response
}

func (service *memberaddressServiceImpl) UpdateIsDefaultAddress(request model.UpdateIsDefaultRequest, params ...string) (response string) {
	validation.ValidateIsDefault(request)

	service.MemberAddressRepository.UpdateNonDefaultAddress(request)
	service.MemberAddressRepository.UpdateIsDefaultAddress(request)

	return "Berhasil merubah menjadi alamat default"
}

func (service *memberaddressServiceImpl) DeleteAddress(request model.DeleteAddressRequest, params ...string) (response string) {
	validation.ValidateDeleteAddress(request)

	reqCekDefault := model.GetAddressOneRequest{
		Type:   "id",
		Filter: request.Id,
	}
	resCekDefault := service.MemberAddressRepository.FindAddressOne(reqCekDefault)

	if resCekDefault.UserId == "" {
		exception.PanicIfNeeded("Maaf alamat tidak ditemukan")
	}

	if resCekDefault.IsDefault == true {
		exception.PanicIfNeeded("Maaf tidak bisa menghapus alamat default")
	}

	service.MemberAddressRepository.DeleteAddress(request)

	return "Berhasil menghapus alamat"
}

func GetDataLongLat(address string) (response model.ApiResponse) {
	client := http.Client{Timeout: 10 * time.Second}
	request, _ := http.NewRequest("GET", os.Getenv("MapURL")+"address="+address+"&key="+os.Getenv("MapAPI"), nil)
	request.Header.Set("Content-type", "application/json")
	// request.Header.Set("Private-Authentication", token)
	// request.Header.Set("Public-Authorization", os.Getenv("core.client.properties.publicAuthorization"))
	resp, _ := client.Do(request)

	// defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	response, _ = decode(string(body), string(body))

	return response
}

func decode(input string, inputok string) (output model.ApiResponse, err error) {
	firstresponse := model.ApiFirstResponse{}
	err = json.Unmarshal([]byte(input), &firstresponse)

	if firstresponse.Status == "OK" {
		err = json.Unmarshal([]byte(inputok), &output)
		if err != nil {
			exception.PanicIfNeeded("Tidak dapat convert ke interface")
			// return output, err
		}

	}

	// output.Status = firstresponse.Status

	return output, nil
}

func (service *memberaddressServiceImpl) FindAddress(request model.GetAddressRequest, params ...string) (response []model.GetAddressFindResponse) {

	response = service.MemberAddressRepository.FindAddress(request)
	if response == nil {
		response = []model.GetAddressFindResponse{}
	}
	return response
}
