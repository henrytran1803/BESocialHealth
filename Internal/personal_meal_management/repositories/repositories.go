package mealrepositories

import "gorm.io/gorm"

type MealRepository struct {
	DB *gorm.DB
}

func NewMealRepository(db *gorm.DB) *MealRepository {
	return &MealRepository{DB: db}
}
