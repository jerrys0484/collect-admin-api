package collet

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/collect/model/entity"
)

type StepsSearchReq struct {
	g.Meta `path:"/steps/list" tags:"Steps" method:"get" summary:"List Steps"`
	Name   string `p:"name"`
	Uuid   string `p:"uuid"`
	common.PageReq
}

type StepsSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.Steps `json:"list"`
	common.ListRes
}

type StepsAdd struct {
	Name     string `p:"name"  v:"required#Name can not empty"`
	Type     string `p:"type"  v:"required#Type can not empty"`
	Template string `p:"template"`
	Request  string `p:"request"`
	Response string `p:"response"`
}

type StepsAddReq struct {
	g.Meta `path:"/steps/add" tags:"Steps" method:"post" summary:"Add Steps"`
	*StepsAdd
}

type StepsAddRes struct{}

type StepsGetReq struct {
	g.Meta `path:"/steps/get" tags:"Steps" method:"get" summary:"Get Steps"`
	Uuid   string `p:"uuid"`
}

type StepsGetRes struct {
	g.Meta `mime:"application/json"`
	Data   *entity.Steps `json:"data"`
}

type StepsEdit struct {
	Uuid     string `p:"uuid" v:"required#Uuid can not empty"`
	Name     string `p:"name"`
	Type     string `p:"type"`
	Template string `p:"template"`
	Request  string `p:"request"`
	Response string `p:"response"`
}

type StepsEditReq struct {
	g.Meta `path:"/steps/edit" tags:"Steps" method:"put" summary:"Edit Steps"`
	*StepsEdit
}

type StepsEditRes struct{}

type StepsDeleteReq struct {
	g.Meta `path:"/steps/delete" tags:"Steps" method:"delete" summary:"Del Steps"`
	Uuids  []string `p:"uuids"`
}

type StepsDeleteRes struct{}

type StepsDebugReq struct {
	g.Meta       `path:"/steps/debug" tags:"Steps" method:"post" summary:"Debug Steps"`
	Uuid         string `p:"uuid" v:"required#Uuid can not empty"`
	Group        string `p:"group" v:"required#Group can not empty"`
	HttpResponse string `p:"httpResponse"`
	Params       string `p:"params" v:"required#Params can not empty"`
}

type StepsDebugRes struct {
	g.Meta        `mime:"application/json"`
	HttpResponse  string `json:"httpResponse"`
	DebugResponse string `json:"debugResponse"`
}

type StepDebugCollectRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    CollectData `json:"data"`
}

type CollectData struct {
	Code     int                    `json:"code"`
	Response map[string]interface{} `json:"response"`
	Collect  map[string]interface{} `json:"collect"`
}
