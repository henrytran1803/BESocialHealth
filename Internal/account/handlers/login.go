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
		}
		user, err := accountInteractor.Login(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest,
				comon.Response{
					Status:  "fail",
					Message: err.Error(),
				})
		}
		c.JSON(http.StatusOK, comon.Response{
			Status:  "ok",
			Message: "success",
			Data:    user,
		})
	}
}
