package userhandler

import (
	userinteractors "BESocialHealth/Internal/user_management/interactors"
	usermodels "BESocialHealth/Internal/user_management/models"
	userrepositories "BESocialHealth/Internal/user_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		userRepo := userrepositories.NewUserRepository(db)
		userInteractor := userinteractors.NewUserInteractor(userRepo)
		var user *usermodels.UserDetail
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
		if err := userInteractor.CreateUser(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}
