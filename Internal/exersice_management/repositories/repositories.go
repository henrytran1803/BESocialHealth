package exersicerepositories

import "gorm.io/gorm"

type ExersiceRepository struct {
	DB *gorm.DB
}

func NewExersiceRepository(db *gorm.DB) *ExersiceRepository {
	return &ExersiceRepository{DB: db}
}
