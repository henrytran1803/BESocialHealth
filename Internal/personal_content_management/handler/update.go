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

func UpdatePostHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		postId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		db := appctx.GetMainDBConnection()
		repo := personalcontentrepositories.NewPersonalContentRepository(db)
		postInteractor := personalcontentinteractors.NewPersonalContentInteractor(repo)

		var post personalcontentmodels.CreatePost

		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		postInteractor.UpdatePostById(postId, &post)
		c.JSON(http.StatusCreated, gin.H{"Data": post})
	}
}
