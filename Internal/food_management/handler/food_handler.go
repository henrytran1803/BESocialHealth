package foodhandler

import (
	foodinteractors "BESocialHealth/Internal/food_management/interactors"
	foodmodels "BESocialHealth/Internal/food_management/models"
	foodrepositories "BESocialHealth/Internal/food_management/repositories"
	"BESocialHealth/comon"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreateFoodHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var food foodmodels.FoodCreate
		if err := c.ShouldBindWith(&food, binding.Form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
			return
		}
		fileData, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open image file"})
			return
		}
		defer fileData.Close()
		imageData, err := ioutil.ReadAll(fileData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read image file"})
			return
		}
		db := appctx.GetMainDBConnection()
		foodRepo := foodrepositories.NewFoodRepository(db)
		foodInteractor := foodinteractors.NewFoodInteractor(foodRepo)
		newFood := foodmodels.Food{
			Name:        food.Name,
			Description: food.Description,
			Calorie:     food.Calorie,
			Protein:     food.Protein,
			Fat:         food.Fat,
			Carb:        food.Carb,
			Sugar:       food.Sugar,
			Serving:     food.Serving,
		}

		// Call CreateFood interactor
		if err := foodInteractor.CreateFood(&newFood, imageData, file.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Food item created successfully", "food": newFood})
	}
}
func UpdateFoodHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var food foodmodels.FoodCreate
		if err := c.ShouldBind(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
			return
		}

		fileData, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open image file"})
			return
		}
		defer fileData.Close()

		imageData, err := ioutil.ReadAll(fileData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read image file"})
			return
		}

		db := appctx.GetMainDBConnection()
		foodRepo := foodrepositories.NewFoodRepository(db)
		foodInteractor := foodinteractors.NewFoodInteractor(foodRepo)

		newFood := foodmodels.Food{
			Name:        food.Name,
			Description: food.Description,
			Calorie:     food.Calorie,
			Protein:     food.Protein,
			Fat:         food.Fat,
			Carb:        food.Carb,
			Sugar:       food.Sugar,
			Serving:     food.Serving,
		}
		if err := foodInteractor.UpdateFood(&newFood, imageData, file.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK,
			comon.Response{
				Status:  "ok",
				Message: "Food item updated successfully",
				Data:    newFood,
			},
		)
	}
}

func DeleteFoodHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		foodID, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid food ID"})
			return
		}
		db := appctx.GetMainDBConnection()
		foodRepo := foodrepositories.NewFoodRepository(db)
		foodInteractor := foodinteractors.NewFoodInteractor(foodRepo)
		if err := foodInteractor.DeleteFood(foodID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Food item deleted successfully"})
	}
}
func GetListFoodHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		foodRepo := foodrepositories.NewFoodRepository(db)
		foodInteractor := foodinteractors.NewFoodInteractor(foodRepo)
		foods, err := foodInteractor.GetListFood()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}
		c.JSON(http.StatusOK, gin.H{"data": foods})
	}
}
func GetFoodHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		foodID, err := strconv.Atoi(idParam)
		db := appctx.GetMainDBConnection()
		foodRepo := foodrepositories.NewFoodRepository(db)
		foodInteractor := foodinteractors.NewFoodInteractor(foodRepo)
		foods, err := foodInteractor.GetFood(strconv.Itoa(foodID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}
		c.JSON(http.StatusOK, gin.H{"data": foods})
	}
}
