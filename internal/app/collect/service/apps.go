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
	IApps interface {
		List(ctx context.Context, req *collet.AppsSearchReq) (res *collet.AppsSearchRes, err error)
		Add(ctx context.Context, req *collet.AppsAddReq) (err error)
		Get(ctx context.Context, uuid string) (res *collet.AppsGetRes, err error)
		Edit(ctx context.Context, req *collet.AppsEditReq) (err error)
		Delete(ctx context.Context, uuids []string) (err error)
	}
)

var (
	localApps IApps
)

func Apps() IApps {
	if localApps == nil {
		panic("implement not found for interface IApps, forgot register?")
	}
	return localApps
}

func RegisterApps(i IApps) {
	localApps = i
}
