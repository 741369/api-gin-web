package middleware

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		isAccess := false

		reg := regexp.MustCompile(`^http.*lg1024.com[\:0-9]*`)
		if reg.MatchString(origin) {
			isAccess = true
		}

		path := c.Request.URL.Path
		regPath := regexp.MustCompile(`^\/in*|^\/inner*|^\/admin*`)
		if regPath.MatchString(path) {
			isAccess = true
		}

		if isAccess {
			// 核心处理方式
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept,DNT,X-Mx-ReqToken,"+
				"Keep-Alive,User-Agent,If-Modified-Since,Cache-Control,Authorization")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":    0,
				"message": "OK",
				"data":    nil,
			})
		}

		c.Next()
	}
}
