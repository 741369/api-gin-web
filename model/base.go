/**********************************************
** @Des: 基准model
** @Author: 1@lg1024.com
** @Last Modified time: 2020/5/3 上午12:19
***********************************************/

package model

import "time"

type BaseModel struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}
