package accountmodels

import (
	"BESocialHealth/comon"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int     `gorm:"primaryKey; column:id" json:"id"`
	Email     string  `gorm:"type:varchar(255);not null; column:email" json:"email"`
	FirstName string  `gorm:"type:varchar(255); column:fistname" json:"firstname"`
	LastName  string  `gorm:"type:varchar(255); column:lastname" json:"lastname"`
	Role      int     `json:"role"`
	Height    float64 `json:"height"`
	Weight    float64 `json:"weight"`
	BDF       float64 `json:"bdf"`
	TDEE      float64 `json:"tdee"`
	Calorie   float64 `json:"calorie"`
	comon.SQLModel
}

func (User) TableName() string { return "users" }

type CreateUser struct {
	gorm.Model
	Email     string `gorm:"type:varchar(255);not null" json:"email"`
	FirstName string `gorm:"type:varchar(255);column:firstname" json:"firstname"`
	LastName  string `gorm:"type:varchar(255);column:lastname" json:"lastname"`
	Role      int    `gorm:"type:int ;column:role" json:"role"`
}
type CreateAccount struct {
	Email     string `gorm:"type:varchar(255);not null;column:email" json:"email"`
	FirstName string `gorm:"type:varchar(255);column:firstname" json:"firstname"`
	LastName  string `gorm:"type:varchar(255);column:lastname" json:"lastname"`
	Role      int    `gorm:"type:int" json:"role"`
	Password  string `gorm:"type:varchar(255);column:password" json:"password"`
}

type Account struct {
	UserId   int    `gorm:"column:user_id" json:"userid"`
	Password string `gorm:"column:password" json:"password"`
}
type Login struct {
	Email    string `gorm:"type:varchar(255);not null;column:email" json:"email"`
	Password string `gorm:"type:varchar(255);column:password" json:"password"`
}

type PasswordResetToken struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Token     string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (PasswordResetToken) TableName() string { return "password_reset_tokens" }
func (Account) TableName() string            { return "accounts" }
