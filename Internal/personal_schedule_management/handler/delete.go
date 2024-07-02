package schedulehandler

import (
	scheduleinteractors "BESocialHealth/Internal/personal_schedule_management/interactors"
	schedulerepositories "BESocialHealth/Internal/personal_schedule_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteScheduleDetailHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := appctx.GetMainDBConnection()
		repo := schedulerepositories.NewScheduleRepository(db)
		scheduleInteractor := scheduleinteractors.NewScheduleInteractor(repo)

		if err := scheduleInteractor.DeleteScheduleDetailById(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "successfully deleted schedule"})
	}

}
func DeleteScheduleHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := appctx.GetMainDBConnection()
		repo := schedulerepositories.NewScheduleRepository(db)
		scheduleInteractor := scheduleinteractors.NewScheduleInteractor(repo)

		if err := scheduleInteractor.DeleteScheduleById(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "successfully deleted schedule"})
	}

}
