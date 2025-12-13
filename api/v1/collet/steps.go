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

type StepsReq struct {
	Name     string `p:"name"  v:"required#Name can not empty"`
	Type     string `p:"type"  v:"required#Type can not empty"`
	Request  string `p:"request" v:"required#Request can not empty"`
	Response string `p:"response"`
}

type StepsAddReq struct {
	g.Meta `path:"/steps/add" tags:"Steps" method:"post" summary:"Add Steps"`
	*StepsReq
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

type StepsEditReq struct {
	g.Meta `path:"/steps/edit" tags:"Steps" method:"put" summary:"Edit Steps"`
	Uuid   string `p:"uuid" v:"required#Uuid can not empty"`
	*StepsReq
}

type StepsEditRes struct{}

type StepsDeleteReq struct {
	g.Meta `path:"/steps/delete" tags:"Steps" method:"delete" summary:"Del Steps"`
	Uuids  []string `p:"uuids"`
}

type StepsDeleteRes struct{}
