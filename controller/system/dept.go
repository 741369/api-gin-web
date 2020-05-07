package system

import (
	"api-gin-web/controller"
	"api-gin-web/model"
	"api-gin-web/router/middleware/jwt"
	"api-gin-web/utils/errno"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
)

// @Summary 分页部门列表数据
// @Description 分页列表
// @Tags 部门 / dept
// @Param name query string false "name"
// @Param id query string false "id"
// @Param position query string false "position"
// @Success 200 {object} controller.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dept/list [get]
// @Security
func GetDeptList(c *gin.Context) {
	var Dept model.Dept
	Dept.DeptName = c.Request.FormValue("deptName")
	Dept.Status = c.Request.FormValue("status")
	Dept.DeptId, _ = strconv.Atoi(c.Request.FormValue("deptId"))
	Dept.DataScope = jwt.GetUserIdStr(c)
	result, err := Dept.SetDept(true)
	controller.SendResponse(c, err, result)
}

// @Summary 分页部门列表数据
// @Description 分页列表
// @Tags 部门 / dept
// @Param name query string false "name"
// @Param id query string false "id"
// @Param position query string false "position"
// @Success 200 {object} controller.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dept/tree [get]
// @Security
func GetDeptTree(c *gin.Context) {
	var Dept model.Dept
	Dept.DeptName = c.Request.FormValue("deptName")
	Dept.Status = c.Request.FormValue("status")
	Dept.DeptId, _ = strconv.Atoi(c.Request.FormValue("deptId"))
	result, err := Dept.SetDept(false)
	controller.SendResponse(c, err, result)
}

// @Summary 部门列表数据
// @Description 获取JSON
// @Tags 部门 / dept
// @Param deptId path string false "deptId"
// @Param position query string false "position"
// @Success 200 {object} controller.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dept/{deptId} [get]
// @Security
func GetDept(c *gin.Context) {
	var Dept model.Dept
	Dept.DeptId, _ = strconv.Atoi(c.Param("deptId"))
	Dept.DataScope = jwt.GetUserIdStr(c)
	result, err := Dept.Get()
	controller.SendResponse(c, err, result)
}

// @Summary 添加部门
// @Description 获取JSON
// @Tags 部门 / dept
// @Accept  application/json
// @Product application/json
// @Param data body model.Dept true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dept [post]
// @Security Bearer
func InsertDept(c *gin.Context) {
	var data model.Dept
	err := c.BindWith(&data, binding.JSON)
	if err != nil {
		controller.SendResponse(c, errno.ErrParam, nil)
		return
	}
	data.CreateBy = jwt.GetUserIdStr(c)
	result, err := data.Create()
	controller.SendResponse(c, err, result)
}

// @Summary 修改部门
// @Description 获取JSON
// @Tags 部门 / dept
// @Accept  application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body model.Dept true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dept [put]
// @Security Bearer
func UpdateDept(c *gin.Context) {
	var data model.Dept
	err := c.BindWith(&data, binding.JSON)
	if err != nil {
		controller.SendResponse(c, errno.ErrParam, nil)
		return
	}
	data.UpdateBy = jwt.GetUserIdStr(c)
	result, err := data.Update(data.DeptId)
	controller.SendResponse(c, err, result)
}

// @Summary 删除部门
// @Description 删除数据
// @Tags 部门 / dept
// @Param id path int true "id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/dept/{id} [delete]
func DeleteDept(c *gin.Context) {
	var data model.Dept
	id, err := strconv.Atoi(c.Param("id"))
	_, err = data.Delete(id)
	controller.SendResponse(c, err, nil)
}

/*func GetDeptTreeRoleselect(c *gin.Context) {
	var Dept model.Dept
	var SysRole model.SysRole
	id, err := strconv.Atoi(c.Param("roleId"))
	SysRole.RoleId = id
	result, err := Dept.SetDeptLable()
	pkg.HasError(err, msg.NotFound, -1)
	menuIds := make([]int, 0)
	if id != 0 {
		menuIds, err = SysRole.GetRoleDeptId()
		pkg.HasError(err, "抱歉未找到相关信息", -1)
	}
	app.Custum(c, gin.H{
		"code":        200,
		"depts":       result,
		"checkedKeys": menuIds,
	})
}*/
