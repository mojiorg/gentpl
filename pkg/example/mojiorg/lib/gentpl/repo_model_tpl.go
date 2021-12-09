package gentpl

import (
	"github.com/mojiorg/gentpl/pkg/example/mojiorg/lib/gentpl/util"
)

type RepoModelTpl struct {
	BaseDir  string
	Data     interface{}
	FileName string
}

func (p *RepoModelTpl) GetTemplate() string {
	return `package model

import (
	"baibao/settlement/pkg/domain"
	"github.com/jinzhu/copier"
	"github.com/quexer/utee"

)

type {{ .content.Name }} struct {
	{{ range .content.Fields -}}
	{{ .Name }}    {{ .Type }}    {{ .GormTag }}
	{{ end -}}
}

func ({{ .content.Name }}) ModelToDomain(x *{{ .content.Name }}) *domain.{{ .content.Name }} {
	return x.ToDomain()
}

func ({{ .content.Name }}) New(x *domain.{{ .content.Name }}) *{{ .content.Name }} {
	out := &{{ .content.Name }}{}
	utee.Chk(copier.Copy(out, x))
	return out
}

func (p *{{ .content.Name }}) ToDomain() *domain.{{ .content.Name }} {
	out := &domain.{{ .content.Name }}{}
	utee.Chk(copier.Copy(out, p))
	return out
}`
}

func (p *RepoModelTpl) GetPath() string {
	return p.GetBaseDir() + `../repo/internal/model`
}

func (p *RepoModelTpl) GetFileName() string {
	return p.FileName
}

func (p *RepoModelTpl) GetData() interface{} {
	return p.Data
}

func (p *RepoModelTpl) ParseData(file string) {
	f := util.FilePath(file)
	p.Data = util.ParseData(file)
	p.FileName = f.FileName()
	if p.BaseDir == "" {
		p.BaseDir = f.FolderName()
	}
	return
}

func (p *RepoModelTpl) SetBaseDir(dir string) {
	p.BaseDir = dir
}

func (p *RepoModelTpl) GetBaseDir() string {
	return p.BaseDir
}
