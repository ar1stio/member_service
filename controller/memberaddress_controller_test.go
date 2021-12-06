package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"member-service/entity"
	"member-service/model"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMemberAddressController_Create(t *testing.T) {
	memberaddressRepository.DeleteAll()
	userId := "f611bc5f-bdc4-437d-95d0-18a3203101c5"

	createAddressRequest := model.CreateAddressRequest{
		Name:            "dasds",
		IsDefault:       true,
		ReceiverName:    "Tio",
		Phone:           "085523865720",
		IdProvinsi:      "1",
		Provinsi:        "Dki Jakarta",
		IdKotaKabupaten: "2",
		KotaKabupaten:   "bandung",
		IdKecamatan:     "3",
		Kecamatan:       "coblong",
		IdKelurahan:     "2",
		Kelurahan:       "dago",
		IdKodepos:       "2",
		Kodepos:         "10111",
		Address:         "bandung dago",
	}
	requestBody, _ := json.Marshal(createAddressRequest)

	request := httptest.NewRequest("POST", "/api/member/address", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("x-auth-id", userId)
	request.Header.Set("x-auth-username", "daristet")
	request.Header.Set("x-auth-email", "daristio@gmail.com")
	request.Header.Set("x-auth-phone", "085523865720")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	jsonData, _ := json.Marshal(webResponse.Data)
	createAddressResponse := model.CreateAddressResponse{}
	json.Unmarshal(jsonData, &createAddressResponse)
	assert.NotNil(t, createAddressResponse.Id)
	assert.Equal(t, createAddressRequest.Name, createAddressResponse.Name)
}

func TestMemberAddressController_Update(t *testing.T) {
	memberaddressRepository.DeleteAll()
	userId := "f611bc5f-bdc4-437d-95d0-18a3203101c5"

	createAddressRequest := entity.Address{
		Id:              uuid.New().String(),
		Name:            "dasds",
		IsDefault:       true,
		ReceiverName:    "Tio",
		Phone:           "085523865720",
		IdProvinsi:      "1",
		Provinsi:        "Dki Jakarta",
		IdKotaKabupaten: "2",
		KotaKabupaten:   "bandung",
		IdKecamatan:     "3",
		Kecamatan:       "coblong",
		IdKelurahan:     "2",
		Kelurahan:       "dago",
		IdKodepos:       "2",
		Kodepos:         "10111",
		Address:         "bandung dago",
	}
	memberaddressRepository.CreateAddress(createAddressRequest)

	updateAddressRequest := model.CreateAddressRequest{
		Id:              createAddressRequest.Id,
		Name:            "dasds",
		IsDefault:       true,
		ReceiverName:    "Tio 1",
		Phone:           "085523865720",
		IdProvinsi:      "1",
		Provinsi:        "Dki Jakarta",
		IdKotaKabupaten: "2",
		KotaKabupaten:   "bandung",
		IdKecamatan:     "3",
		Kecamatan:       "coblong",
		IdKelurahan:     "2",
		Kelurahan:       "dago",
		IdKodepos:       "2",
		Kodepos:         "10111",
		Address:         "bandung dago",
	}
	requestBody, _ := json.Marshal(updateAddressRequest)

	request := httptest.NewRequest("PUT", "/api/member/address", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("x-auth-id", userId)
	request.Header.Set("x-auth-username", "daristet")
	request.Header.Set("x-auth-email", "daristio@gmail.com")
	request.Header.Set("x-auth-phone", "085523865720")

	response, _ := app.Test(request)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	jsonData, _ := json.Marshal(webResponse.Data)
	createAddressResponse := model.CreateAddressResponse{}
	json.Unmarshal(jsonData, &createAddressResponse)
	assert.NotNil(t, createAddressResponse.Id)
	assert.Equal(t, createAddressRequest.Name, createAddressResponse.Name)
}
