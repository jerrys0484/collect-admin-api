// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Dispatch is the golang structure for table dispatch.
type Dispatch struct {
	Id         int    `json:"id"         orm:"id"          description:"ID"`
	Uuid       string `json:"uuid"       orm:"uuid"        description:"标识"`
	Name       string `json:"name"       orm:"name"        description:"名称"`
	Type       uint   `json:"type"       orm:"type"        description:"类型"`
	Way        uint   `json:"way"        orm:"way"         description:"方式"`
	Rules      string `json:"rules"      orm:"rules"       description:"规则(json)"`
	CreateTime uint   `json:"createTime" orm:"create_time" description:"创建时间"`
	UpdateTime uint   `json:"updateTime" orm:"update_time" description:"更新时间"`
}
