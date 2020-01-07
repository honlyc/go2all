package model

import (
	"github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"
	"github.com/hedan806/h-fast/app/admin/fast/internal/model"
	"github.com/hedan806/h-fast/library/ecode"
)

type Query{{.MName}}Request struct {
	Select string `json:"select"`
	Value  string `json:"value"`
	{{.MName}}
	Pagination
}

type Query{{.MName}}Res struct {
	Result []*{{.MName}} `json:"items"`
	Pagination
}
