package middleware

import (
	"api-gin-web/controller"
	"api-gin-web/utils/errno"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			controller.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
