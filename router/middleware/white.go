package middleware

import (
	"os"

	. "api-gin-web/controller"
	"api-gin-web/utils"
	"api-gin-web/utils/errno"

	"github.com/gin-gonic/gin"
)

// 管理后台访问接口白名单
var whiteIpArr = []string{
	"127.0.0.1",
}

// 推送接口IP白名单
func PushWhite() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 正式环境才有白名单限制
		ip := c.Request.Header.Get("ip")
		if !utils.InArray(ip, whiteIpArr) && os.Getenv("ENV_GO") == "" {
			SendResponse(c, errno.NotWhiteList, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
