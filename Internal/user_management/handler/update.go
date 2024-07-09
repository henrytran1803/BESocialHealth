package userhandler

import (
	userinteractors "BESocialHealth/Internal/user_management/interactors"
	usermodels "BESocialHealth/Internal/user_management/models"
	userrepositories "BESocialHealth/Internal/user_management/repositories"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UpdateUserHandler(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMainDBConnection()
		userRepo := userrepositories.NewUserRepository(db)
		userInteractor := userinteractors.NewUserInteractor(userRepo)

		// Khởi tạo user như một biến không phải con trỏ
		var user usermodels.UserDetail

		// Bind JSON từ request vào user
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		// Kiểm tra giá trị của user và các thuộc tính cần thiết
		if user.Email == "" || user.FirstName == "" || user.LastName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"err": "Required fields are empty"})
			return
		}

		// Debug log để kiểm tra giá trị của user
		log.Printf("UserUpdate(id: %d, email: %s, firstname: %s, lastname: %s, role: %d, height: %f, weight: %f, bdf: %f, tdee: %f, calorie: %f, status: %d)",
			user.Id, user.Email, user.FirstName, user.LastName, user.Role, user.Height, user.Weight, user.BDF, user.TDEE, user.Calorie, user.Status)

		// Gọi phương thức UpdateUser
		if err := userInteractor.UpdateUser(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}
