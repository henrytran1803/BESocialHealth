package mealhandler

import (
	mealinteractors "BESocialHealth/Internal/personal_meal_management/interactors"
	mealmodels "BESocialHealth/Internal/personal_meal_management/models"
	mealrepositories "BESocialHealth/Internal/personal_meal_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateMealDetail(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		mealID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		var meal mealmodels.CreateMealDetail
		if err := c.ShouldBindJSON(&meal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		db := appctx.GetMainDBConnection()
		repo := mealrepositories.NewMealRepository(db)
		mealInteractor := mealinteractors.NewMealInteractor(repo)

		if err := mealInteractor.UpdateMealDetail(mealID, &meal); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"data": meal})
	}
}
