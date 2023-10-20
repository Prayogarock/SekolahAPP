package handler

import "sekolahApp/features/users"

type UserRequest struct {
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	PhoneNumber  string `json:"phone_number" form:"phone_number"`
	Password     string `json:"password" form:"password"`
	ProfilePhoto string `json:"profile_photo" form:"profile_photo"`
	Gender       string `json:"gender" form:"gender"`
	Address      string `json:"address" form:"address"`
	RoleID       uint   `json:"role_id" form:"role_id"`
}

func UserRequestToCore(input UserRequest) users.UserCore {
	var userCore = users.UserCore{
		Name:         input.Name,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		Password:     input.Password,
		Address:      input.Address,
		Gender:       input.Gender,
		ProfilePhoto: input.ProfilePhoto,
		RoleID:       input.RoleID,
	}
	return userCore
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
