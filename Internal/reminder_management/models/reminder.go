package remindermodels

import (
	"time"
)

type Reminder struct {
	ID             int       `json:"id" gorm:"primary_key; column:id"`
	UserID         int       `json:"user_id" gorm:"column:user_id""`
	Description    string    `json:"description" gorm:"comment:'description'"`
	ScheduleID     *int      `json:"schedule_id" gorm:"column:schedule_id""`
	MealID         *int      `json:"meal_id" gorm:"not null ;column:meal_id"`
	ReminderTypeID int       `json:"reminder_type_id" gorm:"column:reminder_type" `
	Date           time.Time `json:"date" gorm:"column:date" `
	Status         string    `json:"status" gorm:"column:status" `
}

func (Reminder) TableName() string { return "reminders" }

type ReminderType struct {
	ID   int    `json:"ID" gorm:"primary_key; column:id"`
	Name string `json:"Name" gorm:"type:varchar(255); comment:'name'"`
}
type ReminderCreate struct {
	ID             int    `json:"id" gorm:"primary_key; column:id"`
	UserID         int    `json:"user_id" gorm:"column:user_id""`
	Description    string `json:"description" gorm:"comment:'description'"`
	ScheduleID     *int   `json:"schedule_id" gorm:"column:schedule_id""`
	MealID         *int   `json:"meal_id" gorm:"not null ;column:meal_id"`
	ReminderTypeID int    `json:"reminder_type_id" gorm:"column:reminder_type" `
	Date           string `json:"date" gorm:"column:date" `
	Status         string `json:"status" gorm:"column:status" `
}
