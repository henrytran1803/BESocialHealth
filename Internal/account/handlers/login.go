package accounthandlers

import (
	accountinteractors "BESocialHealth/Internal/account/interactors"
	accountmodels "BESocialHealth/Internal/account/models"
	accountrepositories "BESocialHealth/Internal/account/repositories"
	accountuntils "BESocialHealth/Internal/account/untils"
	"BESocialHealth/comon"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		roleRepo := accountrepositories.NewAccountRepository(db)
		accountInteractor := accountinteractors.NewAccountInteractor(roleRepo)

		var login accountmodels.Login
		if err := c.ShouldBind(&login); err != nil {
			c.JSON(http.StatusBadRequest,
				comon.Response{
					Status:  "fail",
					Message: err.Error(),
				},
			)
			return
		}

		user, err := accountInteractor.Login(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest,
				comon.Response{
					Status:  "fail",
					Message: err.Error(),
				})
			return
		}

		user.JWTSecret = accountuntils.GenerateRandomString(32)
		if err := db.Save(user).Error; err != nil {
			c.JSON(http.StatusInternalServerError,
				comon.Response{
					Status:  "fail",
					Message: "Failed to update user secret",
				})
			return
		}

		// Generate JWT token
		token, err := accountuntils.GenerateJWT(user.Id, user.JWTSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				comon.Response{
					Status:  "fail",
					Message: "Failed to generate token",
				})
			return
		}

		c.JSON(http.StatusOK, comon.Response{
			Status:  "ok",
			Message: "success",
			Data:    gin.H{"token": token, "id": user.Id, "role": user.Role},
		})
	}
}
