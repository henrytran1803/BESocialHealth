package reminderrepositories

import "gorm.io/gorm"

type ReminderRepository struct {
	DB *gorm.DB
}

func NewReminderRepository(db *gorm.DB) *ReminderRepository {
	return &ReminderRepository{DB: db}
}
