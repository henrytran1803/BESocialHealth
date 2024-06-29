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
)

func CreateExersiceHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		exersiceRepo := exersicerepositories.NewExersiceRepository(db)
		exersiceInteractor := interactors.NewExersiceInteractor(exersiceRepo)

		var exersice exersicemodels.CreateExersice
		if err := c.ShouldBindWith(&exersice, binding.Form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if exersice.Name == "" || exersice.Description == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name, Description, and Image are required fields"})
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

		// Gọi interactor để tạo bài tập mới
		if err := exersiceInteractor.CreateExersice(&newEx, imageData, file.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Exersice item created successfully", "exersice": newEx})
	}
}
