package controller

import (
	"member-service/model"
	"member-service/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AddressController struct {
	MemberService  service.MemberService
	AddressService service.AddressService
}

func NewAddressController(memberService *service.MemberService, addressService *service.AddressService) AddressController {
	return AddressController{MemberService: *memberService, AddressService: *addressService}
}

func (controller *AddressController) GetClient(c *fiber.Ctx) string {
	return c.Method() + " " + c.OriginalURL() + ", client:" + c.IP()
}

func (controller *AddressController) RouteAddress(app *fiber.App) {
	app.Get("/api/member/provinsi/:id/namaprovinsi/:provinsi?", controller.GetProvinsi)
	app.Get("/api/member/col-provinsi/:id/namaprovinsi/:provinsi?", controller.GetColProvinsi) //done

	app.Get("/api/member/kota-kabupaten/:id/namakotakabupaten/:kotakabupaten?", controller.GetKotaKabupaten)
	app.Get("/api/member/col-kota-kabupaten/:id/namakotakabupaten/:kotakabupaten?", controller.GetColKotaKabupaten) //done

	app.Get("/api/member/kecamatan/:id/namakecamatan/:kecamatan?", controller.GetKecamatan)
	app.Get("/api/member/col-kecamatan/:id/namakecamatan/:kecamatan?", controller.GetColKecamatan) //done

	app.Get("/api/member/kelurahan/:id/namakelurahan/:kelurahan?", controller.GetKelurahan)
	app.Get("/api/member/col-kelurahan/:id/namakelurahan/:kelurahan?", controller.GetColKelurahan) //done

	app.Get("/api/member/kodepos/:id", controller.GetKodepos)
	app.Get("/api/member/all-address/:filter", controller.GetAllAddress)

	app.Get("/api/member/allkodepos/:id", controller.GetKodeposByKode)

	//backoffice
	app.Get("/backoffice-api/member/provinsi/:id/namaprovinsi/:provinsi?", controller.GetProvinsi)
	app.Get("/backoffice-api/member/col-provinsi/:id/namaprovinsi/:provinsi?", controller.GetColProvinsi)

	app.Get("/backoffice-api/member/kota-kabupaten/:id/namakotakabupaten/:kotakabupaten?", controller.GetKotaKabupaten)
	app.Get("/backoffice-api/member/col-kota-kabupaten/:id/namakotakabupaten/:kotakabupaten?", controller.GetColKotaKabupaten)

	app.Get("/backoffice-api/member/kecamatan/:id/namakecamatan/:kecamatan?", controller.GetKecamatan)
	app.Get("/backoffice-api/member/col-kecamatan/:id/namakecamatan/:kecamatan?", controller.GetColKecamatan)

	app.Get("/backoffice-api/member/kelurahan/:id/namakelurahan/:kelurahan?", controller.GetKelurahan)
	app.Get("/backoffice-api/member/col-kelurahan/:id/namakelurahan/:kelurahan?", controller.GetColKelurahan)

	app.Get("/backoffice-api/member/kodepos/:id", controller.GetKodepos)
	app.Get("/backoffice-api/member/all-address/:filter", controller.GetAllAddress)

	app.Get("/backoffice-api/member/allkodepos/:id", controller.GetKodeposByKode)

	// app.Get("/api/auth", controller.List)
}

func (controller *AddressController) GetProvinsi(c *fiber.Ctx) error {
	user_id := c.Params("id")
	// i64, _ := strconv.ParseInt(user_id, 10, 32)
	// i := int(i64)
	provinsi := c.Params("provinsi")

	request := model.GetProvinsiRequest{
		Id:       user_id,
		Provinsi: provinsi,
	}

	response := controller.AddressService.GetProvinsi(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *AddressController) GetColProvinsi(c *fiber.Ctx) error {
	user_id := c.Params("id")
	i64, _ := strconv.ParseInt(user_id, 10, 32)
	i := int(i64)
	provinsi := c.Params("provinsi")

	request := model.GetColProvinsiRequest{
		Id:       i,
		Provinsi: provinsi,
	}

	response := controller.AddressService.GetColProvinsi(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *AddressController) GetKotaKabupaten(c *fiber.Ctx) error {
	user_id := c.Params("id")
	// i64, _ := strconv.ParseInt(user_id, 10, 32)
	// i := int(i64)
	kotakabupaten := c.Params("kotakabupaten")

	request := model.GetKotaKabupatenRequest{
		Id:            user_id,
		KotaKabupaten: kotakabupaten,
	}

	response := controller.AddressService.GetKotaKabupaten(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *AddressController) GetColKotaKabupaten(c *fiber.Ctx) error {
	user_id := c.Params("id")
	// i64, _ := strconv.ParseInt(user_id, 10, 32)
	// i := int(i64)
	kotakabupaten := c.Params("kotakabupaten")

	request := model.GetKotaKabupatenRequest{
		Id:            user_id,
		KotaKabupaten: kotakabupaten,
	}

	response := controller.AddressService.GetColKotaKabupaten(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *AddressController) GetKecamatan(c *fiber.Ctx) error {
	user_id := c.Params("id")
	// i64, _ := strconv.ParseInt(user_id, 10, 32)
	// i := int(i64)
	kecamatan := c.Params("kecamatan")

	request := model.GetKecamatanRequest{
		Id:        user_id,
		Kecamatan: kecamatan,
	}

	response := controller.AddressService.GetKecamatan(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *AddressController) GetColKecamatan(c *fiber.Ctx) error {
	user_id := c.Params("id")
	// i64, _ := strconv.ParseInt(user_id, 10, 32)
	// i := int(i64)
	kecamatan := c.Params("kecamatan")

	request := model.GetKecamatanRequest{
		Id:        user_id,
		Kecamatan: kecamatan,
	}

	response := controller.AddressService.GetColKecamatan(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *AddressController) GetKelurahan(c *fiber.Ctx) error {
	user_id := c.Params("id")
	// i64, _ := strconv.ParseInt(user_id, 10, 32)
	// i := int(i64)
	kelurahan := c.Params("kelurahan")

	request := model.GetKelurahanRequest{
		Id:        user_id,
		Kelurahan: kelurahan,
	}

	response := controller.AddressService.GetKelurahan(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *AddressController) GetColKelurahan(c *fiber.Ctx) error {
	user_id := c.Params("id")
	// i64, _ := strconv.ParseInt(user_id, 10, 32)
	// i := int(i64)
	kelurahan := c.Params("kelurahan")

	request := model.GetKelurahanRequest{
		Id:        user_id,
		Kelurahan: kelurahan,
	}

	response := controller.AddressService.GetColKelurahan(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *AddressController) GetKodepos(c *fiber.Ctx) error {
	user_id := c.Params("id")
	// i64, _ := strconv.ParseInt(user_id, 10, 32)
	// i := int(i64)

	request := model.GetAddressByIdKecamatanRequest{
		Id: user_id,
	}

	response := controller.AddressService.GetKodepos(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *AddressController) GetKodeposByKode(c *fiber.Ctx) error {
	id := c.Params("id")
	// i64, _ := strconv.ParseInt(user_id, 10, 32)
	// i := int32(i64)

	request := model.GetPostcodeRequest{
		Kodepos: id,
	}

	response := controller.AddressService.GetKodeposByKode(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *AddressController) GetAllAddress(c *fiber.Ctx) error {
	filter := c.Params("filter")

	request := model.GetAddressRequest{
		Filter: filter,
	}

	response := controller.AddressService.GetAllAddress(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
