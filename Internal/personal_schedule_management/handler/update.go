package schedulehandler

import (
	scheduleinteractors "BESocialHealth/Internal/personal_schedule_management/interactors"
	schedulemodels "BESocialHealth/Internal/personal_schedule_management/models"
	schedulerepositories "BESocialHealth/Internal/personal_schedule_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateScheduleHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		repo := schedulerepositories.NewScheduleRepository(db)
		scheduleInteractor := scheduleinteractors.NewScheduleInteractor(repo)

		var schedule schedulemodels.ScheduleCreate
		if err := c.ShouldBindJSON(&schedule); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if err := scheduleInteractor.UpdateSchedule(&schedule); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Schedule updated"})
	}
}
func UpdateScheduleDetailHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		repo := schedulerepositories.NewScheduleRepository(db)
		scheduleInteractor := scheduleinteractors.NewScheduleInteractor(repo)

		var schedule schedulemodels.ScheduleDetailCreate
		if err := c.ShouldBindJSON(&schedule); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if err := scheduleInteractor.UpdateScheduleDetail(&schedule); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Scheduledetail updated"})
	}
}
