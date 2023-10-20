package handler

import "sekolahApp/features/users"

type UserResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	ProfilePhoto string `json:"profile_photo"`
	Address      string `json:"address"`
	Gender       string `json:"gender"`
	RoleID       uint   `json:"role_id"`
}

func UserCoreToResponseAll(input users.UserCore) UserResponse {
	var userResp = UserResponse{
		ID:           input.ID,
		Name:         input.Name,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		ProfilePhoto: input.ProfilePhoto,
		Address:      input.Address,
		Gender:       input.Gender,
		RoleID:       input.RoleID,
	}
	return userResp
}

type LoginResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
