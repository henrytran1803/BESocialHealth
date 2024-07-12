package accountinteractors

import (
	accountmodels "BESocialHealth/Internal/account/models"
	accountrepositories "BESocialHealth/Internal/account/repositories"
	accountuntils "BESocialHealth/Internal/account/untils"
	"BESocialHealth/comon"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AccountInteractor struct {
	AccountRepository *accountrepositories.AccountRepository
}

func NewAccountInteractor(repo *accountrepositories.AccountRepository) *AccountInteractor {
	return &AccountInteractor{
		AccountRepository: repo,
	}
}

type DashBoardInteractor struct {
	DashboardRepository *accountrepositories.DashboardRepository
}

func NewDashBoardInteractor(repo *accountrepositories.DashboardRepository) *DashBoardInteractor {
	return &DashBoardInteractor{
		DashboardRepository: repo,
	}
}
func (i *AccountInteractor) CreateAccount(user *accountmodels.CreateAccount) error {
	exists, err := i.AccountRepository.CheckExists(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("Email already exists")
	}

	return i.AccountRepository.CreateAccount(user)
}
func (i *AccountInteractor) Login(login *accountmodels.Login) (*accountmodels.User, error) {
	if login.Email == "" || login.Password == "" {
		return nil, errors.New("email or password is empty")
	}

	user, err := i.AccountRepository.Login(login.Email, login.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (i *AccountInteractor) RequestPasswordReset(email string) error {
	user, err := i.AccountRepository.FindByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	token, err := comon.GenerateResetToken()
	if err != nil {
		return err
	}

	resetToken := accountmodels.PasswordResetToken{
		UserID:    uint(user.Id),
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	if err := i.AccountRepository.SavePasswordResetToken(&resetToken); err != nil {
		return err
	}

	return accountuntils.SendPasswordResetEmail(email, token)
}
func (i *AccountInteractor) ResetPassword(token string, newPassword string) error {
	resetToken, err := i.AccountRepository.FindPasswordResetToken(token)
	if err != nil {
		return errors.New("invalid or expired reset token")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err := i.AccountRepository.UpdatePassword(resetToken.UserID, string(hashedPassword)); err != nil {
		return err
	}

	return nil
}
func (i *AccountInteractor) ChangePass(account *accountmodels.Account) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if err := i.AccountRepository.UpdatePassword(uint(account.UserId), string(hashedPassword)); err != nil {
		return err
	}
	return nil
}

func (i *DashBoardInteractor) GetDashboard() (*accountmodels.DashBoard, error) {
	dashboard, err := i.DashboardRepository.GetDashboardData()
	if err != nil {
		return nil, err
	}
	return dashboard, nil
}
