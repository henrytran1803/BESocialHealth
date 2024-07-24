package userhandler

import (
	userinteractors "BESocialHealth/Internal/user_management/interactors"
	usermodels "BESocialHealth/Internal/user_management/models"
	userrepositories "BESocialHealth/Internal/user_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UpdateUserHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		userRepo := userrepositories.NewUserRepository(db)
		userInteractor := userinteractors.NewUserInteractor(userRepo)
		var user usermodels.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		if user.Email == "" || user.FirstName == "" || user.LastName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"err": "Required fields are empty"})
			return
		}

		log.Printf("UserUpdate(id: %d, email: %s, firstname: %s, lastname: %s, role: %d, height: %f, weight: %f, bdf: %f, tdee: %f, calorie: %f, status: %d)",
			user.Id, user.Email, user.FirstName, user.LastName, user.Role, user.Height, user.Weight, user.BDF, user.TDEE, user.Calorie, user.Status)
		user.Role = 1
		if err := userInteractor.UpdateUser(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}
