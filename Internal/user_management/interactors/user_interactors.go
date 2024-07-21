package userinteractors

import (
	usermodels "BESocialHealth/Internal/user_management/models"
	userrepositories "BESocialHealth/Internal/user_management/repositories"
	"errors"
)

type UserInteractor struct {
	UserRepository *userrepositories.UserRepository
}

func NewUserInteractor(repo *userrepositories.UserRepository) *UserInteractor {
	return &UserInteractor{
		UserRepository: repo,
	}
}

func (i *UserInteractor) GetUserById(id int) (*usermodels.UserPhoto, error) {
	user, err := i.UserRepository.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (i *UserInteractor) GetAllUsers() (*[]usermodels.User, error) {
	users, err := i.UserRepository.GetAllUser()
	if err != nil {

		return nil, err
	}
	return users, nil
}
func (i *UserInteractor) UpdateUser(user *usermodels.User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	if user.Email == "" || user.FirstName == "" || user.LastName == "" {
		return errors.New("required fields are empty")
	}
	// Các xử lý cập nhật user khác
	if err := i.UserRepository.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (i *UserInteractor) DeleteUserById(id int) error {
	if err := i.UserRepository.DeleteUserById(id); err != nil {
		return err
	}
	return nil
}
func (i *UserInteractor) CreateUser(user *usermodels.UserDetail) error {
	if err := i.UserRepository.CreateUser(user); err != nil {
	}
	return nil
}
