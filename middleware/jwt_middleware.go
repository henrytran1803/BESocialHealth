package middleware

import (
	accountmodels "BESocialHealth/Internal/account/models"
	accountuntils "BESocialHealth/Internal/account/untils"
	"BESocialHealth/comon"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy token từ header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, comon.Response{
				Status:  "fail",
				Message: "Authorization header is required",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, comon.Response{
				Status:  "fail",
				Message: "Token is required",
			})
			c.Abort()
			return
		}

		// Xác thực và phân tích token
		claims, err := accountuntils.VerifyJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, comon.Response{
				Status:  "fail",
				Message: "Invalid token",
			})
			c.Abort()
			return
		}

		// Tìm user trong database
		db := appctx.GetMainDBConnection()
		var user accountmodels.User
		if err := db.Where("id = ?", claims.UserID).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, comon.Response{
				Status:  "fail",
				Message: "User not found",
			})
			c.Abort()
			return
		}

		// Kiểm tra secret
		if user.JWTSecret != claims.Secret {
			c.JSON(http.StatusUnauthorized, comon.Response{
				Status:  "fail",
				Message: "Invalid token secret",
			})
			c.Abort()
			return
		}

		// Lưu thông tin user vào context
		c.Set("user", &user)
		c.Next()
	}
}
