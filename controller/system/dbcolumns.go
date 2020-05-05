package system

import (
	"api-gin-web/controller"
	"api-gin-web/model/base"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary 分页列表数据 / page list data
// @Description 数据库表列分页列表 / database table column page list
// @Tags 工具 / system
// @Param tableName query string false "tableName / 数据表名称"
// @Param pageSize query int false "pageSize / 页条数"
// @Param pageIndex query int false "pageIndex / 页码"
// @Success 200 {object} controller.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/db/columns/page [get]
func GetDBColumnList(c *gin.Context) {
	var data base.DBColumns
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize, _ = strconv.Atoi(size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex, _ = strconv.Atoi(index)
	}

	data.TableName = c.Request.FormValue("tableName")
	//pkg.Assert(data.TableName == "", "table name cannot be empty！", 500)
	result, count, err := data.GetPage(pageSize, pageIndex)
	if err != nil {
		controller.SendResponse(c, err, nil)
		return
	}

	var mp = make(map[string]interface{}, 3)
	mp["list"] = result
	mp["count"] = count
	mp["pageIndex"] = pageIndex
	mp["pageSize"] = pageSize
	controller.SendResponse(c, nil, mp)
}
