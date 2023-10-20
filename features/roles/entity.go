package roles

import (
	"sekolahApp/features/users"
	"time"
)

type RoleCore struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAT time.Time
	UserID    []users.UserCore
}

type RoleDataInterface interface {
}

type RoleServiceInterface interface {
}
