package system

import (
	"api-gin-web/controller"
	"api-gin-web/model/base"
	"api-gin-web/utils"
	"github.com/741369/go_utils/log"
	"github.com/gin-gonic/gin"
)

// @Summary 分页列表数据 / page list data
// @Description 数据库表列分页列表 / database table column page list
// @Tags 工具 / system
// @Param table_name query string true "table_name / 数据表名称"
// @Param page query int false "page / 页码"
// @Param page_size query int false "page_size / 页条数"
// @Success 200 {object} controller.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/db/columns/page [get]
func GetDBColumnList(c *gin.Context) {
	var data base.DBColumns

	reqParam := utils.PostParam2(c)
	log.Infof(c, "req_param = %#v", reqParam)
	offset, limit := utils.GetPagePageSize(reqParam)
	data.TableName = reqParam["table_name"]
	result, count, err := data.GetPage(offset, limit)
	if err != nil {
		controller.SendResponse(c, err, nil)
		return
	}

	var mp = make(map[string]interface{}, 2)
	mp["list"] = result
	mp["count"] = count
	controller.SendResponse(c, nil, mp)
}
