package userrepositories

import (
	usermodels "BESocialHealth/Internal/user_management/models"
)

func (r *UserRepository) GetUserById(id int) (*usermodels.User, error) {
	var user usermodels.User
	if err := r.DB.Table(usermodels.User{}.TableName()).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
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
