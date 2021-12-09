package pkg

import (
	"bytes"
	"fmt"
	"text/template"
)

type GenTpl struct {
}

func (p *GenTpl) Gen(file string, tpls []ITemplate) {
	for _, tpl := range tpls {
		tpl.ParseData(file)
		t, err := template.New(tpl.GetFileName()).Parse(tpl.GetTemplate())
		if err != nil {
			fmt.Printf("Parse err: %#v", err.Error())
			return
		}
		var buf bytes.Buffer
		if err := t.Execute(&buf, tpl.GetData()); err != nil {
			fmt.Printf("Execute err: %#v", err.Error())
			return
		}
		Chk(MkdirPathIfNotExist(tpl.GetPath()))
		if err := SaveFile(tpl.GetPath(), tpl.GetFileName(), buf.Bytes()); err != nil {
			fmt.Printf("save file error: %#v", err)
			return
		}
	}
}
