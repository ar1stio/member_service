package repository

import (
	"member-service/entity"
	"member-service/model"
)

type MemberRepository interface {
	Create(member entity.Member)
	Update(auths entity.Member)
	UpdateProfile(auths entity.Member)

	FindAll() (member []entity.Member)
	CheckUser(request string, check string) (login model.GetMemberResponse)
	FindMember(request model.GetFindUserRequest) (member model.GetFindUserResponse)
	FindAllMember(request model.GetMemberFindRequest) (response []model.GetMemberResponse)
	DeleteFieldPicture(auths entity.Member)
}
