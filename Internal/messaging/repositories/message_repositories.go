package messagerepositories

import (
	messagemodels "BESocialHealth/Internal/messaging/models"
	"time"
)

func (r *MessageRepository) CreateMessage(conversationID, senderID int, content string) (int, error) {
	message := messagemodels.Message{
		ConversationID: conversationID,
		SenderID:       senderID,
		Content:        content,
		Timestamp:      time.Now(),
		DeletedBy:      []int{},
	}
	if err := r.DB.Create(&message).Error; err != nil {
		return 0, err
	}
	return message.ID, nil
}

func (r *MessageRepository) GetConversationMessages(conversationID int) ([]messagemodels.Message, error) {
	var messages []messagemodels.Message
	err := r.DB.Where("conversation_id = ?", conversationID).Find(&messages).Error
	return messages, err
}
func (r *MessageRepository) GetUserConversations(userID int) ([]messagemodels.Conversation, error) {
	var conversations []messagemodels.Conversation

	rows, err := r.DB.Raw(`
		SELECT c.conversation_id, c.created_at, cp.user_id
		FROM conversations c
		JOIN conversationparticipants cp ON c.conversation_id = cp.conversation_id
		WHERE cp.conversation_id IN (
			SELECT conversation_id FROM conversationparticipants WHERE user_id = ?
		)
	`, userID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	conversationMap := make(map[int]*messagemodels.Conversation)
	for rows.Next() {
		var conversationID int
		var createdAt time.Time
		var participantID int

		err := rows.Scan(&conversationID, &createdAt, &participantID)
		if err != nil {
			return nil, err
		}

		if conversation, exists := conversationMap[conversationID]; exists {
			conversation.Participants = append(conversation.Participants, participantID)
		} else {
			conversationMap[conversationID] = &messagemodels.Conversation{
				ID:           conversationID,
				CreatedAt:    createdAt,
				Participants: []int{participantID},
			}
		}
	}

	for _, conversation := range conversationMap {
		conversations = append(conversations, *conversation)
	}

	return conversations, nil
}

func (r *MessageRepository) CreateConversation(participants []int) (int, error) {
	conversation := messagemodels.Conversation{}
	if err := r.DB.Create(&conversation).Error; err != nil {
		return 0, err
	}

	for _, userID := range participants {
		conversationParticipant := messagemodels.ConversationParticipant{
			ConversationID: conversation.ID,
			UserID:         userID,
		}
		if err := r.DB.Table(messagemodels.ConversationParticipant{}.TableName()).Create(&conversationParticipant).Error; err != nil {
			return 0, err
		}
	}
	return conversation.ID, nil
}

func (r *MessageRepository) ListConversationsByUserID(userID int) ([]messagemodels.Conversation, error) {
	var conversations []messagemodels.Conversation

	rows, err := r.DB.Table("conversations").
		Select("conversations.conversation_id, conversations.created_at, cp.user_id").
		Joins("JOIN conversationparticipants cp ON conversations.conversation_id = cp.conversation_id").
		Where("cp.user_id = ?", userID).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	conversationMap := make(map[int]*messagemodels.Conversation)
	for rows.Next() {
		var conversationID int
		var createdAt time.Time
		var participantID int

		err := rows.Scan(&conversationID, &createdAt, &participantID)
		if err != nil {
			return nil, err
		}

		if conversation, exists := conversationMap[conversationID]; exists {
			conversation.Participants = append(conversation.Participants, participantID)
		} else {
			conversationMap[conversationID] = &messagemodels.Conversation{
				ID:           conversationID,
				CreatedAt:    createdAt,
				Participants: []int{participantID},
			}
		}
	}

	for _, conversation := range conversationMap {
		conversations = append(conversations, *conversation)
	}

	return conversations, nil
}
