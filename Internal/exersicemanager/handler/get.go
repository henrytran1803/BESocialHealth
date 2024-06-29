package exersicehandler

import (
	"BESocialHealth/Internal/exersicemanager/interactors"
	exersicerepositories "BESocialHealth/Internal/exersicemanager/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetExersiceByIdHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		exersiceID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}
		db := appctx.GetMainDBConnection()
		exersiceRepo := exersicerepositories.NewExersiceRepository(db)
		exersiceInteractor := interactors.NewExersiceInteractor(exersiceRepo)
		exersice, err := exersiceInteractor.GetExersice(exersiceID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
		c.JSON(http.StatusOK, gin.H{"data": exersice})

	}
}
