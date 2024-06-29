package usermodels

import "BESocialHealth/comon"

type User struct {
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

type UserDetail struct {
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
}
