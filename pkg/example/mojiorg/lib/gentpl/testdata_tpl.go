package gentpl

import (
	"github.com/mojiorg/gentpl/pkg/example/mojiorg/lib/gentpl/util"
)

type TestDataTpl struct {
	BaseDir  string
	Data     interface{}
	FileName string
}

func (p *TestDataTpl) GetTemplate() string {
	return `package testdata

func {{ .content.Name }} (id ...int) *domain.{{ .content.Name }} {
	obj := &domain.{{ .content.Name }} {
        {{ range .content.Fields -}}
            {{.Name}}  : {{.MockVal}},
		{{ end -}}
	}
	if len(id) > 0 {
        obj.Id = id[0]
    }
	return obj
}`
}

func (p *TestDataTpl) GetPath() string {
	return p.GetBaseDir() + `../testdata`
}

func (p *TestDataTpl) GetFileName() string {
	return p.FileName
}

func (p *TestDataTpl) GetData() interface{} {
	return p.Data
}

func (p *TestDataTpl) ParseData(file string) {
	f := util.FilePath(file)
	p.Data = util.ParseData(file)
	p.FileName = f.FileName()
	if p.BaseDir == "" {
		p.BaseDir = f.FolderName()
	}
	return
}

func (p *TestDataTpl) SetBaseDir(dir string) {
	p.BaseDir = dir
}

func (p *TestDataTpl) GetBaseDir() string {
	return p.BaseDir
}
