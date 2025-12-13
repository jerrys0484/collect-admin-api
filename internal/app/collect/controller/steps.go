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

var Steps = stepsController{}

type stepsController struct{}

func (c *stepsController) List(ctx context.Context, req *collet.StepsSearchReq) (res *collet.StepsSearchRes, err error) {
	res, err = service.Steps().List(ctx, req)
	return
}

func (c *stepsController) Add(ctx context.Context, req *collet.StepsAddReq) (res *collet.StepsAddRes, err error) {
	err = service.Steps().Add(ctx, req)
	return
}

func (c *stepsController) Get(ctx context.Context, req *collet.StepsGetReq) (res *collet.StepsGetRes, err error) {
	res, err = service.Steps().Get(ctx, req.Uuid)
	return
}

func (c *stepsController) Edit(ctx context.Context, req *collet.StepsEditReq) (res *collet.StepsEditRes, err error) {
	err = service.Steps().Edit(ctx, req)
	return
}

func (c *stepsController) Delete(ctx context.Context, req *collet.StepsDeleteReq) (res *collet.StepsDeleteRes, err error) {
	err = service.Steps().Delete(ctx, req.Uuids)
	return
}
