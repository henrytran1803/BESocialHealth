package middleware

import (
	accountrepositories "BESocialHealth/Internal/account/repositories"
	"BESocialHealth/component/appctx"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2/jwt"
	"net/http"
	"strings"
)

var jwtSecret = []byte("your_global_secret_key")

func AuthMiddleware(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		userID := claims["user_id"].(string)
		secret := claims["secret"].(string)

		db := appctx.GetMainDBConnection()
		roleRepo := accountrepositories.NewAccountRepository(db)
		user, err := roleRepo.FindByID(userID)
		if err != nil || user.JWTSecret != secret {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
