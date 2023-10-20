package service

import (
	"mime/multipart"
	"sekolahApp/app/middlewares"
	"sekolahApp/features/users"
)

type UserService struct {
	userData users.UserDataInterface
}

func NewUsersLogic(repo users.UserDataInterface) users.UserServiceInterface {
	return &UserService{
		userData: repo,
	}
}

func (service *UserService) LoginService(email string, password string) (dataLogin users.UserCore, token string, err error) {
	dataLogin, err = service.userData.LoginQuery(email, password)
	if err != nil {
		return users.UserCore{}, "", err
	}
	token, err = middlewares.CreateToken(dataLogin.ID, dataLogin.RoleID)
	if err != nil {
		return users.UserCore{}, "", err
	}
	return dataLogin, token, nil
}

func (service *UserService) GetAll(page uint, userPerPage uint, searchName string) ([]users.UserCore, bool, error) {
	result, count, err := service.userData.ReadAll(page, userPerPage, searchName)

	next := true
	var pages int64
	if userPerPage != 0 {
		pages = count / int64(userPerPage)
		if count%int64(userPerPage) != 0 {
			pages += 1
		}
		if page == uint(pages) {
			next = false
		}
	}

	return result, next, err
}

func (service *UserService) Create(input users.UserCore, file multipart.File, filename string) error {
	return service.userData.Insert(input, file, filename)
}

func (service *UserService) Delete(id uint) error {
	return service.userData.Delete(id)
}

func (service *UserService) GetById(id uint) (users.UserCore, error) {
	return service.userData.SelectById(id)
}

func (service *UserService) Get(ID uint) ([]users.UserCore, error) {
	return service.userData.GetProfile(ID)
}

func (service *UserService) Update(id uint, input users.UserCore, file multipart.File, filename string) error {
	return service.userData.UpdateProfile(id, input, file, filename)
}
