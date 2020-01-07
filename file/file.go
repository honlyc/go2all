package file

import (
	"bytes"
	"github.com/honlyc/struct2all/model"
	"io/ioutil"
	"strings"
	"text/template"
)

func init() {

	MyFunc = map[string]interface{}{}
	MyFunc["ToPropName"] = ToPropName
	MyFunc["To18IN"] = To18IN
}

func ToPropName(str string) string {
	return "{{ row." + strings.ToLower(str) + "}}"
}

// {{ $t('table.add') }}
func To18IN(str string) string {
	return "{{ $t('" + str + "') }}"
}

var MyFunc map[string]interface{}

func WriteData(name, tpl string, d *model.Page) (err error) {
	data, err := parseData(tpl, d)
	if err != nil {
		return
	}
	return ioutil.WriteFile(name, data, 0644)
}

func parseData(s string, data *model.Page) ([]byte, error) {
	t, err := template.New(s).Funcs(MyFunc).ParseFiles("./tpl/" + s)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
