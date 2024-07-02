package messagemodels

import (
	"time"
)

// CREATE TABLE conversations (
// conversation_id INT AUTO_INCREMENT PRIMARY KEY,
// created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// );
//
// CREATE TABLE conversationparticipants (
// conversation_id INT,
// user_id INT,
// UNIQUE (conversation_id, user_id),
// FOREIGN KEY (conversation_id) REFERENCES conversations(conversation_id),
// FOREIGN KEY (user_id) REFERENCES users(id)
// );
// CREATE TABLE messages (
// message_id INT AUTO_INCREMENT PRIMARY KEY,
// conversation_id INT,
// sender_id INT,
// content TEXT,
// timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// deleted_by JSON,
// FOREIGN KEY (conversation_id) REFERENCES conversations(conversation_id),
// FOREIGN KEY (sender_id) REFERENCES users(id)
// );

type Conversation struct {
	ID           int       `json:"id" gorm:"column:conversation_id"`
	CreatedAt    time.Time `json:"created_at"`
	Participants []int     `json:"participants" gorm:"-"`
}

type Message struct {
	ID             int       `json:"id"`
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
