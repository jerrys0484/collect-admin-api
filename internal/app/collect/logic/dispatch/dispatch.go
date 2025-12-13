package dispatch

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
	"github.com/tiger1103/gfast/v3/api/v1/collet"
	"github.com/tiger1103/gfast/v3/internal/app/collect/dao"
	"github.com/tiger1103/gfast/v3/internal/app/collect/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/collect/service"
	systemConsts "github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/library/liberr"
	"time"
)

func init() {
	service.RegisterDispatch(New())
}

func New() *sDispatch {
	return &sDispatch{}
}

type sDispatch struct{}

func (s *sDispatch) List(ctx context.Context, req *collet.DispatchSearchReq) (res *collet.DispatchSearchRes, err error) {
	res = new(collet.DispatchSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.Dispatch.Ctx(ctx)
		if req != nil {
			if req.Name != "" {
				m = m.Where("name like ?", "%"+req.Name+"%")
			}
		}
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		res.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = systemConsts.PageSize
		}
		err = m.Page(req.PageNum, req.PageSize).Order("id desc").Scan(&res.List)
		liberr.ErrIsNil(ctx, err, "List Failed")
	})
	return
}

func (s *sDispatch) Add(ctx context.Context, req *collet.DispatchAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		liberr.ErrIsNil(ctx, err)
		uid, _ := uuid.NewUUID()
		_, err = dao.Dispatch.Ctx(ctx).Insert(do.Dispatch{
			Uuid:       uid.String(),
			Name:       req.Name,
			Way:        0,
			Type:       0,
			Rules:      req.Rules,
			CreateTime: time.Now().UTC().Unix(),
		})
		liberr.ErrIsNil(ctx, err, "Add Failed")
	})
	return
}

func (s *sDispatch) Get(ctx context.Context, uuid string) (res *collet.DispatchGetRes, err error) {
	res = new(collet.DispatchGetRes)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Dispatch.Ctx(ctx).Where("uuid = ?", uuid).Scan(&res.Data)
		liberr.ErrIsNil(ctx, err, "Get Failed")
	})
	return
}

func (s *sDispatch) Edit(ctx context.Context, req *collet.DispatchEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		liberr.ErrIsNil(ctx, err)
		_, err = dao.Dispatch.Ctx(ctx).Where("uuid = ?", req.Uuid).Update(do.Dispatch{
			Name:       req.Name,
			Way:        req.Way,
			Type:       req.Type,
			Rules:      req.Rules,
			UpdateTime: time.Now().UTC().Unix(),
		})
		liberr.ErrIsNil(ctx, err, "Edit Failed")
	})
	return
}

func (s *sDispatch) Delete(ctx context.Context, uuids []string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Dispatch.Ctx(ctx).Delete(dao.Dispatch.Columns().Uuid+" in (?)", uuids)
		liberr.ErrIsNil(ctx, err, "Delete Failed")
	})
	return
}
