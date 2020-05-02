/**********************************************
** @Des:
** @Author: 1@lg1024.com
** @Last Modified time: 2020/5/3 上午12:00
***********************************************/

package model

type Login struct {
	Username string `form:"UserName" json:"username" binding:"required"`
	Password string `form:"Password" json:"password" binding:"required"`
	Code     string `form:"Code" json:"code" binding:"required"`
	UUID     string `form:"UUID" json:"uuid" binding:"required"`
}

type SysRole struct {
	RoleId    int    `json:"roleId" gorm:"primary_key;AUTO_INCREMENT"` // 角色编码
	RoleName  string `json:"roleName" gorm:"type:varchar(128);"`       // 角色名称
	Status    string `json:"status" gorm:"type:int(1);"`               //
	RoleKey   string `json:"roleKey" gorm:"type:varchar(128);"`        //角色代码
	RoleSort  int    `json:"roleSort" gorm:"type:int(4);"`             //角色排序
	Flag      string `json:"flag" gorm:"type:varchar(128);"`           //
	CreateBy  string `json:"createBy" gorm:"type:varchar(128);"`       //
	UpdateBy  string `json:"updateBy" gorm:"type:varchar(128);"`       //
	Remark    string `json:"remark" gorm:"type:varchar(255);"`         //备注
	Admin     bool   `json:"admin" gorm:"type:char(1);"`
	DataScope string `json:"dataScope" gorm:"-"`
	Params    string `json:"params" gorm:"-"`
	MenuIds   []int  `json:"menuIds" gorm:"-"`
	DeptIds   []int  `json:"deptIds" gorm:"-"`
	BaseModel
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

	BaseModel
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

func (u *Login) GetUser() (user SysUser, role SysRole, e error) {
	return
}
