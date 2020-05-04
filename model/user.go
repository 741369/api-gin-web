/**********************************************
** @Des:
** @Author: 1@lg1024.com
** @Last Modified time: 2020/5/3 上午12:00
***********************************************/

package model

import (
	"api-gin-web/model/base"
	"errors"
	"fmt"
	"github.com/741369/go_utils/log"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

type Login struct {
	Username string `form:"UserName" json:"username" binding:"required"`
	Password string `form:"Password" json:"password" binding:"required"`
	Code     string `form:"Code" json:"code" binding:"required"`
	UUID     string `form:"UUID" json:"uuid" binding:"required"`
}

// User
type User struct {
	// key
	IdentityKey string
	// 用户名
	UserName  string
	FirstName string
	LastName  string
	// 角色
	Role string
}

type UserName struct {
	Username string `gorm:"type:varchar(64)" json:"username"`
}

type PassWord struct {
	// 密码
	Password string `gorm:"type:varchar(128)" json:"password"`
}

type LoginM struct {
	UserName
	PassWord
}

type SysUserId struct {
	UserId int `gorm:"primary_key;AUTO_INCREMENT"  json:"userId"` // 编码
}

type SysUserB struct {
	NickName  string `gorm:"type:varchar(128)" json:"nickName"` // 昵称
	Phone     string `gorm:"type:varchar(11)" json:"phone"`     // 手机号
	RoleId    int    `gorm:"type:int(11)" json:"roleId"`        // 角色编码
	Salt      string `gorm:"type:varchar(255)" json:"salt"`     //盐
	Avatar    string `gorm:"type:varchar(255)" json:"avatar"`   //头像
	Sex       string `gorm:"type:varchar(255)" json:"sex"`      //性别
	Email     string `gorm:"type:varchar(128)" json:"email"`    //邮箱
	DeptId    int    `gorm:"type:int(11)" json:"deptId"`        //部门编码
	PostId    int    `gorm:"type:int(11)" json:"postId"`        //职位编码
	CreateBy  string `gorm:"type:varchar(128)" json:"createBy"` //
	UpdateBy  string `gorm:"type:varchar(128)" json:"updateBy"` //
	Remark    string `gorm:"type:varchar(255)" json:"remark"`   //备注
	Status    string `gorm:"type:int(1);" json:"status"`
	DataScope string `gorm:"-" json:"dataScope"`
	Params    string `gorm:"-" json:"params"`
	base.BaseModel
}

type SysUser struct {
	SysUserId
	SysUserB
	LoginM
}

func (SysUser) TableName() string {
	return "sys_user"
}

type SysUserPwd struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type SysUserPage struct {
	SysUserId
	SysUserB
	LoginM
	DeptName string `gorm:"-" json:"deptName"`
}

type SysUserView struct {
	SysUserId
	SysUserB
	LoginM
	RoleName string `gorm:"column:role_name"  json:"role_name"`
}

func (u *Login) GetUser() (user SysUser, role SysRole, err error) {
	err = base.DB.TestDB.Where("username = ? ", u.Username).Find(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			err = errors.New("账号或密码错误(代码204)")
		}
		return
	}
	log.Infof(nil, "input_pwd = %s, user_pwd = %s", u.Password, user.Password)
	_, err = CompareHashAndPassword(user.Password, u.Password)
	if err != nil {
		if strings.Contains(err.Error(), "hashedPassword is not the hash of the given password") {
			err = errors.New("账号或密码错误(代码201)")
		}
		return
	}
	err = base.DB.TestDB.Where("role_id = ? ", user.RoleId).First(&role).Error
	return
}

// 获取用户数据
func (e *SysUser) Get() (SysUserView SysUserView, err error) {
	db := base.DB.TestDB.Select([]string{"sys_user.*", "sys_role.role_name"})
	db = db.Joins("left join sys_role on sys_user.role_id=sys_role.role_id")
	if e.UserId != 0 {
		db = db.Where("user_id = ?", e.UserId)
	}

	if e.Username != "" {
		db = db.Where("username = ?", e.Username)
	}

	if e.Password != "" {
		db = db.Where("password = ?", e.Password)
	}

	if e.RoleId != 0 {
		db = db.Where("role_id = ?", e.RoleId)
	}

	if e.DeptId != 0 {
		db = db.Where("dept_id = ?", e.DeptId)
	}

	if e.PostId != 0 {
		db = db.Where("post_id = ?", e.PostId)
	}

	if err = db.First(&SysUserView).Error; err != nil {
		return
	}
	return
}

func (e *SysUser) GetPage(offset int, limit int) ([]SysUserPage, int, error) {
	var doc []SysUserPage

	db := base.DB.TestDB.Select("sys_user.*,sys_dept.dept_name")
	db = db.Joins("left join sys_dept on sys_dept.dept_id = sys_user.dept_id")

	if e.Username != "" {
		db = db.Where("username = ?", e.Username)
	}

	if e.DeptId != 0 {
		db = db.Where("sys_user.dept_id in (select dept_id from sys_dept where dept_path like ? )", "%"+strconv.Itoa(e.DeptId)+"%")
	}

	// 数据权限控制
	dataPermission := new(DataPermission)
	dataPermission.UserId, _ = strconv.Atoi(e.DataScope)
	db = dataPermission.GetDataScope("sys_user", db)

	var count int

	if err := db.Offset(offset).Limit(limit).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	db.Count(&count)
	return doc, count, nil
}

//添加
func (e SysUser) InsertUser() (id int, err error) {
	if err = e.Encrypt(); err != nil {
		return
	}

	db := base.DB.TestDB
	// check 用户名
	var count int
	db.Where("username = ?", e.Username).Count(&count)
	if count > 0 {
		err = errors.New("账户已存在！")
		return
	}

	//添加数据
	if err = db.Create(&e).Error; err != nil {
		return
	}
	id = e.UserId
	return
}

//加密
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
