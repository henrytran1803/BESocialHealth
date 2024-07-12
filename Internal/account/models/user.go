package accountmodels

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int       `gorm:"primaryKey; column:id" json:"id"`
	Email     string    `gorm:"type:varchar(255);not null; column:email" json:"email"`
	FirstName string    `gorm:"type:varchar(255); column:firstname" json:"firstname"`
	LastName  string    `gorm:"type:varchar(255); column:lastname" json:"lastname"`
	Role      int       `json:"role"`
	Height    float64   `json:"height"`
	Weight    float64   `json:"weight"`
	BDF       float64   `json:"bdf"`
	TDEE      float64   `json:"tdee"`
	Calorie   float64   `json:"calorie"`
	JWTSecret string    `gorm:"column:jwt_secret" json:"jwt_secret"`
	Status    int       `gorm:"default:0; column:status" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
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
type DashBoard struct {
	All_user       int          `json:"all_user"`
	User_disable   int          `json:"user_disable"`
	Count_posts    int          `json:"count_posts"`
	Count_food     int          `json:"count_food"`
	Count_exersice int          `json:"count_exersice"`
	Count_photos   int          `json:"count_photos"`
	Active_user    int          `json:"active_user"`
	List_active    []UserActive `json:"list_active"`
}
type UserActive struct {
	Id_user   string    `gorm:"primaryKey" json:"id_user"`
	LastLogin time.Time `json:"last_login"`
}

func (PasswordResetToken) TableName() string { return "password_reset_tokens" }
func (Account) TableName() string            { return "accounts" }
