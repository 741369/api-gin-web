package system

import (
	"api-gin-web/controller"
	"api-gin-web/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
)

// @Summary 字典数据列表
// @Description 获取JSON
// @Tags 字典数据
// @Param status query string false "status"
// @Param dictCode query string false "dictCode"
// @Param dictType query string false "dictType"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} controller.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/data/list [get]
// @Security
func GetDictDataList(c *gin.Context) {
	var data model.DictData
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize, _ = strconv.Atoi(size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex, _ = strconv.Atoi(index)
	}

	data.DictLabel = c.Request.FormValue("dictLabel")
	data.Status = c.Request.FormValue("status")
	data.DictType = c.Request.FormValue("dictType")
	id := c.Request.FormValue("dictCode")
	data.DictCode, _ = strconv.Atoi(id)
	//data.DataScope = utils.GetUserIdStr(c)
	data.DataScope = "0"
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

// @Summary 通过编码获取字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Param dictCode path int true "字典编码"
// @Success 200 {object} controller.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/data/{dictCode} [get]
// @Security
func GetDictData(c *gin.Context) {
	var DictData model.DictData
	DictData.DictLabel = c.Request.FormValue("dictLabel")
	DictData.DictCode, _ = strconv.Atoi(c.Param("dictCode"))
	result, err := DictData.GetByCode()
	if err != nil {
		controller.SendResponse(c, err, nil)
		return
	}
	controller.SendResponse(c, nil, result)
}

// @Summary 通过字典类型获取字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Param dictType path int true "dictType"
// @Success 200 {object} controller.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/databyType/{dictType} [get]
// @Security
func GetDictDataByDictType(c *gin.Context) {
	var DictData model.DictData
	DictData.DictType = c.Param("dictType")
	result, err := DictData.Get()
	if err != nil {
		controller.SendResponse(c, err, nil)
		return
	}
	controller.SendResponse(c, nil, result)
}

// @Summary 添加字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Accept  application/json
// @Product application/json
// @Param data body model.DictType true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dict/data [post]
// @Security Bearer
func InsertDictData(c *gin.Context) {
	var data model.DictData
	err := c.BindWith(&data, binding.JSON)
	data.CreateBy = "0"
	result, err := data.Create()
	if err != nil {
		controller.SendResponse(c, err, nil)
		return
	}
	controller.SendResponse(c, nil, result)
}

// @Summary 修改字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Accept  application/json
// @Product application/json
// @Param data body model.DictType true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dict/data [put]
// @Security Bearer
func UpdateDictData(c *gin.Context) {
	var data model.DictData
	err := c.BindWith(&data, binding.JSON)
	data.UpdateBy = "0"
	if err != nil {
		controller.SendResponse(c, err, nil)
		return
	}
	result, err := data.Update(data.DictCode)
	if err != nil {
		controller.SendResponse(c, err, nil)
		return
	}
	controller.SendResponse(c, nil, result)
}

// @Summary 删除字典数据
// @Description 删除数据
// @Tags 字典数据
// @Param dictCode path int true "dictCode"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/dict/data/{dictCode} [delete]
func DeleteDictData(c *gin.Context) {
	var data model.DictData
	id, err := strconv.Atoi(c.Param("dictCode"))
	data.UpdateBy = "0"
	_, err = data.Delete(id)
	if err != nil {
		controller.SendResponse(c, err, nil)
		return
	}
	controller.SendResponse(c, nil, nil)
}
