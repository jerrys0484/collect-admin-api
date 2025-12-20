// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Apps is the golang structure for table apps.
type Apps struct {
	Id         int    `json:"id"         orm:"id"          description:"ID"`
	Uuid       string `json:"uuid"       orm:"uuid"        description:"标识"`
	Name       string `json:"name"       orm:"name"        description:"名称"`
	Template   string `json:"template"   orm:"template"    description:"模板标识"`
	Rules      string `json:"rules"      orm:"rules"       description:"调度配置(json)"`
	CreateTime uint   `json:"createTime" orm:"create_time" description:"创建时间"`
	UpdateTime uint   `json:"updateTime" orm:"update_time" description:"更新时间"`
}
