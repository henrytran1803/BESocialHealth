package foodrepositories

import "gorm.io/gorm"

type FoodRepository struct {
	DB *gorm.DB
}

func NewFoodRepository(db *gorm.DB) *FoodRepository {
	return &FoodRepository{DB: db}
}
