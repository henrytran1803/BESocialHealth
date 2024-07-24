package messagemodels

import (
	usermodels "BESocialHealth/Internal/user_management/models"
	"time"
)

type Conversation struct {
	ID           int                    `json:"id" gorm:"column:conversation_id"`
	CreatedAt    time.Time              `json:"created_at"`
	Participants []int                  `json:"participants" gorm:"-"`
	Users        []usermodels.UserPhoto `json:"users" `
}
type ConversationCreate struct {
	ID           int       `json:"id" gorm:"column:conversation_id"`
	CreatedAt    time.Time `json:"created_at"`
	Participants []int     `json:"participants" gorm:"-"`
}
type Message struct {
	ID             int       `json:"id" gorm:"column:message_id"`
	ConversationID int       `json:"conversation_id"`
	SenderID       int       `json:"sender_id"`
	Content        string    `json:"content"`
	Timestamp      time.Time `json:"timestamp"`
	DeletedBy      []int     `json:"deleted_by" gorm:"type:json"`
}

type ConversationParticipant struct {
	ConversationID int `json:"conversation_id"`
	UserID         int `json:"user_id"`
}

func (ConversationParticipant) TableName() string { return "conversationparticipants" }

type GetMessageConvertion struct {
	Users    []usermodels.UserPhoto `json:"users"`
	Messages []Message              `json:"messages"`
}
