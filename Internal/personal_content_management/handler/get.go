package personalcontenthandler

import (
	personalcontentinteractors "BESocialHealth/Internal/personal_content_management/interactors"
	personalcontentrepositories "BESocialHealth/Internal/personal_content_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetPostByIdHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idconv, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		db := appctx.GetMainDBConnection()
		repo := personalcontentrepositories.NewPersonalContentRepository(db)
		postInteractor := personalcontentinteractors.NewPersonalContentInteractor(repo)

		post, err := postInteractor.GetPostById(idconv)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"data": post})

	}
}
func GetAllPostHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appctx.GetMainDBConnection()
		repo := personalcontentrepositories.NewPersonalContentRepository(db)
		postInteractor := personalcontentinteractors.NewPersonalContentInteractor(repo)

		posts, err := postInteractor.GetAllPost()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"data": posts})

	}
}
func GetAllComentByPostIdHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idconv, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		db := appctx.GetMainDBConnection()
		repo := personalcontentrepositories.NewPersonalContentRepository(db)
		postInteractor := personalcontentinteractors.NewPersonalContentInteractor(repo)

		comments, err := postInteractor.GetAllComentByPostId(idconv)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"data": comments})

	}
}
