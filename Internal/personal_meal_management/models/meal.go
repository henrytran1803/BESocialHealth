package mealmodels

import (
	"BESocialHealth/comon"
	"time"
)

type Mealdetail struct {
	Id      int     `json:"id" gorm:"column:id"`
	Serving float64 `json:"serving" gorm:"column:serving"`
}
type Meal struct {
	UserId int          `json:"user_id" gorm:"column:user_id"`
	Dishes []Mealdetail `json:"dishes" gorm:"foreignKey:MealId;references:ID"`
}
type CreateMealDetail struct {
	DishId  int     `json:"dish_id" gorm:"column:dish_id"`
	MealId  int     `json:"meal_id" gorm:"column:meal_id"`
	Serving float64 `json:"serving" gorm:"column:serving"`
}
type CreateMeal struct {
	ID     int `json:"id" gorm:"column:id; primaryKey:auto_increment"`
	UserId int `json:"user_id" gorm:"column:user_id"`
}

func (CreateMeal) TableName() string       { return "meals" }
func (CreateMealDetail) TableName() string { return "meal_detail" }

type GetMeal struct {
	ID           int        `json:"id" gorm:"primaryKey"`
	UserId       int        `json:"user_id" gorm:"column:user_id"`
	Description  string     `json:"description" gorm:"column:description"`
	Date         *time.Time `json:"date" gorm:"column:date"`
	TotalCalorie float64    `json:"total_calorie" gorm:"column:total_calorie"`
	Dishes       []GetDish  `json:"dishes" gorm:"foreignKey:MealId"`
	comon.SQLModel
}

func (GetMeal) TableName() string {
	return "meals"
}

type GetDish struct {
	ID      int     `json:"id" gorm:"primaryKey"`
	DishId  int     `json:"dish_id" gorm:"column:dish_id"`
	MealId  int     `json:"meal_id" gorm:"column:meal_id"`
	Serving float64 `json:"serving" gorm:"column:serving"`
	Calorie float64 `json:"calorie" gorm:"column:calorie"`
	comon.SQLModel
}

func (GetDish) TableName() string {
	return "meal_detail"
}
