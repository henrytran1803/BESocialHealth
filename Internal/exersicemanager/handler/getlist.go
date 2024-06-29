package exersicehandler

import (
	"BESocialHealth/Internal/exersicemanager/interactors"
	exersicerepositories "BESocialHealth/Internal/exersicemanager/repositories"
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
