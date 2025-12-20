// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Steps is the golang structure for table steps.
type Steps struct {
	Id         int    `json:"id"         orm:"id"          description:"ID"`
	Uuid       string `json:"uuid"       orm:"uuid"        description:"标识"`
	Name       string `json:"name"       orm:"name"        description:"名称"`
	Type       string `json:"type"       orm:"type"        description:"类型"`
	Template   string `json:"template"   orm:"template"    description:"模板标识"`
	Request    string `json:"request"    orm:"request"     description:"请求参数(json)"`
	Response   string `json:"response"   orm:"response"    description:"响应参数配置(json)"`
	CreateTime uint   `json:"createTime" orm:"create_time" description:"创建时间"`
	UpdateTime uint   `json:"updateTime" orm:"update_time" description:"更新时间"`
}
