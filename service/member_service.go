package service

import (
	"member-service/model"
)

type MemberService interface {
	// List() (responses []model.GetMemberResponse)
	FindAllMember(request model.GetMemberFindRequest, params ...string) (response []model.GetMemberResponse)
	FindMember(request model.GetFindUserRequest) (response model.GetFindUserResponse)
	Create(request model.CreateMemberRequest, params ...string) (response model.CreateMemberResponse)
	CreateConsumer(msg []byte, queueName string)
	CheckUser(request string, check string) (response model.GetMemberResponse)
	DeletePictureProfile(request model.DeleteProfilePictureRequest, params ...string) (response string)
}
