package users

import (
	"mime/multipart"
	"time"
)

type UserCore struct {
	ID           uint
	Name         string
	Email        string
	PhoneNumber  string
	Password     string
	Address      string
	ProfilePhoto string
	Gender       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAT    time.Time
	RoleID       uint
}

type UserDataInterface interface {
	LoginQuery(email, password string) (UserCore, error)
	ReadAll(page uint, userPerPage uint, searchName string) ([]UserCore, int64, error)
	Insert(input UserCore, file multipart.File, filename string) error
	Delete(id uint) error
	SelectById(id uint) (UserCore, error)
	GetProfile(ID uint) ([]UserCore, error)
	UpdateProfile(id uint, input UserCore, file multipart.File, filename string) error
}

type UserServiceInterface interface {
	LoginService(email, password string) (UserCore, string, error)
	GetAll(page, userPerPage uint, searchName string) ([]UserCore, bool, error)
	Create(input UserCore, file multipart.File, filename string) error
	Delete(id uint) error
	GetById(id uint) (UserCore, error)
	Get(ID uint) ([]UserCore, error)
	Update(id uint, input UserCore, file multipart.File, filename string) error
}
