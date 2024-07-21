package userrepositories

import (
	usermodels "BESocialHealth/Internal/user_management/models"
	"gorm.io/gorm"
)

func (r *UserRepository) GetUserById(id int) (*usermodels.UserPhoto, error) {
	var user usermodels.User
	var userPhoto usermodels.UserPhoto
	var photo usermodels.Photo
	if err := r.DB.Table(usermodels.User{}.TableName()).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Table("photos").Where("user_id = ?", user.Id).First(&photo).Error; err != nil {
		if err := gorm.ErrRecordNotFound; err == gorm.ErrRecordNotFound {
			userPhoto.Photo = nil
		} else {
			return nil, err
		}
	} else {
		userPhoto.Photo = &photo
	}

	// Populate userPhoto fields from user
	userPhoto.Id = user.Id
	userPhoto.Email = user.Email
	userPhoto.FirstName = user.FirstName
	userPhoto.Height = user.Height
	userPhoto.Weight = user.Weight
	userPhoto.BDF = user.BDF
	userPhoto.TDEE = user.TDEE
	userPhoto.Calorie = user.Calorie
	userPhoto.Status = user.Status

	return &userPhoto, nil
}

func (r *UserRepository) GetAllUser() (*[]usermodels.User, error) {

	var users []usermodels.User
	if err := r.DB.Table(usermodels.User{}.TableName()).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *UserRepository) UpdateUser(user *usermodels.User) error {
	if err := r.DB.Table(usermodels.User{}.TableName()).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUserById(id int) error {
	if err := r.DB.Table(usermodels.User{}.TableName()).Where("id = ?", id).Update("status", 1).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) CreateUser(user *usermodels.UserDetail) error {
	if err := r.DB.Table(usermodels.User{}.TableName()).Create(&user).Error; err != nil {
		return err
	}
	return nil
}
