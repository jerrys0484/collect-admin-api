package collet

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/collect/model/entity"
)

type DispatchSearchReq struct {
	g.Meta `path:"/dispatch/list" tags:"Dispatch" method:"get" summary:"List Dispatch"`
	Name   string `p:"name"`
	Uuid   string `p:"uuid"`
	common.PageReq
}

type DispatchSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.Dispatch `json:"list"`
	common.ListRes
}

type DispatchReq struct {
	Name  string `p:"name"  v:"required#Name can not empty"`
	Rules string `p:"rules"  v:"required#Rules can not empty"`
	Type  int    `p:"type"`
	Way   int    `p:"way"`
}

type DispatchAddReq struct {
	g.Meta `path:"/dispatch/add" tags:"Dispatch" method:"post" summary:"Add Dispatch"`
	*DispatchReq
}

type DispatchAddRes struct{}

type DispatchGetReq struct {
	g.Meta `path:"/dispatch/get" tags:"Dispatch" method:"get" summary:"Get Dispatch"`
	Uuid   string `p:"uuid"`
}

type DispatchGetRes struct {
	g.Meta `mime:"application/json"`
	Data   *entity.Dispatch `json:"data"`
}

type DispatchEditReq struct {
	g.Meta `path:"/dispatch/edit" tags:"Dispatch" method:"put" summary:"Edit Dispatch"`
	Uuid   string `p:"uuid" v:"required|Uuid can not empty"`
	*DispatchReq
}

type DispatchEditRes struct{}

type DispatchDeleteReq struct {
	g.Meta `path:"/dispatch/delete" tags:"Dispatch" method:"delete" summary:"Del Dispatch"`
	Uuids  []string `p:"uuids"`
}

type DispatchDeleteRes struct{}
