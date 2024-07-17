package mealhandler

import (
	mealinteractors "BESocialHealth/Internal/personal_meal_management/interactors"
	mealmodels "BESocialHealth/Internal/personal_meal_management/models"
	mealrepositories "BESocialHealth/Internal/personal_meal_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateMealHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var meal mealmodels.Meal
		if err := c.ShouldBind(&meal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		db := appctx.GetMainDBConnection()
		mealRepo := mealrepositories.NewMealRepository(db)
		mealInteractor := mealinteractors.NewMealInteractor(mealRepo)

		id, err := mealInteractor.CreateMeal(&meal)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"id": id})

	}
}
func CreateMealDetailHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var mealdetail mealmodels.CreateMealDetail
		if err := c.ShouldBind(&mealdetail); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		db := appctx.GetMainDBConnection()
		mealRepo := mealrepositories.NewMealRepository(db)
		mealInteractor := mealinteractors.NewMealInteractor(mealRepo)
		if err := mealInteractor.CreateMealDetail(&mealdetail); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"data": mealdetail})
	}
}
