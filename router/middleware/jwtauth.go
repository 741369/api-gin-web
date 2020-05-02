/**********************************************
** @Des:
** @Author: 1@lg1024.com
** @Last Modified time: 2020/5/3 上午12:31
***********************************************/

package middleware

import (
	"api-gin-web/controller/user"
	"api-gin-web/router/middleware/jwt"
	"time"
)

func AuthInit() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     "secret",
		PayloadFunc:     user.PayloadFunc,
		IdentityHandler: user.IdentityHandler,
		Authenticator:   user.Authenticator,
		Authorizator:    user.Authorizator,
		Unauthorized:    user.Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
}
