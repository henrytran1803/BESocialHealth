package accountrepositories

import (
	accountmodels "BESocialHealth/Internal/account/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (r *AccountRepository) Create(user *accountmodels.User) error {
	return r.DB.Table(accountmodels.User{}.TableName()).Create(user).Error
}
func (r *AccountRepository) FindByEmail(username string) (*accountmodels.User, error) {
	var user accountmodels.User
	if err := r.DB.Table(accountmodels.User{}.TableName()).Where("email = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AccountRepository) CheckExists(username string) (bool, error) {
	var user accountmodels.User
	if err := r.DB.Table(accountmodels.User{}.TableName()).Where("email = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (r *AccountRepository) Update(user *accountmodels.User) error {
	return r.DB.Table(accountmodels.User{}.TableName()).Updates(user).Error
}
func (r *AccountRepository) Delete(username string) error {
	var user accountmodels.User
	if err := r.DB.Table(accountmodels.User{}.TableName()).Where("name = ?", username).First(&user).Error; err != nil {
		return err
	}
	if err := r.DB.Table(accountmodels.User{}.TableName()).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *AccountRepository) FindByID(id uint) (*accountmodels.User, error) {
	var user accountmodels.User
	if err := r.DB.Table(accountmodels.User{}.TableName()).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *AccountRepository) CreateAccount(user *accountmodels.CreateAccount) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		createUser := accountmodels.CreateUser{
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      1,
		}
		if err := tx.Table(accountmodels.User{}.TableName()).Create(&createUser).Error; err != nil {
			return err
		}

		createAccount := accountmodels.Account{
			UserId:   int(createUser.ID),
			Password: string(hashedPassword), // Lưu mật khẩu đã mã hóa vào cơ sở dữ liệu
		}
		if err := tx.Table(accountmodels.Account{}.TableName()).Create(&createAccount).Error; err != nil {
			return err
		}

		return nil
	})
}
func (r *AccountRepository) Login(username string, password string) (*accountmodels.User, error) {
	var user accountmodels.User
	if err := r.DB.Table(accountmodels.User{}.TableName()).Where("email = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	var account accountmodels.Account
	if err := r.DB.Table(accountmodels.Account{}.TableName()).Where("user_id = ?", user.Id).First(&account).Error; err != nil {
	}
	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (r *AccountRepository) SavePasswordResetToken(token *accountmodels.PasswordResetToken) error {
	return r.DB.Table(accountmodels.PasswordResetToken{}.TableName()).Create(token).Error
}
func (r *AccountRepository) FindPasswordResetToken(token string) (*accountmodels.PasswordResetToken, error) {
	var resetToken accountmodels.PasswordResetToken
	if err := r.DB.Table(accountmodels.PasswordResetToken{}.TableName()).Where("token = ?", token).First(&resetToken).Error; err != nil {
		return nil, err
	}
	return &resetToken, nil
}
func (r *AccountRepository) UpdatePassword(userID uint, hashedPassword string) error {
	return r.DB.Table(accountmodels.Account{}.TableName()).Where("id = ?", userID).Update("password", hashedPassword).Error
}
func (r *AccountRepository) DeletePasswordResetToken(token string) error {
	return r.DB.Table(accountmodels.PasswordResetToken{}.TableName()).Where("token = ?", token).Delete(&accountmodels.PasswordResetToken{}).Error
}

func (r *DashboardRepository) GetDashboardData() (*accountmodels.DashBoard, error) {
	var allUsersCount, disabledUsersCount, postsCount, foodCount, exercisesCount, photosCount int64

	if err := r.DB.Table("users").Count(&allUsersCount).Error; err != nil {
		return nil, err
	}
	if err := r.DB.Table("users").Where("status = ?", 1).Count(&disabledUsersCount).Error; err != nil {
		return nil, err
	}
	if err := r.DB.Table("posts").Count(&postsCount).Error; err != nil {
		return nil, err
	}
	if err := r.DB.Table("dishes").Count(&foodCount).Error; err != nil {
		return nil, err
	}
	if err := r.DB.Table("exersices").Count(&exercisesCount).Error; err != nil {
		return nil, err
	}
	if err := r.DB.Table("photos").Count(&photosCount).Error; err != nil {
		return nil, err
	}

	activeUsers := r.Manager.GetActiveUsers()
	activeUserCount := len(activeUsers)

	return &accountmodels.DashBoard{
		All_user:       int(allUsersCount),
		User_disable:   int(disabledUsersCount),
		Count_posts:    int(postsCount),
		Count_food:     int(foodCount),
		Count_exersice: int(exercisesCount),
		Count_photos:   int(photosCount),
		Active_user:    activeUserCount,
		List_active:    activeUsers,
	}, nil
}
