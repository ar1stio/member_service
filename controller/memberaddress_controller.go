package controller

import (
	"member-service/exception"
	"member-service/model"
	"member-service/service"

	"github.com/gofiber/fiber/v2"
)

type MemberAddressController struct {
	MemberService        service.MemberService
	MemberAddressService service.MemberAddressService
	AddressService       service.AddressService
}

func NewMemberAddressController(memberAddressService *service.MemberAddressService, memberService *service.MemberService, addressService *service.AddressService) MemberAddressController {
	return MemberAddressController{MemberAddressService: *memberAddressService, MemberService: *memberService, AddressService: *addressService}
}

func (controller *MemberAddressController) RouteMemberAddress(app *fiber.App) {

	app.Post("/api/member/address", controller.CreateAddress)
	app.Put("/api/member/address", controller.UpdateAddress)
	app.Delete("/api/member/address/:id", controller.DeleteAddress)
	app.Post("/api/member/address-find", controller.FindAddress)
	app.Put("/api/member/address-is-default", controller.UpdateAddressIsDefault)

	app.Post("/backoffice-api/member/address", controller.CreateAddress)
	app.Put("/backoffice-api/member/address", controller.UpdateAddress)
	app.Delete("/backoffice-api/member/address/:id", controller.DeleteAddress)
	app.Post("/backoffice-api/member/address-find", controller.FindAddress)
	app.Put("/backoffice-api/member/address-is-default", controller.UpdateAddressIsDefault)

	// app.Get("/api/auth", controller.List)
}

func (controller *MemberAddressController) GetClient(c *fiber.Ctx) string {
	return c.Method() + " " + c.OriginalURL() + ", client:" + c.IP()
}

func (controller *MemberAddressController) FindAddress(c *fiber.Ctx) error {
	var request model.GetAddressRequest
	err := c.BodyParser(&request)

	exception.PanicIfNeeded(err)

	response := controller.MemberAddressService.FindAddress(request, controller.GetClient(c))
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *MemberAddressController) CreateAddress(c *fiber.Ctx) error {
	var request model.CreateAddressRequest
	err := c.BodyParser(&request)

	id := c.Get("x-auth-id")
	request.UserId = id

	exception.PanicIfNeeded(err)

	response := controller.MemberAddressService.CreateAddress(request, controller.GetClient(c))
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *MemberAddressController) UpdateAddress(c *fiber.Ctx) error {
	var request model.CreateAddressRequest
	err := c.BodyParser(&request)

	id := c.Get("x-auth-id")
	request.UserId = id

	exception.PanicIfNeeded(err)

	response := controller.MemberAddressService.UpdateAddress(request, controller.GetClient(c))
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *MemberAddressController) UpdateAddressIsDefault(c *fiber.Ctx) error {
	var request model.UpdateIsDefaultRequest
	err := c.BodyParser(&request)

	id := c.Get("x-auth-id")
	request.UserId = id
	exception.PanicIfNeeded(err)

	response := controller.MemberAddressService.UpdateIsDefaultAddress(request, controller.GetClient(c))
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *MemberAddressController) DeleteAddress(c *fiber.Ctx) error {
	var request model.DeleteAddressRequest
	id := c.Params("id")

	request = model.DeleteAddressRequest{
		Id: id,
	}

	response := controller.MemberAddressService.DeleteAddress(request, controller.GetClient(c))
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
