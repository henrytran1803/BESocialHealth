package personalcontenthandler

import (
	personalcontentinteractors "BESocialHealth/Internal/personal_content_management/interactors"
	personalcontentmodels "BESocialHealth/Internal/personal_content_management/models"
	personalcontentrepositories "BESocialHealth/Internal/personal_content_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdatePostHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		repo := personalcontentrepositories.NewPersonalContentRepository(db)
		postInteractor := personalcontentinteractors.NewPersonalContentInteractor(repo)

		var post personalcontentmodels.CreatePost

		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		postInteractor.UpdatePostById(&post)
		c.JSON(http.StatusOK, gin.H{"Data": post})
	}
}
