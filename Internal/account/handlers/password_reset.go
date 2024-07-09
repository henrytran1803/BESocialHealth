package accounthandlers

import (
	accountinteractors "BESocialHealth/Internal/account/interactors"
	accountmodels "BESocialHealth/Internal/account/models"
	accountrepositories "BESocialHealth/Internal/account/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestPasswordResetHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Email string `json:"email" binding:"required,email"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		db := appctx.GetMainDBConnection()
		roleRepo := accountrepositories.NewAccountRepository(db)
		accountInteractor := accountinteractors.NewAccountInteractor(roleRepo)
		err := accountInteractor.RequestPasswordReset(request.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Password reset email sent"})
	}
}
func ConfirmPasswordResetHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Token       string `json:"token" binding:"required"`
			NewPassword string `json:"new_password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		db := appctx.GetMainDBConnection()
		roleRepo := accountrepositories.NewAccountRepository(db)
		accountInteractor := accountinteractors.NewAccountInteractor(roleRepo)
		err := accountInteractor.ResetPassword(request.Token, request.NewPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Password has been reset"})
	}
}
func ChangePasswordHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var account accountmodels.Account
		if err := c.ShouldBindJSON(&account); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		}
		db := appctx.GetMainDBConnection()
		roleRepo := accountrepositories.NewAccountRepository(db)
		accountInteractor := accountinteractors.NewAccountInteractor(roleRepo)
		if err := accountInteractor.ChangePass(&account); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Password has been changed"})

	}
}
