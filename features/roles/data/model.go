package data

import (
	"sekolahApp/features/users"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name   string `gorm:"column:name;not null"`
	UserID []users.UserCore
}
