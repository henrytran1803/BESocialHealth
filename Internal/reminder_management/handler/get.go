package reminderhandler

import (
	reminderinteractors "BESocialHealth/Internal/reminder_management/interactors"
	reminderrepositories "BESocialHealth/Internal/reminder_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetReminderByIdHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		repo := reminderrepositories.NewReminderRepository(db)
		reminderInteractor := reminderinteractors.NewReminderInteractor(repo)
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{})
		}

		reminder, err := reminderInteractor.GetReminderById(idInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"data": reminder})
	}
}

func GetReminderByIdUserHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		repo := reminderrepositories.NewReminderRepository(db)
		reminderInteractor := reminderinteractors.NewReminderInteractor(repo)
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{})
		}

		reminder, err := reminderInteractor.GetAllReminderById(idInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"data": reminder})
	}
}
