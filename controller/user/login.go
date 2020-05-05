/**********************************************
** @Des:
** @Author: 1@lg1024.com
** @Last Modified time: 2020/5/2 下午11:57
***********************************************/

package user

import (
	"api-gin-web/controller"
	"api-gin-web/model"
	"api-gin-web/router/middleware/jwt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mojocn/base64Captcha"

	//"github.com/mssola/user_agent"
	"log"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(model.SysUser)
		r, _ := v["role"].(model.SysRole)
		return jwt.MapClaims{
			jwt.IdentityKey:  u.UserId,
			jwt.RoleIdKey:    r.RoleId,
			jwt.RoleKey:      r.RoleKey,
			jwt.NiceKey:      "user",
			jwt.DataScopeKey: r.DataScope,
			jwt.RoleNameKey:  "role",
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims["identity"],
		"UserName":    claims["nice"],
		"RoleKey":     claims["rolekey"],
		"UserId":      claims["identity"],
		"RoleIds":     claims["roleid"],
		"DataScope":   claims["datascope"],
	}
}

// @Summary 登陆
// @Tags login
// @Description 获取token
// LoginHandler can be used by clients to get a jwt token.
// Payload needs to be json in the form of {"username": "USERNAME", "password": "PASSWORD"}.
// Reply will be of the form {"token": "TOKEN"}.
// @Accept  application/json
// @Product application/json
// @Param username body model.Login  true "Add account"
// @Success 200 {string} string "{"code": 200, "expire": "2019-08-07T12:45:48+08:00", "token": ".eyJleHAiOjE1NjUxNTMxNDgsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU2NTE0OTU0OH0.-zvzHvbg0A" }"
// @Router /login [post]
func Authenticator(c *gin.Context) (interface{}, error) {
	var loginVals model.Login

	if err := c.ShouldBind(&loginVals); err != nil {
		//loginlog.Create()
		return nil, jwt.ErrMissingLoginValues
	}
	if !store.Verify(loginVals.UUID, loginVals.Code, true) {
		return nil, jwt.ErrInvalidVerificationode
	}

	user, role, e := loginVals.GetUser()
	if e == nil {
		return map[string]interface{}{"user": user, "role": role}, nil
	} else {
		log.Println(e.Error())
	}

	return nil, jwt.ErrFailedAuthentication
}

// @Summary 退出登录
// @Description 获取token
// LoginHandler can be used by clients to get a jwt token.
// Reply will be of the form {"token": "TOKEN"}.
// @Accept  application/json
// @Product application/json
// @Success 200 {string} string "{"code": 200, "msg": "成功退出系统" }"
// @Router /logout [post]
// @Security
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "退出成功",
	})

}

// @Summary 创建用户
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body model.SysUser true "用户数据"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/sysuser [post]
func InsertSysUser(c *gin.Context) {
	var sysuser model.SysUser
	err := c.BindWith(&sysuser, binding.JSON)
	if err != nil {
		controller.SendResponse(c, err, nil)
		return
	}

	sysuser.CreateBy = "test"
	id, err := sysuser.InsertUser()
	if err != nil {
		controller.SendResponse(c, err, nil)
		return
	}
	controller.SendResponse(c, nil, id)
}

func Authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(model.SysUser)
		r, _ := v["role"].(model.SysRole)
		c.Set("role", r.RoleName)
		c.Set("roleIds", r.RoleId)
		c.Set("userId", u.UserId)
		c.Set("userName", u.UserName)
		c.Set("dataScope", r.DataScope)
		return true
	}
	return false
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}
