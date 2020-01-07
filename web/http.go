package web

import (
	"errors"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/honlyc/struct2all/program"
	"github.com/urfave/cli"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
)

// project project config
var ()

func CreateHttp(c *cli.Context) (err error) {
	file := c.String("file")
	if file == "" {
		file, _ = os.Getwd()
	}
	//fi, err := os.Stat(file)
	//if err != nil {
	//	log.Error("get file info [%s] failed:%v", file, err)
	//	return err
	//}

	pattern := c.String("struct")
	if pattern == "" {
		return errors.New("struct is empty")
	}

	/*out := c.String("out")
	if out == "" {
		return errors.New("output file is empty")
	}*/

	fset := token.NewFileSet()
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Error("read [file:%s] failed:%v", file, err)
		return err
	}
	f, err := parser.ParseFile(fset, file, string(data), parser.ParseComments)
	if err != nil {
		log.Error("parse [file:%s] failed:%v", file, err)
		return err
	}

	matchFunc := func(structName string) bool {
		match, _ := filepath.Match(pattern, structName)
		return match
	}
	var types []*ast.TypeSpec

	types = program.FindMatchStruct([]*ast.File{f}, matchFunc)

	log.Info("get %d matched struct", len(types))

	sqls := []string{}
	for _, typ := range types {
		ms, err := NewWebGenerator(typ)
		if err != nil {
			log.Error("create model struct failed:%v", err)
			return err
		}

		sql, err := ms.CreateWebPage(pattern)
		if err != nil {
			log.Error("generate sql failed:%v", err)
			return err
		}
		sqls = append(sqls, sql)
	}

	log.Info("%s", sqls)

	return nil
}
