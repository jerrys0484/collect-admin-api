// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Steps is the golang structure of table ct_steps for DAO operations like Where/Data.
type Steps struct {
	g.Meta     `orm:"table:ct_steps, do:true"`
	Id         any // ID
	Uuid       any // 标识
	Name       any // 名称
	Type       any // 类型
	Data       any // 数据模板(json)
	Vars       any // 变量模板(json)
	Request    any // 请求参数(json)
	Response   any // 响应参数配置(json)
	CreateTime any // 创建时间
	UpdateTime any // 更新时间
}
