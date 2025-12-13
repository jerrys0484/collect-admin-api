// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Dispatch is the golang structure of table ct_dispatch for DAO operations like Where/Data.
type Dispatch struct {
	g.Meta     `orm:"table:ct_dispatch, do:true"`
	Id         any // ID
	Uuid       any // 标识
	Name       any // 名称
	Type       any // 类型
	Way        any // 方式
	Rules      any // 规则(json)
	CreateTime any // 创建时间
	UpdateTime any // 更新时间
}
