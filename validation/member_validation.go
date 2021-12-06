package validation

import (
	"member-service/exception"
	"member-service/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func Validate(request model.CreateMemberRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.FullName, validation.Required),
		validation.Field(&request.BirthDate, validation.Required),
		validation.Field(&request.Gender, validation.Required),
		// validation.Field(&request.ProfilePicture, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
