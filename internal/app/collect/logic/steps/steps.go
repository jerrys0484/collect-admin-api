package steps

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/google/uuid"
	"github.com/tiger1103/gfast/v3/api/v1/collet"
	"github.com/tiger1103/gfast/v3/internal/app/collect/dao"
	"github.com/tiger1103/gfast/v3/internal/app/collect/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/collect/service"
	systemConsts "github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterSteps(New())
}

func New() *sSteps {
	return &sSteps{}
}

type sSteps struct{}

func (s *sSteps) List(ctx context.Context, req *collet.StepsSearchReq) (res *collet.StepsSearchRes, err error) {
	res = new(collet.StepsSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.Steps.Ctx(ctx)
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

func (s *sSteps) Get(ctx context.Context, uuid string) (res *collet.StepsGetRes, err error) {
	res = new(collet.StepsGetRes)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Steps.Ctx(ctx).Where("uuid = ?", uuid).Scan(&res.Data)
		liberr.ErrIsNil(ctx, err, "Get Failed")
	})
	return
}

func (s *sSteps) Edit(ctx context.Context, req *collet.StepsEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		liberr.ErrIsNil(ctx, err)
		_, err = dao.Steps.Ctx(ctx).Where("uuid = ?", req.Uuid).Update(do.Steps{
			Name:       req.Name,
			Type:       req.Type,
			Request:    req.Request,
			Response:   req.Response,
			UpdateTime: time.Now().UTC().Unix(),
		})
		liberr.ErrIsNil(ctx, err, "Edit Failed")
		_, err = g.Redis().Del(ctx, "CT_STEPS_"+req.Uuid)
	})
	return
}

func (s *sSteps) Delete(ctx context.Context, uuids []string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Steps.Ctx(ctx).Delete(dao.Steps.Columns().Uuid+" in (?)", uuids)
		liberr.ErrIsNil(ctx, err, "Delete Failed")
	})
	return
}

func (s *sSteps) Debug(ctx context.Context, req *collet.StepsDebugReq) (res *collet.StepsDebugRes, err error) {
	res = new(collet.StepsDebugRes)
	err = g.Try(ctx, func(ctx context.Context) {
		host := g.Cfg().MustGet(ctx, "collect.host").String()
		endpoint := g.Cfg().MustGet(ctx, "collect.endpoint").String()
		url := host + endpoint + "/" + req.Uuid
		var err1 error
		var response *gclient.Response
		var params g.Map
		err = json.Unmarshal([]byte(req.Params), &params)
		response, err1 = g.Client().ContentJson().Timeout(60*time.Second).Post(ctx, url, params)
		defer func(response *gclient.Response) {
			err = response.Close()
		}(response)
		str := response.ReadAllString()
		if str == "" {
			res.HttpResponse = "Empty Collect API Server: " + url
			return
		}
		var collectRes collet.StepDebugCollectRes
		err1 = json.Unmarshal([]byte(str), &collectRes)
		if err1 != nil {
			res.HttpResponse = str
			res.DebugResponse = err1.Error()
			return
		}
		b2, _ := json.Marshal(collectRes.Data.Collect)
		b3, _ := json.Marshal(collectRes.Data.Response)
		res.HttpResponse = string(b3)
		res.DebugResponse = string(b2)
	})
	return
}
