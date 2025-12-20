// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Template is the golang structure for table template.
type Template struct {
	Id         int    `json:"id"         orm:"id"          description:"ID"`
	Uuid       string `json:"uuid"       orm:"uuid"        description:"标识"`
	Name       string `json:"name"       orm:"name"        description:"名称"`
	Vars       string `json:"vars"       orm:"vars"        description:"变量模板(json)"`
	Data       string `json:"data"       orm:"data"        description:"数据模板(json)"`
	CreateTime uint   `json:"createTime" orm:"create_time" description:"创建时间"`
	UpdateTime uint   `json:"updateTime" orm:"update_time" description:"更新时间"`
}
