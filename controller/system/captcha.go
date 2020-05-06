/**********************************************
** @Des:
** @Author: 1@lg1024.com
** @Last Modified time: 2020/5/5 下午10:59
***********************************************/

package system

import (
	"api-gin-web/controller"
	"api-gin-web/model"
	"github.com/gin-gonic/gin"
)

// @Summary 获取图片验证码
// @Tags 工具 / system
// @Description 获取图片验证码
// @Accept  application/json
// @Product application/json
// @Success 200 {string} string "{"data": "图片base64", "id":"图片唯一id" }"
// @Router /api/v1/captcha [get]
func GenerateCaptcha(c *gin.Context) {
	id, b64s, err := model.DriverDigitFunc()
	if err != nil {
		controller.SendResponse(nil, err, "获取验证码失败")
		return
	}
	controller.SendResponse(c, nil, map[string]interface{}{"data": b64s, "id": id})
}
