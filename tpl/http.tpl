package http

import (
	"github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"
	"github.com/hedan806/h-fast/app/admin/fast/internal/model"
	"github.com/hedan806/h-fast/library/ecode"
)

// http

func init{{.MName}}(api *blademaster.RouterGroup){
    {{.Name}} := api.Group("/{{.Name}}", verify)
    {
        {{.Name}}.POST("/list", {{.Name}}List)
        {{.Name}}.GET("/detail", {{.Name}}Detail)
        {{.Name}}.POST("/create", {{.Name}}Create)
        {{.Name}}.POST("/update", {{.Name}}Update)
        {{.Name}}.DELETE("/", {{.Name}}Del)
    }
}


func {{.Name}}List(c *blademaster.Context) {
	qmr := model.Query{{.MName}}Request{}
	if err := c.BindWith(&qmr, binding.JSON); err != nil {
		c.JSON(nil, err)
		return
	}
	if err := qmr.Verify(); err != nil {
		c.JSON(nil, err)
		return
	}

	switch qmr.Select {
	}

	c.JSON(svc.{{.MName}}List(c, &qmr))
}

func {{.Name}}Create(c *blademaster.Context) {
	m := new(model.{{.MName}})
	if err := c.BindWith(&m, binding.JSON); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(svc.{{.MName}}Add(c, m))
}

func {{.Name}}Update(c *blademaster.Context) {
    m := new(model.{{.MName}})
	if err := c.BindWith(&m, binding.JSON); err != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(svc.{{.MName}}Update(c, m))
}

func {{.Name}}Detail(c *blademaster.Context) {
	c.JSON(svc.{{.MName}}Info(c, 0))
}

func {{.Name}}Del(c *blademaster.Context) {
	id := atoi(c.Request.Form.Get("id"))
	c.JSON(svc.{{.MName}}Del(c, id))
}

