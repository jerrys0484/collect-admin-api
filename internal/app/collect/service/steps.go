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
	ISteps interface {
		// List 系统参数列表
		List(ctx context.Context, req *collet.StepsSearchReq) (res *collet.StepsSearchRes, err error)
		Add(ctx context.Context, req *collet.StepsAddReq) (err error)
		// Get 获取系统参数
		Get(ctx context.Context, uuid string) (res *collet.StepsGetRes, err error)
		// Edit 修改系统参数
		Edit(ctx context.Context, req *collet.StepsEditReq) (err error)
		// Delete 删除系统参数
		Delete(ctx context.Context, uuids []string) (err error)
	}
)

var (
	localSteps ISteps
)

func Steps() ISteps {
	if localSteps == nil {
		panic("implement not found for interface ISteps, forgot register?")
	}
	return localSteps
}

func RegisterSteps(i ISteps) {
	localSteps = i
}
