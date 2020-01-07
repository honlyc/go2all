package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/honlyc/struct2all/file"
	"github.com/honlyc/struct2all/model"
	"github.com/honlyc/struct2all/util"
	"go/ast"
	"os"

	"path"
	"strings"
)

type WebGenerator struct {
	structName string
	modelType  *ast.StructType
}

func NewWebGenerator(typeSpec *ast.TypeSpec) (*WebGenerator, error) {
	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return nil, errors.New("typeSpec is not struct type")
	}

	return &WebGenerator{
		structName: typeSpec.Name.Name,
		modelType:  structType,
	}, nil
}

func (ms *WebGenerator) CreateWebPage(name string) (res string, err error) {
	var columns []*model.Column
	for _, field := range ms.getStructFieds(ms.modelType) {
		label := strings.Trim(field.Comment.Text(), "\n")
		key := getColumnName(field)
		if label == ""{
			label = key
		}
		column := model.NewColumn(label, key)
		if isPrimaryKey(field) {
			column.CanModify = false
			column.IgnoreEdit = true
			column.IsID = true
		}
		if isIgnoreEdit(field) {
			column.IgnoreEdit = true
		}
		if isIgnoreShow(field) {
			column.IgnoreShow = true
		}
		if isIgnoreFilter(field) {
			column.IgFilter = true
		}
		columns = append(columns, &column)
	}
	for _, col := range columns {
		bytes, _ := json.Marshal(col)
		fmt.Println(string(bytes))
	}
	p := model.Page{}
	p.Columns = columns
	p.Name = strings.ToLower(name)
	p.MName = name

	p.Path = path.Join("out", p.Name)
	if err = os.MkdirAll(p.Path, 0755); err != nil {
		return
	}
	if err = os.MkdirAll(path.Join(p.Path, "/service"), 0755); err != nil {
		return
	}
	if err = os.MkdirAll(path.Join(p.Path, "/model"), 0755); err != nil {
		return
	}
	if err = os.MkdirAll(path.Join(p.Path, "/http"), 0755); err != nil {
		return
	}

	fileName := path.Join(p.Path, p.Name)

	files := map[string]string{
		path.Join(p.Path, "/service", p.Name) + ".go": "service.tpl",
		path.Join(p.Path, "/http", p.Name) + ".go":    "http.tpl",
		path.Join(p.Path, "/model", p.Name) + ".go":   "model.tpl",
		fileName + ".vue": "page.tpl",
	}
	for k, f := range files {
		log.Info("%s-%s", k, f)
		if err = file.WriteData(k, f, &p); err != nil {
			log.Error("write error(%v)", err)
			return
		}
	}
	return "", nil
}

func isPrimaryKey(field *ast.Field) bool {
	tagStr := util.GetFieldTag(field, "gorm").Name
	gormSettings := ParseTagSetting(tagStr)
	if _, ok := gormSettings["PRIMARY_KEY"]; ok {
		return true
	}

	if len(field.Names) > 0 && strings.ToUpper(field.Names[0].Name) == "ID" {
		return true
	}

	return false
}
func isIgnoreEdit(field *ast.Field) bool {
	tagStr := util.GetFieldTag(field, "s2a").Name
	gormSettings := ParseTagSetting(tagStr)
	if _, ok := gormSettings["IGNOREEDIT"]; ok {
		return true
	}

	return false
}

func isIgnoreShow(field *ast.Field) bool {
	tagStr := util.GetFieldTag(field, "s2a").Name
	gormSettings := ParseTagSetting(tagStr)
	if _, ok := gormSettings["IGNORESHOW"]; ok {
		return true
	}

	return false
}

func isIgnoreFilter(field *ast.Field) bool {
	tagStr := util.GetFieldTag(field, "s2a").Name
	gormSettings := ParseTagSetting(tagStr)
	if _, ok := gormSettings["IGFILTER"]; ok {
		return true
	}

	return false
}

func getColumnName(field *ast.Field) string {
	tagStr := util.GetFieldTag(field, "json").Name
	gormSettings := ParseTagSetting(tagStr)
	if columnName, ok := gormSettings["COLUMN"]; ok {
		return columnName
	}

	if len(field.Names) > 0 {
		//return fmt.Sprintf("%s", casee.ToSnakeCase(field.Names[0].Name))
		return fmt.Sprintf("%s", tagStr)
	}

	return ""
}

func ParseTagSetting(str string) map[string]string {
	tags := strings.Split(str, ";")
	setting := map[string]string{}
	for _, value := range tags {
		v := strings.Split(value, ":")
		k := strings.TrimSpace(strings.ToUpper(v[0]))
		if len(v) == 2 {
			setting[k] = v[1]
		} else {
			setting[k] = k
		}
	}
	return setting
}

func (ms *WebGenerator) getStructFieds(node ast.Node) []*ast.Field {
	var fields []*ast.Field
	nodeType, ok := node.(*ast.StructType)
	if !ok {
		return nil
	}
	for _, field := range nodeType.Fields.List {
		if util.GetFieldTag(field, "sql").Name == "-" {
			continue
		}

		switch t := field.Type.(type) {
		case *ast.Ident:
			if t.Obj != nil && t.Obj.Kind == ast.Typ {
				if typeSpec, ok := t.Obj.Decl.(*ast.TypeSpec); ok {
					fields = append(fields, ms.getStructFieds(typeSpec.Type)...)
				}
			} else {
				fields = append(fields, field)
			}
		case *ast.SelectorExpr:
			fields = append(fields, field)
		default:
			log.Info("filed %s not supported, ignore", util.GetFieldName(field))
		}
	}

	return fields
}
