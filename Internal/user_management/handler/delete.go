package userhandler

import (
	userinteractors "BESocialHealth/Internal/user_management/interactors"
	userrepositories "BESocialHealth/Internal/user_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteUserHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		userRepo := userrepositories.NewUserRepository(db)
		userInteractor := userinteractors.NewUserInteractor(userRepo)
		id := c.Param("id")
		userId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		if err := userInteractor.DeleteUserById(userId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"userId": userId})
	}
}
