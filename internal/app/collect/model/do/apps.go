// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Apps is the golang structure of table ct_apps for DAO operations like Where/Data.
type Apps struct {
	g.Meta     `orm:"table:ct_apps, do:true"`
	Id         any // ID
	Uuid       any // 标识
	Name       any // 名称
	Template   any // 模板标识
	Rules      any // 调度配置(json)
	CreateTime any // 创建时间
	UpdateTime any // 更新时间
}
