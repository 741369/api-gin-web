package system

import (
	"api-gin-web/controller"
	"api-gin-web/model/base"
	"api-gin-web/utils"
	"github.com/gin-gonic/gin"
)

// @Summary 分页列表数据 / page list data
// @Description 数据库表分页列表 / database table page list
// @Tags 工具 / system
// @Param table_name query string false "table_name / 数据表名称"
// @Param page query int false "page / 页码"
// @Param page_size query int false "page_size / 页条数"
// @Success 200 {object} controller.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/db/tables/page [get]
func GetDBTableList(c *gin.Context) {
	var data base.DBTables
	reqParam := utils.PostParam2(c)
	offset, limit := utils.GetPagePageSize(reqParam)
	data.TableName = reqParam["table_name"]
	result, count, err := data.GetPage(offset, limit)
	if err != nil {
		controller.SendResponse(c, err, nil)
		return
	}

	var mp = make(map[string]interface{}, 3)
	mp["list"] = result
	mp["count"] = count
	controller.SendResponse(c, nil, mp)
}
