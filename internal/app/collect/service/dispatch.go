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
	IDispatch interface {
		List(ctx context.Context, req *collet.DispatchSearchReq) (res *collet.DispatchSearchRes, err error)
		Add(ctx context.Context, req *collet.DispatchAddReq) (err error)
		Get(ctx context.Context, uuid string) (res *collet.DispatchGetRes, err error)
		Edit(ctx context.Context, req *collet.DispatchEditReq) (err error)
		Delete(ctx context.Context, uuids []string) (err error)
	}
)

var (
	localDispatch IDispatch
)

func Dispatch() IDispatch {
	if localDispatch == nil {
		panic("implement not found for interface IDispatch, forgot register?")
	}
	return localDispatch
}

func RegisterDispatch(i IDispatch) {
	localDispatch = i
}
