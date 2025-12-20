// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Template is the golang structure of table ct_template for DAO operations like Where/Data.
type Template struct {
	g.Meta     `orm:"table:ct_template, do:true"`
	Id         any // ID
	Uuid       any // 标识
	Name       any // 名称
	Vars       any // 变量模板(json)
	Data       any // 数据模板(json)
	CreateTime any // 创建时间
	UpdateTime any // 更新时间
}
