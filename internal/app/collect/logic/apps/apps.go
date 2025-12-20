package dispatch

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
	"github.com/tiger1103/gfast/v3/api/v1/collet"
	"github.com/tiger1103/gfast/v3/internal/app/collect/dao"
	"github.com/tiger1103/gfast/v3/internal/app/collect/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/collect/service"
	systemConsts "github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterApps(New())
}

func New() *sApps {
	return &sApps{}
}

type sApps struct{}

func (s *sApps) List(ctx context.Context, req *collet.AppsSearchReq) (res *collet.AppsSearchRes, err error) {
	res = new(collet.AppsSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.Apps.Ctx(ctx)
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

func (s *sApps) Add(ctx context.Context, req *collet.AppsAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		liberr.ErrIsNil(ctx, err)
		uid, _ := uuid.NewUUID()
		_, err = dao.Apps.Ctx(ctx).Insert(do.Apps{
			Uuid:       uid.String(),
			Name:       req.Name,
			Template:   req.Template,
			Rules:      req.Rules,
			CreateTime: time.Now().UTC().Unix(),
		})
		liberr.ErrIsNil(ctx, err, "Add Failed")
	})
	return
}

func (s *sApps) Get(ctx context.Context, uuid string) (res *collet.AppsGetRes, err error) {
	res = new(collet.AppsGetRes)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Apps.Ctx(ctx).Where("uuid = ?", uuid).Scan(&res.Data)
		liberr.ErrIsNil(ctx, err, "Get Failed")
	})
	return
}

func (s *sApps) Edit(ctx context.Context, req *collet.AppsEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		liberr.ErrIsNil(ctx, err)
		_, err = dao.Apps.Ctx(ctx).Where("uuid = ?", req.Uuid).Update(do.Apps{
			Name:       req.Name,
			Template:   req.Template,
			Rules:      req.Rules,
			UpdateTime: time.Now().UTC().Unix(),
		})
		liberr.ErrIsNil(ctx, err, "Edit Failed")
	})
	return
}

func (s *sApps) Delete(ctx context.Context, uuids []string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Apps.Ctx(ctx).Delete(dao.Apps.Columns().Uuid+" in (?)", uuids)
		liberr.ErrIsNil(ctx, err, "Delete Failed")
	})
	return
}
