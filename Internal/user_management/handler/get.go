package userhandler

import (
	userinteractors "BESocialHealth/Internal/user_management/interactors"
	userrepositories "BESocialHealth/Internal/user_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserByIdHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		userRepo := userrepositories.NewUserRepository(db)
		userInteractor := userinteractors.NewUserInteractor(userRepo)
		id := c.Param("id")
		userId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}
		user, err := userInteractor.GetUserById(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})

		}
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}
func GetAllUserHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		userRepo := userrepositories.NewUserRepository(db)
		userInteractor := userinteractors.NewUserInteractor(userRepo)
		users, err := userInteractor.GetAllUsers()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})

		}
		c.JSON(http.StatusOK, gin.H{"data": users})

	}
}
