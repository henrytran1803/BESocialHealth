package exersicehandler

import (
	"BESocialHealth/Internal/exersicemanager/interactors"
	exersicemodels "BESocialHealth/Internal/exersicemanager/models"
	exersicerepositories "BESocialHealth/Internal/exersicemanager/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io/ioutil"
	"net/http"
	"strconv"
)

func UpdateExersiceHandeler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		exersiceID, err := strconv.Atoi(idParam)
		db := appctx.GetMainDBConnection()
		exersiceRepo := exersicerepositories.NewExersiceRepository(db)
		exersiceInteractor := interactors.NewExersiceInteractor(exersiceRepo)

		var exersice exersicemodels.CreateExersice
		if err := c.ShouldBindWith(&exersice, binding.Form); err != nil {
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

		newEx := exersicemodels.Exersice{
			Name:          exersice.Name,
			Description:   exersice.Description,
			Calorie:       exersice.Calorie,
			Time_serving:  exersice.Time_serving,
			Rep_serving:   exersice.Rep_serving,
			Exersice_type: exersice.Exersice_type,
		}

		// Call CreateFood interactor
		if err := exersiceInteractor.UpdateExersice(exersiceID, &newEx, imageData, file.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Food item created successfully", "food": newEx})

	}
}
