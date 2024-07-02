package messagehandler

import (
	messageinteractors "BESocialHealth/Internal/messaging/interactors"
	messagemodels "BESocialHealth/Internal/messaging/models"
	messagerepositories "BESocialHealth/Internal/messaging/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SendMessageHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var message messagemodels.Message
		if err := c.ShouldBindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db := appctx.GetMainDBConnection()
		messageRepo := messagerepositories.NewMessageRepository(db)
		messageInteractor := messageinteractors.NewMessageInteractor(messageRepo)
		messageID, err := messageInteractor.CreateMessage(message.ConversationID, message.SenderID, message.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message_id": messageID})
	}
}

func ListUserConversationsHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}
		db := appctx.GetMainDBConnection()
		messageRepo := messagerepositories.NewMessageRepository(db)
		messageInteractor := messageinteractors.NewMessageInteractor(messageRepo)
		conversations, err := messageInteractor.GetUserConversations(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, conversations)
	}
}

func ListConversationMessagesHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		conversationID, err := strconv.Atoi(c.Param("conversation_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid conversation ID"})
			return
		}
		db := appctx.GetMainDBConnection()
		messageRepo := messagerepositories.NewMessageRepository(db)
		messageInteractor := messageinteractors.NewMessageInteractor(messageRepo)
		messages, err := messageInteractor.GetConversationMessages(conversationID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, messages)
	}
}
