/**********************************************
** @Des: 用户操作日志
** @Author: 1@lg1024.com
** @Last Modified time: 2020/5/3 下午11:53
***********************************************/

package model

import (
	"api-gin-web/model/base"
	"time"
)

//sys_operlog
type SysOperLog struct {
	OperId        int       `json:"operId" gorm:"primary_key;AUTO_INCREMENT"` //日志编码
	Title         string    `json:"title" gorm:"size(255);comment:'操作模块'"`    //操作模块
	BusinessType  string    `json:"businessType" gorm:"type:varchar(128);"`   //操作类型
	BusinessTypes string    `json:"businessTypes" gorm:"type:varchar(128);"`
	Method        string    `json:"method" gorm:"type:varchar(128);"`        //函数
	RequestMethod string    `json:"requestMethod" gorm:"type:varchar(128);"` //请求方式
	OperatorType  string    `json:"operatorType" gorm:"type:varchar(128);"`  //操作类型
	OperName      string    `json:"operName" gorm:"type:varchar(128);"`      //操作者
	DeptName      string    `json:"deptName" gorm:"type:varchar(128);"`      //部门名称
	OperUrl       string    `json:"operUrl" gorm:"type:varchar(255);"`       //访问地址
	OperIp        string    `json:"operIp" gorm:"type:varchar(128);"`        //客户端ip
	OperLocation  string    `json:"operLocation" gorm:"type:varchar(128);"`  //访问位置
	OperParam     string    `json:"operParam" gorm:"type:varchar(255);"`     //请求参数
	Status        string    `json:"status" gorm:"type:int(1);"`              //操作状态
	OperTime      time.Time `json:"operTime" gorm:"type:timestamp;"`         //操作时间
	JsonResult    string    `json:"jsonResult" gorm:"type:varchar(255);"`    //返回数据
	CreateBy      string    `json:"createBy" gorm:"type:varchar(128);"`      //创建人
	UpdateBy      string    `json:"updateBy" gorm:"type:varchar(128);"`      //更新者
	DataScope     string    `json:"dataScope" gorm:"-"`                      //数据
	Params        string    `json:"params" gorm:"-"`                         //参数
	Remark        string    `json:"remark" gorm:"type:varchar(255);"`        //备注
	LatencyTime   string    `json:"latencyime" gorm:"type:varchar(128);"`    //耗时
	UserAgent     string    `json:"userAgent" gorm:"type:varchar(255);"`     //ua
	base.BaseModel
}

func (SysOperLog) TableName() string {
	return "sys_operlog"
}

// 条件查询用户操作日志
func (e *SysOperLog) Get() (doc SysOperLog, err error) {
	db := base.DB.TestDB
	if e.OperIp != "" {
		db = db.Where("oper_ip = ?", e.OperIp)
	}
	if e.OperId != 0 {
		db = db.Where("oper_id = ?", e.OperId)
	}
	err = db.First(&doc).Error
	return
}

// 分页查询用户操作日志列表
func (e *SysOperLog) GetSysOperLogList(offset int, limit int) (doc []SysOperLog, count int, err error) {
	db := base.DB.TestDB
	if e.OperIp != "" {
		db = db.Where("oper_ip = ?", e.OperIp)
	}
	if e.Status != "" {
		db = db.Where("status = ?", e.Status)
	}
	if e.OperName != "" {
		db = db.Where("oper_name = ?", e.OperName)
	}
	if e.BusinessType != "" {
		db = db.Where("business_type = ?", e.BusinessType)
	}

	err = db.Order("oper_id desc").Offset(offset).Limit(limit).Find(&doc).Count(&count).Error
	return
}

// 保存用户操作日志
func (e *SysOperLog) Create() (doc SysOperLog, err error) {
	e.CreateBy = "0"
	e.UpdateBy = "0"
	if err = base.DB.TestDB.Create(&e).Error; err != nil {
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 修改用户操作日志
func (e *SysOperLog) Update(id int) (update SysOperLog, err error) {
	db := base.DB.TestDB
	if err = db.First(&update, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	err = db.Model(&update).Updates(&e).Error
	return
}

// 批量删除指定id的用户操作日志
func (e *SysOperLog) BatchDelete(id []int) (rows int64, err error) {
	dbRes := base.DB.TestDB.Where(" oper_id in (?)", id).Delete(&SysOperLog{})
	rows = dbRes.RowsAffected
	err = dbRes.Error
	return
}
