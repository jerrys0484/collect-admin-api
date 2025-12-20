// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/collet"
)

type (
	ITemplate interface {
		List(ctx context.Context, req *collet.TemplateSearchReq) (res *collet.TemplateSearchRes, err error)
		Add(ctx context.Context, req *collet.TemplateAddReq) (err error)
		Get(ctx context.Context, uuid string) (res *collet.TemplateGetRes, err error)
		Edit(ctx context.Context, req *collet.TemplateEditReq) (err error)
		Delete(ctx context.Context, uuids []string) (err error)
	}
)

var (
	localTemplate ITemplate
)

func Template() ITemplate {
	if localTemplate == nil {
		panic("implement not found for interface ITemplate, forgot register?")
	}
	return localTemplate
}

func RegisterTemplate(i ITemplate) {
	localTemplate = i
}
