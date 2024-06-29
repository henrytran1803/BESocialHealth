package exersicehandler

import (
	"BESocialHealth/Internal/exersicemanager/interactors"
	exersicerepositories "BESocialHealth/Internal/exersicemanager/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteExersiceHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		exersiceID, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid food ID"})
			return
		}
		db := appctx.GetMainDBConnection()
		exersiceRepo := exersicerepositories.NewExersiceRepository(db)
		exersiceInteractor := interactors.NewExersiceInteractor(exersiceRepo)
		if err := exersiceInteractor.DeleteExersice(exersiceID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
}
