package collet

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/collect/model/entity"
)

type TemplateSearchReq struct {
	g.Meta `path:"/template/list" tags:"Template" method:"get" summary:"List Template"`
	Name   string `p:"name"`
	Uuid   string `p:"uuid"`
	common.PageReq
}

type TemplateSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.Template `json:"list"`
	common.ListRes
}

type TemplateAdd struct {
	Name string `p:"name"  v:"required#Name can not empty"`
	Data string `p:"data"`
	Vars string `p:"vars"`
}

type TemplateAddReq struct {
	g.Meta `path:"/template/add" tags:"Template" method:"post" summary:"Add Template"`
	*TemplateAdd
}

type TemplateAddRes struct{}

type TemplateGetReq struct {
	g.Meta `path:"/template/get" tags:"Template" method:"get" summary:"Get Template"`
	Uuid   string `p:"uuid"`
}

type TemplateGetRes struct {
	g.Meta `mime:"application/json"`
	Data   *entity.Template `json:"data"`
}

type TemplateEdit struct {
	Uuid string `p:"uuid" v:"required#Uuid can not empty"`
	Name string `p:"name"`
	Vars string `p:"vars"`
	Data string `p:"data"`
}

type TemplateEditReq struct {
	g.Meta `path:"/template/edit" tags:"Template" method:"put" summary:"Edit Template"`
	*TemplateEdit
}

type TemplateEditRes struct{}

type TemplateDeleteReq struct {
	g.Meta `path:"/template/delete" tags:"Template" method:"delete" summary:"Delete Template"`
	Uuids  []string `p:"uuids"`
}

type TemplateDeleteRes struct{}
