package middleware

import (
	"BESocialHealth/comon"
	"BESocialHealth/component/appctx"
	"github.com/gin-gonic/gin"
)

func Recover(ac appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")
				if appErr, ok := err.(*comon.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
					return
				}
			}
		}()
		c.Next()
	}
}
