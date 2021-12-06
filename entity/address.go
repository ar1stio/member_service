package entity

type Address struct {
	Id              string
	UserId          string
	Name            string
	IsDefault       bool
	ReceiverName    string
	Phone           string
	IdProvinsi      string
	Provinsi        string
	IdKotaKabupaten string
	KotaKabupaten   string
	IdKecamatan     string
	Kecamatan       string
	IdKelurahan     string
	Kelurahan       string
	IdKodepos       string
	Kodepos         string
	Latitude        float64
	Longitude       float64
	Address         string
	Created         interface{}
	Modified        interface{}
}
