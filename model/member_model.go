package model

import "time"

type CreateMemberRequest struct {
	Id                 string    `json:"_id"`
	FullName           string    `json:"full_name"`
	BirthDate          string    `json:"birth_date"`
	ProfilePicture     string    `json:"profile_picture"`
	ObjectPicture      string    `json:"object_picture"`
	Gender             string    `json:"gender"`
	CreatedAt          int64     `json:"created_at"`
	CreatedDateLocal   string    `json:"created_date_local"`
	CreatedDateServer  time.Time `json:"created_date_server"`
	ModifiedAt         int64     `json:"modified_at"`
	ModifiedDateLocal  string    `json:"modified_date_local"`
	ModifiedDateServer time.Time `json:"modified_date_server"`
}

type Created struct {
	By         string
	At         int64
	DateServer time.Time
	DateLocal  string
}

type Modified struct {
	By         string
	At         int64
	DateServer time.Time
	DateLocal  string
}

type CreateMemberResponse struct {
	// Id                 string    `json:"_id"`
	FullName           string    `json:"full_name"`
	BirthDate          string    `json:"birth_date"`
	ProfilePicture     string    `json:"profile_picture"`
	ObjectPicture      string    `json:"object_picture"`
	Gender             string    `json:"gender"`
	CreatedAt          int64     `json:"created_at"`
	CreatedDateLocal   string    `json:"created_date_local"`
	CreatedDateServer  time.Time `json:"created_date_server"`
	ModifiedAt         int64     `json:"modified_at"`
	ModifiedDateLocal  string    `json:"modified_date_local"`
	ModifiedDateServer time.Time `json:"modified_date_server"`
}

type GetMemberRequest struct {
	Id             string `json:"_id"`
	FullName       string `json:"full_name"`
	BirthDate      string `json:"birth_date"`
	ProfilePicture string `json:"profile_picture"`
	Gender         string `json:"gender"`
}

type GetMemberFindRequest struct {
	Filter string `json:"filter"`
}

type GetMemberResponse struct {
	Id             string `json:"_id"`
	FullName       string `json:"full_name"`
	BirthDate      string `json:"birth_date"`
	ProfilePicture string `json:"profile_picture"`
	Gender         string `json:"gender"`
}

type GetFindUserRequest struct {
	UserId string `json:"user_id"`
}

type DeleteProfilePictureRequest struct {
	UserId string `json:"user_id"`
}

type UploadResponse struct {
	Object string `json:"object"`
	Domain string `json:"domain"`
}

type GetFindUserResponse struct {
	Id             string `json:"_id"`
	FullName       string `json:"full_name"`
	BirthDate      string `json:"birth_date"`
	ProfilePicture string `json:"profile_picture"`
	ObjectPicture  string `json:"object_picture"`
	Gender         string `json:"gender"`
}

type RegisterUserEvent struct {
	Code             string `json:"code"`
	VerificationCode string `json:"verification_code"`
	UserId           string `json:"user_id"`
	Username         string `json:"user_name"`
	Email            string `json:"email"`
	EmailToken       string `json:"email_token"`
}

type VerificationCodeEvent struct {
	Code             string `json:"code"`
	VerificationCode string `json:"verification_code"`
	UserId           string `json:"user_id"`
	Username         string `json:"user_name"`
	Email            string `json:"email"`
	PhoneToken       string `json:"phone_token"`
	Phone            string `json:"phone"`
	Channel          string `json:"channel"`
	TypeChannel      string `json:"type_channel"`
}

type EditProfileRequest struct {
	Id             string      `json:"_id"`
	FullName       string      `json:"full_name"`
	BirthDate      string      `json:"birth_date"`
	ProfilePicture string      `json:"profile_picture"`
	Gender         string      `json:"gender"`
	Modified       interface{} `json:"modified"`
}
