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

var Apps = appsController{}

type appsController struct{}

func (c *appsController) List(ctx context.Context, req *collet.AppsSearchReq) (res *collet.AppsSearchRes, err error) {
	res, err = service.Apps().List(ctx, req)
	return
}

func (c *appsController) Add(ctx context.Context, req *collet.AppsAddReq) (res *collet.AppsAddRes, err error) {
	err = service.Apps().Add(ctx, req)
	return
}

func (c *appsController) Get(ctx context.Context, req *collet.AppsGetReq) (res *collet.AppsGetRes, err error) {
	res, err = service.Apps().Get(ctx, req.Uuid)
	return
}

func (c *appsController) Edit(ctx context.Context, req *collet.AppsEditReq) (res *collet.AppsEditRes, err error) {
	err = service.Apps().Edit(ctx, req)
	return
}

func (c *appsController) Delete(ctx context.Context, req *collet.AppsDeleteReq) (res *collet.AppsDeleteRes, err error) {
	err = service.Apps().Delete(ctx, req.Uuids)
	return
}
