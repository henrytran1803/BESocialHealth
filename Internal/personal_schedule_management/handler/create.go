package schedulehandler

import (
	scheduleinteractors "BESocialHealth/Internal/personal_schedule_management/interactors"
	schedulemodels "BESocialHealth/Internal/personal_schedule_management/models"
	schedulerepositories "BESocialHealth/Internal/personal_schedule_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateScheduleHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		repo := schedulerepositories.NewScheduleRepository(db)
		scheduleInteractor := scheduleinteractors.NewScheduleInteractor(repo)

		var schedule schedulemodels.ScheduleCreateFull
		if err := c.ShouldBindJSON(&schedule); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if err := scheduleInteractor.CreateSchedule(&schedule); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"schedule": schedule})
	}
}
func CreateScheduleDetailHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		repo := schedulerepositories.NewScheduleRepository(db)
		scheduleInteractor := scheduleinteractors.NewScheduleInteractor(repo)

		var scheduleDetail schedulemodels.ScheduleDetailCreate
		if err := c.ShouldBindJSON(&scheduleDetail); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if err := scheduleInteractor.CreateScheduleDetail(&scheduleDetail); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusCreated, gin.H{"schedule": scheduleDetail})
	}
}
