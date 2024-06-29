package comon

import "time"

type SQLModel struct {
	Id       int        `json:"id" gorm:"column:id"`
	Status   int        `json:"status" gorm:"column:status"`
	CreateAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdateAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeleteAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
