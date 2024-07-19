package schedulehandler

import (
	scheduleinteractors "BESocialHealth/Internal/personal_schedule_management/interactors"
	schedulerepositories "BESocialHealth/Internal/personal_schedule_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllScheduleHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		repo := schedulerepositories.NewScheduleRepository(db)
		scheduleInteractor := scheduleinteractors.NewScheduleInteractor(repo)
		schedules, err := scheduleInteractor.GetAllSchedule()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}
		c.JSON(http.StatusOK, gin.H{"data": schedules})
	}
}

func GetScheduleByIdHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := appctx.GetMainDBConnection()
		repo := schedulerepositories.NewScheduleRepository(db)
		scheduleInteractor := scheduleinteractors.NewScheduleInteractor(repo)
		schedules, err := scheduleInteractor.GetScheduleById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}
		c.JSON(http.StatusOK, gin.H{"data": schedules})
	}
}
func GetScheduleByUserIdHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := appctx.GetMainDBConnection()
		repo := schedulerepositories.NewScheduleRepository(db)
		scheduleInteractor := scheduleinteractors.NewScheduleInteractor(repo)
		schedules, err := scheduleInteractor.GetListScheduleByUserId(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}
		c.JSON(http.StatusOK, gin.H{"data": schedules})
	}
}

func GetScheduleByDateHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		date := c.Param("date")
		db := appctx.GetMainDBConnection()
		repo := schedulerepositories.NewScheduleRepository(db)
		scheduleInteractor := scheduleinteractors.NewScheduleInteractor(repo)
		schedules, err := scheduleInteractor.GetScheduleBydate(&id, &date)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}
		c.JSON(http.StatusOK, gin.H{"data": schedules})
	}
}

func GetScheduleDateToDateHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		fromdate := c.Param("fromdate")
		date := c.Param("date")
		db := appctx.GetMainDBConnection()
		repo := schedulerepositories.NewScheduleRepository(db)
		scheduleInteractor := scheduleinteractors.NewScheduleInteractor(repo)
		schedules, err := scheduleInteractor.GetScheduleDateToDate(&fromdate, &date, &id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}
		c.JSON(http.StatusOK, gin.H{"data": schedules})
	}
}
