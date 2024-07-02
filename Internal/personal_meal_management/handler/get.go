package mealhandler

import (
	mealinteractors "BESocialHealth/Internal/personal_meal_management/interactors"
	mealrepositories "BESocialHealth/Internal/personal_meal_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMealsByUserIdHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := appctx.GetMainDBConnection()
		repo := mealrepositories.NewMealRepository(db)
		mealInteractor := mealinteractors.NewMealInteractor(repo)

		meals, err := mealInteractor.GetMealByUserID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"data": meals})
	}
}
func GetMealByIdHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		db := appctx.GetMainDBConnection()
		repo := mealrepositories.NewMealRepository(db)
		mealInteractor := mealinteractors.NewMealInteractor(repo)
		meal, err := mealInteractor.GetMeal(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"data": meal})
	}
}