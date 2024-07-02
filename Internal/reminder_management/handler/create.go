package reminderhandler

import (
	reminderinteractors "BESocialHealth/Internal/reminder_management/interactors"
	remindermodels "BESocialHealth/Internal/reminder_management/models"
	reminderrepositories "BESocialHealth/Internal/reminder_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateReminderHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		repo := reminderrepositories.NewReminderRepository(db)
		reminderInteractor := reminderinteractors.NewReminderInteractor(repo)

		var reminder remindermodels.ReminderCreate
		if err := c.ShouldBindJSON(&reminder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if err := reminderInteractor.CreateReminder(&reminder); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"data": reminder})
	}
}
