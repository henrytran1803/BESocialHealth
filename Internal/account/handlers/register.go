package accounthandlers

import (
	accountinteractors "BESocialHealth/Internal/account/interactors"
	accountmodels "BESocialHealth/Internal/account/models"
	accountrepositories "BESocialHealth/Internal/account/repositories"
	"BESocialHealth/comon"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAccountHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		roleRepo := accountrepositories.NewAccountRepository(db)
		accountInteractor := accountinteractors.NewAccountInteractor(roleRepo)

		var createAccountData accountmodels.CreateAccount
		if err := c.ShouldBindJSON(&createAccountData); err != nil {
			c.JSON(http.StatusBadRequest, comon.Response{
				Status:  "fail",
				Message: err.Error(),
			},
			)
			return
		}

		if err := accountInteractor.CreateAccount(&createAccountData); err != nil {
			c.JSON(http.StatusInternalServerError,
				comon.Response{
					Status:  "fail",
					Message: err.Error(),
				},
			)
			return
		}

		c.JSON(http.StatusOK,
			comon.Response{
				Status:  "success",
				Message: "Account created",
				Data:    createAccountData,
			},
		)
	}
}
