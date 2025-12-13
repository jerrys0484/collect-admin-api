package steps

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
	"github.com/tiger1103/gfast/v3/api/v1/collet"
	"github.com/tiger1103/gfast/v3/internal/app/collect/dao"
	"github.com/tiger1103/gfast/v3/internal/app/collect/model/do"
	systemConsts "github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/library/liberr"
	"time"
)

func init() {
	//service.RegisterSteps(New())
}

func New() *sSteps {
	return &sSteps{}
}

type sSteps struct{}

// List 系统参数列表
func (s *sSteps) List(ctx context.Context, req *collet.StepsSearchReq) (res *collet.StepsSearchRes, err error) {
	res = new(collet.StepsSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.Steps.Ctx(ctx)
		if req != nil {
			if req.Name != "" {
				m = m.Where("config_name like ?", "%"+req.Name+"%")
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

func (s *sSteps) Add(ctx context.Context, req *collet.StepsAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		liberr.ErrIsNil(ctx, err)
		uid, _ := uuid.NewUUID()
		_, err = dao.Steps.Ctx(ctx).Insert(do.Steps{
			Uuid:       uid.String(),
			Name:       req.Name,
			Type:       req.Type,
			Request:    req.Request,
			Response:   req.Response,
			CreateTime: time.Now().UTC().Unix(),
		})
		liberr.ErrIsNil(ctx, err, "Add Failed")
	})
	return
}

// Get 获取系统参数
func (s *sSteps) Get(ctx context.Context, uuid string) (res *collet.StepsGetRes, err error) {
	res = new(collet.StepsGetRes)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Steps.Ctx(ctx).Where("uuid = ?", uuid).Scan(&res.Data)
		liberr.ErrIsNil(ctx, err, "Get Failed")
	})
	return
}

// Edit 修改系统参数
func (s *sSteps) Edit(ctx context.Context, req *collet.StepsEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		liberr.ErrIsNil(ctx, err)
		_, err = dao.Steps.Ctx(ctx).WherePri(req.Uuid).Update(do.Steps{
			Name:       req.Name,
			Type:       req.Type,
			Request:    req.Request,
			Response:   req.Response,
			UpdateTime: time.Now().UTC().Unix(),
		})
		liberr.ErrIsNil(ctx, err, "Edit Failed")
	})
	return
}

// Delete 删除系统参数
func (s *sSteps) Delete(ctx context.Context, uuids []string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Steps.Ctx(ctx).Delete(dao.Steps.Columns().Uuid+" in (?)", uuids)
		liberr.ErrIsNil(ctx, err, "Delete Failed")
	})
	return
}
