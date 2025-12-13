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

var Dispatch = dispatchController{}

type dispatchController struct{}

func (c *dispatchController) List(ctx context.Context, req *collet.DispatchSearchReq) (res *collet.DispatchSearchRes, err error) {
	res, err = service.Dispatch().List(ctx, req)
	return
}

func (c *dispatchController) Add(ctx context.Context, req *collet.DispatchAddReq) (res *collet.DispatchAddRes, err error) {
	err = service.Dispatch().Add(ctx, req)
	return
}

func (c *dispatchController) Get(ctx context.Context, req *collet.DispatchGetReq) (res *collet.DispatchGetRes, err error) {
	res, err = service.Dispatch().Get(ctx, req.Uuid)
	return
}

func (c *dispatchController) Edit(ctx context.Context, req *collet.DispatchEditReq) (res *collet.DispatchEditRes, err error) {
	err = service.Dispatch().Edit(ctx, req)
	return
}

func (c *dispatchController) Delete(ctx context.Context, req *collet.DispatchDeleteReq) (res *collet.DispatchDeleteRes, err error) {
	err = service.Dispatch().Delete(ctx, req.Uuids)
	return
}
