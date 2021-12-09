package gentpl

import (
	"github.com/mojiorg/gentpl/pkg/example/mojiorg/lib/gentpl/util"
)

type RepoTpl struct {
	BaseDir  string
	Data     interface{}
	FileName string
}

func (p *RepoTpl) GetTemplate() string {
	return `package repo

import (
	"context"
)
//go:generate mockgen -destination=../mocks/mrepo/{{ .fileName }} -package=mrepo . {{ .content.Name }}Repo

type {{ .content.Name }}Repo interface {
    MustGet(ctx context.Context, id int) (*domain.{{ .content.Name }}, error)
    MultiGet(ctx context.Context, id ...int) (domain.{{ .content.Name }}List, error)
    Create(ctx context.Context, in *domain.{{ .content.Name }}) error
    Update(ctx context.Context, in *domain.{{ .content.Name }}) error
}`
}

func (p *RepoTpl) GetPath() string {
	return p.GetBaseDir() + `../repo`
}

func (p *RepoTpl) GetFileName() string {
	return p.FileName
}

func (p *RepoTpl) GetData() interface{} {
	return p.Data
}

func (p *RepoTpl) ParseData(file string) {
	f := util.FilePath(file)
	p.Data = util.ParseData(file)
	p.FileName = f.FileName()
	if p.BaseDir == "" {
		p.BaseDir = f.FolderName()
	}
	return
}

func (p *RepoTpl) SetBaseDir(dir string) {
	p.BaseDir = dir
}

func (p *RepoTpl) GetBaseDir() string {
	return p.BaseDir
}
