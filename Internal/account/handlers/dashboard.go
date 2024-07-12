package accounthandlers

import (
	accountinteractors "BESocialHealth/Internal/account/interactors"
	accountrepositories "BESocialHealth/Internal/account/repositories"
	"BESocialHealth/component/appctx"
	"BESocialHealth/component/ws"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DashBoardHandler(appctx appctx.AppContext, wsManager *ws.WebSocketManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		accountRepo := accountrepositories.NewDashboardRepositoryy(db, wsManager)
		accountInteractor := accountinteractors.NewDashBoardInteractor(accountRepo)

		dashboard, err := accountInteractor.GetDashboard()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, dashboard)
	}
}
