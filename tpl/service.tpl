package service

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/hedan806/h-fast/app/admin/fast/internal/model"
)

func (s *Service) {{.MName}}List(c context.Context, qbr *model.Query{{.MName}}Request) (res *model.Query{{.MName}}Res, err error) {
	res = &model.Query{{.MName}}Res{}
    var (
        db = s.dao.ORM.Table(model.{{.MName}}{}.TableName())
    )
    db.Count(&res.TotalSize)
    if err = db.Order("id DESC").Offset((qbr.PageNum - 1) * qbr.PageSize).Limit(qbr.PageSize).Find(&res.Result).Error; err != nil {
        log.Error("s.Get{{.MName}}s query error(%v)", err)
    }

    res.PageSize = qbr.PageSize /**/
    res.PageNum = qbr.PageNum
    return
}

func (s *Service) {{.MName}}Add(c context.Context, m *model.{{.MName}}) (res int64, err error) {
	if err = s.dao.ORM.Model(model.{{.MName}}{}).Create(m).Error; err != nil {
		log.Error("s.{{.MName}}Add error(%v)", err)
	}
	return
}

func (s *Service) {{.MName}}Update(c context.Context, m *model.{{.MName}}) (res int64, err error) {
	if err = s.dao.Updates(model.{{.MName}}{Id: m.Id}, m); err != nil {
		return
	}
	return
}

func (s *Service) {{.MName}}Info(c context.Context, id uint64) (res *model.{{.MName}}, err error) {
	res = &model.{{.MName}}{}
	where := model.{{.MName}}{}
	where.Id = id
	if _, err = s.dao.First(&where, &res); err != nil {
		return
	}
	return
}

func (s *Service) {{.MName}}Del(c context.Context, id int) (res int, err error) {
	if err = s.dao.ORM.Table(model.{{.MName}}{}.TableName()).Where("id = ?", id).Delete(&model.{{.MName}}{}).Error;err != nil{
	    return
	}
	return
}