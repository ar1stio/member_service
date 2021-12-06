package service

import (
	"encoding/json"
	"member-service/config"
	"member-service/entity"
	"member-service/exception"
	"member-service/logger"
	"member-service/model"
	"member-service/producer"
	"member-service/repository"
	"member-service/util"
	"member-service/validation"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewMemberService(memberRepository *repository.MemberRepository, registerProducer *producer.RegisterProducer, config config.Config, bucket *oss.Bucket) MemberService {
	return &memberServiceImpl{
		MemberRepository: *memberRepository,
		RegisterProducer: *registerProducer,
		Config:           config,
		Bucketprofile:    bucket,
	}
}

type memberServiceImpl struct {
	MemberRepository repository.MemberRepository
	RegisterProducer producer.RegisterProducer
	Config           config.Config
	Bucketprofile    *oss.Bucket
}

func (service *memberServiceImpl) Create(request model.CreateMemberRequest, params ...string) (response model.CreateMemberResponse) {
	validation.Validate(request)
	var reqMember = model.GetFindUserRequest{
		UserId: request.Id,
	}

	resultUsername := service.MemberRepository.FindMember(reqMember)

	// if resultUsername.Id != "" {
	// 	exception.PanicIfNeeded("data sudah ada")
	// }
	domain := service.Config.Get("OSS_BUCKET_DOMAIN")
	dir := service.Config.Get("FILE_DIR")
	act := service.Config.Get("ACT")

	if resultUsername.ObjectPicture != "" && resultUsername.Id != "" && request.ProfilePicture != "" {
		util.DeleteImage(service.Bucketprofile, resultUsername.Id, resultUsername.ObjectPicture)
	}

	webImg := model.UploadResponse{}
	if request.ProfilePicture != "" {
		// webImg = util.ObjectStorageService(request.Id, request.ProfilePicture, buck, endpoint, keyId, keySecret, domain, dir, act)
		webImg = util.ObjectStorageService(service.Bucketprofile, request.Id, request.ProfilePicture, domain, dir, act)
	}

	nowInt := time.Now().UnixNano() / int64(time.Millisecond)
	loc, _ := time.LoadLocation(service.Config.Get("TIME_ZONE"))
	server := time.Now().Local()
	local := time.Now().In(loc).String()

	created := model.Created{
		By:         request.Id,
		At:         nowInt,
		DateServer: server,
		DateLocal:  local,
	}

	modified := model.Modified{
		By:         request.Id,
		At:         nowInt,
		DateServer: server,
		DateLocal:  local,
	}

	domainUrl := ""
	if webImg.Domain != "" {
		domainUrl = webImg.Domain + "/" + webImg.Object
	}

	memberinput := entity.Member{
		Id:             request.Id,
		FullName:       request.FullName,
		BirthDate:      request.BirthDate,
		Gender:         request.Gender,
		ProfilePicture: domainUrl,
		ObjectPicture:  webImg.Object,
		Created:        created,
		Modified:       modified,
	}

	if resultUsername.Id == "" {
		service.MemberRepository.Create(memberinput)
	} else {
		service.MemberRepository.Update(memberinput)
	}

	response = model.CreateMemberResponse{
		// Id:             memberinput.Id,
		FullName:       request.FullName,
		BirthDate:      request.BirthDate,
		Gender:         request.Gender,
		ProfilePicture: domainUrl,
		ObjectPicture:  webImg.Object,
	}

	zlog := logger.NewLogger()
	zlog.Info().
		Str("endpoint", params[0]).
		Int16("httpResCode", 200).
		// RawJSON("req", util.JsonStruct(request)).
		RawJSON("req", util.JsonStruct("cek")).
		// RawJSON("res", util.JsonStruct(response)).
		RawJSON("res", util.JsonStruct("cek")).
		Msg("new member created")

	return response
}

func (service *memberServiceImpl) CreateConsumer(msg []byte, queueName string) {

	var message map[string]interface{}
	err := json.Unmarshal(msg, &message)
	exception.PanicIfNeeded(err)

	userId := message["user_id"].(string)
	var reqMember = model.GetFindUserRequest{
		UserId: userId,
	}

	resultUsername := service.MemberRepository.FindMember(reqMember)
	if resultUsername.Id == "" {
		nowInt := time.Now().UnixNano() / int64(time.Millisecond)
		loc, _ := time.LoadLocation(service.Config.Get("TIME_ZONE"))
		server := time.Now().Local()
		local := time.Now().In(loc).String()

		created := model.Created{
			By:         userId,
			At:         nowInt,
			DateServer: server,
			DateLocal:  local,
		}
		modified := model.Modified{
			By:         userId,
			At:         nowInt,
			DateServer: server,
			DateLocal:  local,
		}
		memberinput := entity.Member{
			Id:             userId,
			FullName:       "",
			BirthDate:      "",
			Gender:         "",
			ProfilePicture: "",
			ObjectPicture:  "",
			Created:        created,
			Modified:       modified,
		}
		service.MemberRepository.Create(memberinput)
	}

}

func (service *memberServiceImpl) DeletePictureProfile(request model.DeleteProfilePictureRequest, params ...string) (response string) {
	// validation.Validate(request)
	var reqMember = model.GetFindUserRequest{
		UserId: request.UserId,
	}

	resultUsername := service.MemberRepository.FindMember(reqMember)
	if resultUsername.Id == "" || resultUsername.ObjectPicture == "" {
		exception.PanicIfNeeded("Maaf, anda belum memiliki foto")
	}

	deleteImg := ""
	// webImg = util.ObjectStorageService(request.Id, request.ProfilePicture, buck, endpoint, keyId, keySecret, domain, dir, act)
	deleteImg = util.DeleteImage(service.Bucketprofile, resultUsername.Id, resultUsername.ObjectPicture)

	reqDeleteField := entity.Member{
		Id: resultUsername.Id,
	}
	response = deleteImg
	service.MemberRepository.DeleteFieldPicture(reqDeleteField)

	zlog := logger.NewLogger()
	zlog.Info().
		Str("endpoint", params[0]).
		Int16("httpResCode", 200).
		RawJSON("req", util.JsonStruct(request)).
		RawJSON("res", util.JsonStruct(response)).
		Msg("Picture Profile Deleted")

	return deleteImg
}

func (service *memberServiceImpl) CheckUser(request string, check string) (response model.GetMemberResponse) {
	// validation.ValidateLogin(request)

	response = service.MemberRepository.CheckUser(request, check)

	return response
}

func (service *memberServiceImpl) FindMember(request model.GetFindUserRequest) (response model.GetFindUserResponse) {
	// validation.ValidateLogin(request)

	response = service.MemberRepository.FindMember(request)

	return response
}

func (service *memberServiceImpl) FindAllMember(request model.GetMemberFindRequest, params ...string) (response []model.GetMemberResponse) {

	response = service.MemberRepository.FindAllMember(request)
	if response == nil {
		response = []model.GetMemberResponse{}
	}
	return response
}
