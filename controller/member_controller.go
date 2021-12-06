package controller

import (
	"member-service/exception"
	"member-service/model"
	"member-service/service"

	"github.com/gofiber/fiber/v2"
)

type MemberController struct {
	MemberService  service.MemberService
	AddressService service.AddressService
}

func NewMemberController(memberService *service.MemberService, addressService *service.AddressService) MemberController {
	return MemberController{MemberService: *memberService, AddressService: *addressService}
}

func (controller *MemberController) Route(app *fiber.App) {
	app.Post("/api/member/update", controller.Create)
	// app.Post("/api/member/find", controller.FindAllMember)
	app.Delete("/api/member/deleted-picture", controller.DeletePicture)

	app.Post("/backoffice-api/member/update", controller.Create)
	app.Post("/backoffice-api/member/find", controller.FindAllMember)
	app.Delete("/backoffice-api/member/deleted-picture/:id", controller.DeletePicture)

	// app.Get("/api/auth", controller.List)
}

func (controller *MemberController) GetClient(c *fiber.Ctx) string {
	return c.Method() + " " + c.OriginalURL() + ", client:" + c.IP()
}

func (controller *MemberController) Create(c *fiber.Ctx) error {
	var request model.CreateMemberRequest
	err := c.BodyParser(&request)
	id := c.Get("x-auth-id")
	request.Id = id

	exception.PanicIfNeeded(err)

	response := controller.MemberService.Create(request, controller.GetClient(c))
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *MemberController) DeletePicture(c *fiber.Ctx) error {
	var request model.DeleteProfilePictureRequest
	// user_id := c.Params("id")
	user_id := c.Get("x-auth-id")

	request = model.DeleteProfilePictureRequest{
		UserId: user_id,
	}

	response := controller.MemberService.DeletePictureProfile(request, controller.GetClient(c))
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *MemberController) FindAllMember(c *fiber.Ctx) error {
	var request model.GetMemberFindRequest
	err := c.BodyParser(&request)

	exception.PanicIfNeeded(err)

	response := controller.MemberService.FindAllMember(request, controller.GetClient(c))
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

// func (controller *MemberController) List(c *fiber.Ctx) error {
// 	responses := controller.MemberService.List()
// 	return c.JSON(model.WebResponse{
// 		Code:   200,
// 		Status: "OK",
// 		Data:   responses,
// 	})
// }
