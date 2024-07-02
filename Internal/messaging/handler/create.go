package messagehandler

import (
	messageinteractors "BESocialHealth/Internal/messaging/interactors"
	messagerepositories "BESocialHealth/Internal/messaging/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateConversationHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Participants []int `json:"participants"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db := appctx.GetMainDBConnection()
		messageRepo := messagerepositories.NewMessageRepository(db)
		messageInteractor := messageinteractors.NewMessageInteractor(messageRepo)
		conversationID, err := messageInteractor.CreateConversation(request.Participants)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"conversation_id": conversationID})
	}
}
