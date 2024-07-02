package messageinteractors

import (
	messagemodels "BESocialHealth/Internal/messaging/models"
	messagerepositories "BESocialHealth/Internal/messaging/repositories"
)

type MessageInteractor struct {
	MessageRepository *messagerepositories.MessageRepository
}

func NewMessageInteractor(repo *messagerepositories.MessageRepository) *MessageInteractor {
	return &MessageInteractor{
		MessageRepository: repo,
	}
}
func (mi *MessageInteractor) CreateConversation(participants []int) (int, error) {
	return mi.MessageRepository.CreateConversation(participants)
}

func (mi *MessageInteractor) CreateMessage(conversationID, senderID int, content string) (int, error) {
	return mi.MessageRepository.CreateMessage(conversationID, senderID, content)
}

func (mi *MessageInteractor) GetUserConversations(userID int) ([]messagemodels.Conversation, error) {
	return mi.MessageRepository.GetUserConversations(userID)
}

func (mi *MessageInteractor) GetConversationMessages(conversationID int) ([]messagemodels.Message, error) {
	return mi.MessageRepository.GetConversationMessages(conversationID)
}
