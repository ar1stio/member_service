package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"member-service/model"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemberController_Provinsi(t *testing.T) {
	request := httptest.NewRequest("GET", "/api/member/provinsi/2/namaprovinsi/", nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	log.Println("response", response)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	list := webResponse.Data.(interface{})

	jsonData, _ := json.Marshal(list)
	getAuthResponse := model.GetProvinsiResponse{}
	json.Unmarshal(jsonData, &getAuthResponse)

	assert.NotNil(t, getAuthResponse.Id)
	assert.NotNil(t, getAuthResponse.Provinsi)

}

func TestMemberController_KotaKabupaten(t *testing.T) {
	request := httptest.NewRequest("GET", "/api/member/kota-kabupaten/1/namakotakabupaten/", nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	log.Println("response", response)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	list := webResponse.Data.(interface{})

	jsonData, _ := json.Marshal(list)
	getAuthResponse := model.GetKotaKabupatenResponse{}
	json.Unmarshal(jsonData, &getAuthResponse)

	assert.NotNil(t, getAuthResponse.Id)
	assert.NotNil(t, getAuthResponse.KotaKabupaten)

}

func TestMemberController_Kecamatan(t *testing.T) {
	request := httptest.NewRequest("GET", "/api/member/kecamatan/1/namakecamatan/", nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	log.Println("response", response)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	list := webResponse.Data.(interface{})

	jsonData, _ := json.Marshal(list)
	getAuthResponse := model.GetKecamatanResponse{}
	json.Unmarshal(jsonData, &getAuthResponse)

	assert.NotNil(t, getAuthResponse.Id)
	assert.NotNil(t, getAuthResponse.Kecamatan)

}

func TestMemberController_Kelurahan(t *testing.T) {
	request := httptest.NewRequest("GET", "/api/member/kelurahan/1/namakelurahan/", nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	log.Println("response", response)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	list := webResponse.Data.(interface{})

	jsonData, _ := json.Marshal(list)
	getAuthResponse := model.GetKelurahanResponse{}
	json.Unmarshal(jsonData, &getAuthResponse)

	assert.NotNil(t, getAuthResponse.Id)
	assert.NotNil(t, getAuthResponse.Kelurahan)

}

func TestMemberController_PostcodeByKecamatan(t *testing.T) {
	request := httptest.NewRequest("GET", "/api/member/kodepos/1", nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	log.Println("response", response)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	list := webResponse.Data.(interface{})

	jsonData, _ := json.Marshal(list)
	getAuthResponse := model.GetAddressResponse{}
	json.Unmarshal(jsonData, &getAuthResponse)

	assert.NotNil(t, getAuthResponse.Id)
	assert.NotNil(t, getAuthResponse.KodePos)

}

func TestMemberController_PostcodeByFilter(t *testing.T) {
	request := httptest.NewRequest("GET", "/api/member/all-address/Bandung", nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	log.Println("response", response)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	list := webResponse.Data.(interface{})

	jsonData, _ := json.Marshal(list)
	getAuthResponse := model.GetAddressResponse{}
	json.Unmarshal(jsonData, &getAuthResponse)

	assert.NotNil(t, getAuthResponse.Id)
	assert.NotNil(t, getAuthResponse.KodePos)

}

func TestMemberController_PostcodeById(t *testing.T) {
	request := httptest.NewRequest("GET", "/api/member/allkodepos/40193", nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	log.Println("response", response)

	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	list := webResponse.Data.(interface{})

	jsonData, _ := json.Marshal(list)
	getAuthResponse := model.GetAddressResponse{}
	json.Unmarshal(jsonData, &getAuthResponse)

	assert.NotNil(t, getAuthResponse.Id)
	assert.NotNil(t, getAuthResponse.KodePos)

}
