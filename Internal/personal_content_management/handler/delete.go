package personalcontenthandler

import (
	personalcontentinteractors "BESocialHealth/Internal/personal_content_management/interactors"
	personalcontentmodels "BESocialHealth/Internal/personal_content_management/models"
	personalcontentrepositories "BESocialHealth/Internal/personal_content_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeletePostHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idconv, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "id must be integer"})

		}
		db := appctx.GetMainDBConnection()
		repo := personalcontentrepositories.NewPersonalContentRepository(db)
		postInteractor := personalcontentinteractors.NewPersonalContentInteractor(repo)

		if err := postInteractor.DeletePostById(idconv); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "post deleted successfully"})
	}
}

func DeleteLikeByUserIdAndPostIdHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var like personalcontentmodels.CreateLike

		if err := c.ShouldBindJSON(&like); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		}
		db := appctx.GetMainDBConnection()
		repo := personalcontentrepositories.NewPersonalContentRepository(db)
		postInteractor := personalcontentinteractors.NewPersonalContentInteractor(repo)
		if err := postInteractor.DeleteLikeByUserIDAndPostId(int(like.UserId), int(like.PostId)); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "post deleted successfully"})
	}
}
