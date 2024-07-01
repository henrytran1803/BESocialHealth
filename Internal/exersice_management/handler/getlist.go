package exersicehandler

import (
	"BESocialHealth/Internal/exersice_management/interactors"
	exersicerepositories "BESocialHealth/Internal/exersice_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetistExersiceHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		exersiceRepo := exersicerepositories.NewExersiceRepository(db)
		exersiceInteractor := interactors.NewExersiceInteractor(exersiceRepo)
		exersices, err := exersiceInteractor.GetAllExersice()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"data": exersices,
		})
	}
}
