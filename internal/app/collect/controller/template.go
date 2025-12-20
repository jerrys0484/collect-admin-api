/*
* @desc:验证码获取
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/3/2 17:45
 */

package controller

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/collet"
	"github.com/tiger1103/gfast/v3/internal/app/collect/service"
)

var Template = templateController{}

type templateController struct{}

func (c *templateController) List(ctx context.Context, req *collet.TemplateSearchReq) (res *collet.TemplateSearchRes, err error) {
	res, err = service.Template().List(ctx, req)
	return
}

func (c *templateController) Add(ctx context.Context, req *collet.TemplateAddReq) (res *collet.TemplateAddRes, err error) {
	err = service.Template().Add(ctx, req)
	return
}

func (c *templateController) Get(ctx context.Context, req *collet.TemplateGetReq) (res *collet.TemplateGetRes, err error) {
	res, err = service.Template().Get(ctx, req.Uuid)
	return
}

func (c *templateController) Edit(ctx context.Context, req *collet.TemplateEditReq) (res *collet.TemplateEditRes, err error) {
	err = service.Template().Edit(ctx, req)
	return
}

func (c *templateController) Delete(ctx context.Context, req *collet.TemplateDeleteReq) (res *collet.TemplateDeleteRes, err error) {
	err = service.Template().Delete(ctx, req.Uuids)
	return
}
