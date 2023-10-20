package data

import (
	"sekolahApp/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `gorm:"column:name;not null"`
	Email        string `gorm:"column:email;not null;unique"`
	PhoneNumber  string `gorm:"column:phone_number;unique"`
	Password     string `gorm:"column:password;not null;default:'qwerty'"`
	Address      string `gorm:"column:address"`
	ProfilePhoto string `gorm:"column:profile_photo"`
	Gender       string `gorm:"column:gender"`
	RoleID       uint   `gorm:"column:role_id"`
}

func UserCoreToModel(input users.UserCore) User {
	var userModel = User{
		Model:        gorm.Model{},
		Name:         input.Name,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		Password:     input.Password,
		Address:      input.Address,
		ProfilePhoto: input.ProfilePhoto,
		Gender:       input.Gender,
		RoleID:       input.RoleID,
	}
	return userModel
}

func UserModelToCore(input User) users.UserCore {
	var userCore = users.UserCore{
		ID:           input.ID,
		Name:         input.Name,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		Password:     input.Password,
		Address:      input.Address,
		ProfilePhoto: input.ProfilePhoto,
		Gender:       input.Gender,
		RoleID:       input.RoleID,
	}
	return userCore
}
