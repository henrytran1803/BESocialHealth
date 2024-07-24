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
	userPhoto.LastName = user.LastName
	userPhoto.Height = user.Height
	userPhoto.Weight = user.Weight
	userPhoto.BDF = user.BDF
	userPhoto.TDEE = user.TDEE
	userPhoto.Calorie = user.Calorie
	userPhoto.Status = user.Status

	return &userPhoto, nil
}

func (r *UserRepository) GetAllUser() (*[]usermodels.UserPhoto, error) {
	var users []usermodels.User
	var userPhotos []usermodels.UserPhoto

	// Lấy tất cả người dùng
	if err := r.DB.Table(usermodels.User{}.TableName()).Find(&users).Error; err != nil {
		return nil, err
	}

	// Duyệt qua tất cả người dùng để lấy ảnh
	for _, user := range users {
		var photo usermodels.Photo
		userPhoto := usermodels.UserPhoto{
			Id:        user.Id,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Height:    user.Height,
			Weight:    user.Weight,
			BDF:       user.BDF,
			TDEE:      user.TDEE,
			Calorie:   user.Calorie,
			Status:    user.Status,
		}

		if err := r.DB.Table("photos").Where("user_id = ?", user.Id).First(&photo).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				userPhoto.Photo = nil
			} else {
				return nil, err
			}
		} else {
			userPhoto.Photo = &photo
		}

		userPhotos = append(userPhotos, userPhoto)
	}

	return &userPhotos, nil
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
