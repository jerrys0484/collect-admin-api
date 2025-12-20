package collet

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/collect/model/entity"
)

type AppsSearchReq struct {
	g.Meta `path:"/apps/list" tags:"Apps" method:"get" summary:"List Apps"`
	Name   string `p:"name"`
	Uuid   string `p:"uuid"`
	common.PageReq
}

type AppsSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.Apps `json:"list"`
	common.ListRes
}

type AppsAdd struct {
	Name     string `p:"name"  v:"required#Name can not empty"`
	Template string `p:"template"`
	Rules    string `p:"rules"`
}

type AppsAddReq struct {
	g.Meta `path:"/apps/add" tags:"Apps" method:"post" summary:"Add Apps"`
	*AppsAdd
}

type AppsAddRes struct{}

type AppsGetReq struct {
	g.Meta `path:"/apps/get" tags:"Apps" method:"get" summary:"Get Apps"`
	Uuid   string `p:"uuid"`
}

type AppsGetRes struct {
	g.Meta `mime:"application/json"`
	Data   *entity.Apps `json:"data"`
}

type AppsEdit struct {
	Uuid     string `p:"uuid" v:"required#Uuid can not empty"`
	Name     string `p:"name"`
	Template string `p:"template"`
	Rules    string `p:"rules"`
}

type AppsEditReq struct {
	g.Meta `path:"/apps/edit" tags:"Apps" method:"put" summary:"Edit Apps"`
	*AppsEdit
}

type AppsEditRes struct{}

type AppsDeleteReq struct {
	g.Meta `path:"/apps/delete" tags:"Apps" method:"delete" summary:"Delete Apps"`
	Uuids  []string `p:"uuids"`
}

type AppsDeleteRes struct{}
