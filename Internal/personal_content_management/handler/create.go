package personalcontenthandler

import (
	personalcontentinteractors "BESocialHealth/Internal/personal_content_management/interactors"
	personalcontentmodels "BESocialHealth/Internal/personal_content_management/models"
	personalcontentrepositories "BESocialHealth/Internal/personal_content_management/repositories"
	"BESocialHealth/component/appctx"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io"
	"io/ioutil"
	"net/http"
)

func CreatePostHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form personalcontentmodels.CreatePostFull

		// Bind form data to struct
		if err := c.ShouldBindWith(&form, binding.Form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Parse photos
		formFiles, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		files := formFiles.File["photos"]

		for _, file := range files {
			src, err := file.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			defer src.Close()

			buf := bytes.NewBuffer(nil)
			if _, err := io.Copy(buf, src); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			photo := personalcontentmodels.CreatePhoto{
				Photo_type: "image",
				Image:      buf.Bytes(),
				Url:        "",
			}
			form.CreatePhoto = append(form.CreatePhoto, photo)
		}

		db := appctx.GetMainDBConnection()
		repo := personalcontentrepositories.NewPersonalContentRepository(db)
		postInteractor := personalcontentinteractors.NewPersonalContentInteractor(repo)
		if err := postInteractor.CreatePost(&form); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
	}
}
func LikeHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		repo := personalcontentrepositories.NewPersonalContentRepository(db)
		postInteractor := personalcontentinteractors.NewPersonalContentInteractor(repo)
		var like personalcontentmodels.CreateLike
		if err := c.ShouldBind(&like); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if err := postInteractor.CreateLike(&like); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "Like created successfully"})
	}
}
func CreateComentHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		repo := personalcontentrepositories.NewPersonalContentRepository(db)
		postInteractor := personalcontentinteractors.NewPersonalContentInteractor(repo)
		var coment personalcontentmodels.CreateCommentFull
		if err := c.ShouldBindWith(&coment, binding.Form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		file, err := c.FormFile("photo")
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
		coment.CreatePhoto.Image = imageData
		if err := postInteractor.CreateComent(&coment); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "Comment created successfully"})
	}
}
