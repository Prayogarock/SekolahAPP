package data

import (
	"errors"
	"mime/multipart"
	"sekolahApp/features/users"
	"sekolahApp/helpers"

	"gorm.io/gorm"
)

type UserQuery struct {
	db        *gorm.DB
	dataLogin users.UserCore
}

func NewUsersQuery(db *gorm.DB) users.UserDataInterface {
	return &UserQuery{
		db: db,
	}
}

func (repo *UserQuery) LoginQuery(email string, password string) (dataLogin users.UserCore, err error) {

	var data User

	tx := repo.db.Where("email = ? && password = ?", email, password).Find(&data)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return users.UserCore{}, errors.New("no row affected")
	}
	dataLogin = UserModelToCore(data)
	repo.dataLogin = dataLogin
	return dataLogin, nil
}

func (repo *UserQuery) ReadAll(page uint, userPerPage uint, searchName string) ([]users.UserCore, int64, error) {
	var userData []User
	var totalCount int64

	if page == 0 && userPerPage == 0 {
		tx := repo.db

		if searchName != "" {
			tx = tx.Where("name LIKE ? ", "%"+searchName+"%")
		}
		tx.Find(&userData)
	} else {

		offset := int((page - 1) * userPerPage)

		query := repo.db.Offset(offset).Limit(int(userPerPage))

		if searchName != "" {
			query = query.Where("name LIKE ? ", "%"+searchName+"%")
		}

		tx := query.Find(&userData)
		if tx.Error != nil {
			return nil, 0, tx.Error
		}
	}

	var userCore []users.UserCore
	for _, value := range userData {
		userCore = append(userCore, UserModelToCore(value))
	}

	repo.db.Model(&User{}).Count(&totalCount)

	return userCore, totalCount, nil
}

func (repo *UserQuery) Insert(input users.UserCore, file multipart.File, filename string) error {

	var userModel = UserCoreToModel(input)

	if filename == "default.png" {
		userModel.ProfilePhoto = filename
	} else {
		nameGen, errGen := helpers.GenerateName()
		if errGen != nil {
			return errGen
		}
		userModel.ProfilePhoto = nameGen + filename
		errUp := helpers.Uploader.UploadFile(file, userModel.ProfilePhoto)

		if errUp != nil {
			return errUp
		}
	}

	tx := repo.db.Create(&userModel)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *UserQuery) Delete(id uint) error {
	tx := repo.db.Where("id = ?", id).Delete(&User{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
}

func (repo *UserQuery) SelectById(id uint) (users.UserCore, error) {
	var result User
	tx := repo.db.Find(&result, id)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return users.UserCore{}, errors.New("no row affected")
	}

	resultCore := UserModelToCore(result)
	return resultCore, nil
}

func (repo *UserQuery) GetProfile(ID uint) ([]users.UserCore, error) {
	var userData []User

	query := repo.db.Where("id = ?", ID)
	tx := query.Find(&userData)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var userCore []users.UserCore
	for _, value := range userData {
		userCore = append(userCore, users.UserCore{
			ID:           value.ID,
			Name:         value.Name,
			Email:        value.Email,
			PhoneNumber:  value.PhoneNumber,
			Password:     value.Password,
			Address:      value.Address,
			ProfilePhoto: value.ProfilePhoto,
			Gender:       value.Gender,
			RoleID:       value.RoleID,
		})
	}

	return userCore, nil
}

func (repo *UserQuery) UpdateProfile(id uint, input users.UserCore, file multipart.File, filename string) error {
	var user User
	tx := repo.db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("target not found")
	}

	updatedUser := UserCoreToModel(input)

	if filename != "default.png" {
		nameGen, errGen := helpers.GenerateName()
		if errGen != nil {
			return errGen
		}
		updatedUser.ProfilePhoto = nameGen + filename

		errUp := helpers.Uploader.UploadFile(file, updatedUser.ProfilePhoto)
		if errUp != nil {
			return errUp
		}
	}

	tx = repo.db.Model(&user).Updates(updatedUser)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + " failed to update data")
	}
	return nil
}
