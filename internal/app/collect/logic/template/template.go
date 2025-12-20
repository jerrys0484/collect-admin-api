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
	service.RegisterTemplate(New())
}

func New() *sTemplate {
	return &sTemplate{}
}

type sTemplate struct{}

func (s *sTemplate) List(ctx context.Context, req *collet.TemplateSearchReq) (res *collet.TemplateSearchRes, err error) {
	res = new(collet.TemplateSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.Template.Ctx(ctx)
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

func (s *sTemplate) Add(ctx context.Context, req *collet.TemplateAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		liberr.ErrIsNil(ctx, err)
		uid, _ := uuid.NewUUID()
		_, err = dao.Template.Ctx(ctx).Insert(do.Template{
			Uuid:       uid.String(),
			Name:       req.Name,
			Vars:       req.Vars,
			Data:       req.Data,
			CreateTime: time.Now().UTC().Unix(),
		})
		liberr.ErrIsNil(ctx, err, "Add Failed")
	})
	return
}

func (s *sTemplate) Get(ctx context.Context, uuid string) (res *collet.TemplateGetRes, err error) {
	res = new(collet.TemplateGetRes)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Template.Ctx(ctx).Where("uuid = ?", uuid).Scan(&res.Data)
		liberr.ErrIsNil(ctx, err, "Get Failed")
	})
	return
}

func (s *sTemplate) Edit(ctx context.Context, req *collet.TemplateEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		liberr.ErrIsNil(ctx, err)
		_, err = dao.Template.Ctx(ctx).Where("uuid = ?", req.Uuid).Update(do.Template{
			Name:       req.Name,
			Vars:       req.Vars,
			Data:       req.Data,
			UpdateTime: time.Now().UTC().Unix(),
		})
		liberr.ErrIsNil(ctx, err, "Edit Failed")
	})
	return
}

func (s *sTemplate) Delete(ctx context.Context, uuids []string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Template.Ctx(ctx).Delete(dao.Template.Columns().Uuid+" in (?)", uuids)
		liberr.ErrIsNil(ctx, err, "Delete Failed")
	})
	return
}
