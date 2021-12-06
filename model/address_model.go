package model

type GetProvinsiRequest struct {
	Id       string `json:"_id"`
	Provinsi string `json:"provinsi"`
}

type GetColProvinsiRequest struct {
	Id       int    `json:"_id"`
	Provinsi string `json:"provinsi"`
}

type GetColProvinsiResponse struct {
	Id       int32  `json:"_id"`
	Provinsi string `json:"provinsi"`
}

type GetProvinsiResponse struct {
	Id       string `json:"_id"`
	Provinsi string `json:"provinsi"`
}

type GetKotaKabupatenRequest struct {
	Id            string `json:"_id"`
	KotaKabupaten string `json:"kota_kabupaten"`
}

type GetKotaKabupatenResponse struct {
	Id            int32  `json:"_id"`
	IdProvinsi    string `json:"_id_provinsi"`
	Provinsi      string `json:"provinsi"`
	KotaKabupaten string `json:"kota_kabupaten"`
}

type GetKecamatanRequest struct {
	Id        string `json:"_id"`
	Kecamatan string `json:"kecamatan"`
}

type GetKecamatanResponse struct {
	Id            int32  `json:"_id"`
	IdProvinsi    string `json:"_id_provinsi"`
	Provinsi      string `json:"provinsi"`
	IdKotaKab     string `json:"_id_kota_kabupaten"`
	KotaKabupaten string `json:"kota_kabupaten"`
	Kecamatan     string `json:"kecamatan"`
}

type GetKelurahanRequest struct {
	Id        string `json:"_id"`
	Kelurahan string `json:"kelurahan"`
}

type GetKelurahanResponse struct {
	Id              int32  `json:"_id"`
	IdProvinsi      string `json:"_id_provinsi"`
	Provinsi        string `json:"provinsi"`
	IdKotaKabupaten string `json:"_id_kota_kabupaten"`
	KotaKabupaten   string `json:"kota_kabupaten"`
	IdKecamatan     string `json:"_id_kecamatan"`
	Kecamatan       string `json:"kecamatan"`
	Kelurahan       string `json:"kelurahan"`
}

type GetAddressRequest struct {
	Filter string `json:"filter"`
}

type GetAddressOneRequest struct {
	Filter string `json:"filter"`
	Type   string `json:"type"`
	UserId string `json:"user_id"`
}

type GetCekAddressFindResponse struct {
	Id              string  `json:"_id"`
	UserId          string  `json:"user_id"`
	Name            string  `json:"name"`
	IsDefault       bool    `json:"is_default"`
	ReceiverName    string  `json:"receiver_name"`
	Phone           string  `json:"phone"`
	IdProvinsi      string  `json:"_id_provinsi"`
	Provinsi        string  `json:"provinsi"`
	IdKotaKabupaten string  `json:"_id_kota_kabupaten"`
	KotaKabupaten   string  `json:"kota_kabupaten"`
	IdKecamatan     string  `json:"_id_kecamatan"`
	Kecamatan       string  `json:"kecamatan"`
	IdKelurahan     string  `json:"_id_kelurahan"`
	Kelurahan       string  `json:"kelurahan"`
	IdKodepos       string  `json:"_id_kode_pos"`
	Kodepos         string  `json:"kode_pos"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Address         string  `json:"address"`
}

type GetAddressFindResponse struct {
	Id string `json:"_id"`
	// UserId          string  `json:"user_id"`
	Name            string  `json:"name"`
	IsDefault       bool    `json:"is_default"`
	ReceiverName    string  `json:"receiver_name"`
	Phone           string  `json:"phone"`
	IdProvinsi      string  `json:"_id_provinsi"`
	Provinsi        string  `json:"provinsi"`
	IdKotaKabupaten string  `json:"_id_kota_kabupaten"`
	KotaKabupaten   string  `json:"kota_kabupaten"`
	IdKecamatan     string  `json:"_id_kecamatan"`
	Kecamatan       string  `json:"kecamatan"`
	IdKelurahan     string  `json:"_id_kelurahan"`
	Kelurahan       string  `json:"kelurahan"`
	IdKodepos       string  `json:"_id_kode_pos"`
	Kodepos         string  `json:"kode_pos"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Address         string  `json:"address"`
}

type GetAddressByIdKecamatanRequest struct {
	Id string `json:"_id"`
}

type GetAddressResponse struct {
	Id              int32  `json:"_id"`
	IdProvinsi      string `json:"_id_provinsi"`
	Provinsi        string `json:"provinsi"`
	IdKotaKabupaten string `json:"_id_kota_kabupaten"`
	IdKelurahan     string `json:"_id_kelurahan"`
	KotaKabupaten   string `json:"kota_kabupaten"`
	IdKecamatan     string `json:"_id_kecamatan"`
	Kecamatan       string `json:"kecamatan"`
	Kelurahan       string `json:"kelurahan"`
	KodePos         string `json:"kodepos"`
	Address         string `json:"address"`
	// Created         interface{}
	// Modified        interface{}
}

type CreateAddressRequest struct {
	Id              string  `json:"_id"`
	UserId          string  `json:"user_id"`
	Name            string  `json:"name"`
	IsDefault       bool    `json:"is_default"`
	ReceiverName    string  `json:"receiver_name"`
	Phone           string  `json:"phone"`
	IdProvinsi      string  `json:"_id_provinsi"`
	Provinsi        string  `json:"provinsi"`
	IdKotaKabupaten string  `json:"_id_kota_kabupaten"`
	KotaKabupaten   string  `json:"kota_kabupaten"`
	IdKecamatan     string  `json:"_id_kecamatan"`
	Kecamatan       string  `json:"kecamatan"`
	IdKelurahan     string  `json:"_id_kelurahan"`
	Kelurahan       string  `json:"kelurahan"`
	IdKodepos       string  `json:"_id_kode_pos"`
	Kodepos         string  `json:"kode_pos"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Address         string  `json:"address"`
}

type CreateAddressResponse struct {
	Id string `json:"_id"`
	// UserId          string  `json:"user_id"`
	Name            string  `json:"name"`
	IsDefault       bool    `json:"is_default"`
	ReceiverName    string  `json:"receiver_name"`
	Phone           string  `json:"phone"`
	IdProvinsi      string  `json:"_id_provinsi"`
	Provinsi        string  `json:"provinsi"`
	IdKotaKabupaten string  `json:"_id_kota_kabupaten"`
	KotaKabupaten   string  `json:"kota_kabupaten"`
	IdKecamatan     string  `json:"_id_kecamatan"`
	Kecamatan       string  `json:"kecamatan"`
	IdKelurahan     string  `json:"_id_kelurahan"`
	Kelurahan       string  `json:"kelurahan"`
	IdKodepos       string  `json:"_id_kode_pos"`
	Kodepos         string  `json:"kode_pos"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Address         string  `json:"address"`
}

type DeleteAddressRequest struct {
	Id string `json:"_id"`
}

type UpdateIsDefaultRequest struct {
	Id     string `json:"_id"`
	UserId string `json:"user_id"`
}

type ApiFirstResponse struct {
	Status string `json:"status"`
}

type ApiResponse struct {
	ErrorMessage string         `json:"error_message"`
	Result       []ResultMapAPI `json:"results"`
	Status       string         `json:"status"`
}

type ResultMapAPI struct {
	Address_components []AddressComponent `json:"address_components"`
	Formatted_address  string             `json:"formatted_address"`
	GeometryStruct     GeometryStruct     `json:"geometry"`
	Place_id           string             `json:"place_id"`
	Types              []string           `json:"types"`
}

type AddressComponent struct {
	Long_name  string   `json:"long_name"`
	Short_name string   `json:"short_name"`
	Types      []string `json:"types"`
}

type GeometryStruct struct {
	Bounds        BoundsStruct   `json:"bounds"`
	Location      LocationStruct `json:"location"`
	Location_type string         `json:"location_type"`
	Viewport      ViewPortStruct `json:"viewport"`
}

type BoundsStruct struct {
	Northeast NortheastBounds `json:"northeast"`
	Southwest SouthwestBounds `json:"southwest"`
}

type NortheastBounds struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type SouthwestBounds struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type ViewPortStruct struct {
	northeast NortheastViewPort
	southwest SouthwestViewPort
}

type NortheastViewPort struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type SouthwestViewPort struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type LocationStruct struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type GetPostcodeRequest struct {
	Id      int32  `json:"_id"`
	Kodepos string `json:"kodepos"`
}

type GetPostcodeResponse struct {
	Id              int32  `json:"_id"`
	IdKelurahan     string `json:"_id_kelurahan"`
	IdProvinsi      string `json:"_id_provinsi"`
	Provinsi        string `json:"provinsi"`
	IdKotaKabupaten string `json:"_id_kota_kabupaten"`
	KotaKabupaten   string `json:"kota_kabupaten"`
	IdKecamatan     string `json:"_id_kecamatan"`
	Kecamatan       string `json:"kecamatan"`
	Kelurahan       string `json:"kelurahan"`
	Kodepos         string `json:"kodepos"`
}
