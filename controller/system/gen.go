package system

import (
	"api-gin-web/controller"
	"api-gin-web/model/base"
	"bytes"
	"github.com/gin-gonic/gin"
	"strconv"
	"text/template"
)

func Preview(c *gin.Context) {
	table := base.SysTables{}
	id, err := strconv.Atoi(c.Param("tableId"))
	if err != nil {
		controller.SendResponse(c, err, nil)
		return
	}

	table.TableId = id
	t1, err := template.ParseFiles("template/model.go.template")
	if err != nil {
		controller.SendResponse(c, err, nil)
		return
	}
	t2, err := template.ParseFiles("template/api.go.template")
	if err != nil {
		controller.SendResponse(c, err, nil)
		return
	}
	tab, _ := table.Get()
	var b1 bytes.Buffer
	err = t1.Execute(&b1, tab)
	var b2 bytes.Buffer
	err = t2.Execute(&b2, tab)

	mp := make(map[string]interface{})
	mp["template/model.go.template"] = b1.String()
	mp["template/api.go.template"] = b2.String()
	controller.SendResponse(c, nil, mp)
}
