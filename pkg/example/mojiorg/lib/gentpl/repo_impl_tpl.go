package gentpl

import (
	"github.com/mojiorg/gentpl/pkg/example/mojiorg/lib/gentpl/util"
)

type RepoImplTpl struct {
	BaseDir  string
	Data     interface{}
	FileName string
}

func (p *RepoImplTpl) GetTemplate() string {
	return `package impl

import (
    "context"
    "github.com/ahmetb/go-linq/v3"
    "github.com/pkg/errors"
)

type {{ .content.Name }}RepoImpl struct {
	Ds *database.Ds
}

func (p *{{ .content.Name }}RepoImpl) Update(ctx context.Context, in *domain.{{ .content.Name }}) error {
	if err := in.Valid(); err != nil {
		return err
	}

	err := p.Ds.Gdb().Save(model.{{ .content.Name }}{}.New(in)).Error

	return errors.WithStack(err)
}

func (p *{{ .content.Name }}RepoImpl) Create(ctx context.Context, in *domain.{{ .content.Name }}) error {
	if err := in.Valid(); err != nil {
		return err
	}
	obj := model.{{ .content.Name }}{}.New(in)
	if err := p.Ds.Gdb().Create(obj).Error; err != nil {
		return errors.WithStack(err)
	}
	in.Id = obj.Id
	return nil
}

func (p *{{ .content.Name }}RepoImpl) MustGet(ctx context.Context, id int) (*domain.{{ .content.Name }}, error) {
	var o model.{{ .content.Name }}
	if err := p.Ds.Gdb().Take(&o, id).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return o.ToDomain(), nil
}


func (p *{{ .content.Name }}RepoImpl) MultiGet(ctx context.Context, id ...int) (domain.{{ .content.Name }}List, error) {
	if len(id) == 0 {
		return domain.{{ .content.Name }}List{}, nil
	}

	var l []*model.{{ .content.Name }}
	if err := p.Ds.Gdb().
		Where("id in (?)", id).
		Find(&l).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	var out domain.{{ .content.Name }}List
	linq.From(l).SelectT(model.{{ .content.Name }}{}.ModelToDomain).ToSlice(&out)
	return out, nil
}`
}

func (p *RepoImplTpl) GetPath() string {
	return p.GetBaseDir() + `../repo/impl`
}

func (p *RepoImplTpl) GetFileName() string {
	return p.FileName
}

func (p *RepoImplTpl) GetData() interface{} {
	return p.Data
}

func (p *RepoImplTpl) ParseData(file string) {
	f := util.FilePath(file)
	p.Data = util.ParseData(file)
	p.FileName = f.FileName()
	if p.BaseDir == "" {
		p.BaseDir = f.FolderName()
	}
	return
}

func (p *RepoImplTpl) SetBaseDir(dir string) {
	p.BaseDir = dir
}

func (p *RepoImplTpl) GetBaseDir() string {
	return p.BaseDir
}
