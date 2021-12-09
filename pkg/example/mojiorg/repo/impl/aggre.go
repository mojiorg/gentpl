package impl

import (
	"context"
	"github.com/ahmetb/go-linq/v3"
	"github.com/pkg/errors"
)

type AggreRepoImpl struct {
	Ds *database.Ds
}

func (p *AggreRepoImpl) Update(ctx context.Context, in *domain.Aggre) error {
	if err := in.Valid(); err != nil {
		return err
	}

	err := p.Ds.Gdb().Save(model.Aggre{}.New(in)).Error

	return errors.WithStack(err)
}

func (p *AggreRepoImpl) Create(ctx context.Context, in *domain.Aggre) error {
	if err := in.Valid(); err != nil {
		return err
	}
	obj := model.Aggre{}.New(in)
	if err := p.Ds.Gdb().Create(obj).Error; err != nil {
		return errors.WithStack(err)
	}
	in.Id = obj.Id
	return nil
}

func (p *AggreRepoImpl) MustGet(ctx context.Context, id int) (*domain.Aggre, error) {
	var o model.Aggre
	if err := p.Ds.Gdb().Take(&o, id).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return o.ToDomain(), nil
}

func (p *AggreRepoImpl) MultiGet(ctx context.Context, id ...int) (domain.AggreList, error) {
	if len(id) == 0 {
		return domain.AggreList{}, nil
	}

	var l []*model.Aggre
	if err := p.Ds.Gdb().
		Where("id in (?)", id).
		Find(&l).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	var out domain.AggreList
	linq.From(l).SelectT(model.Aggre{}.ModelToDomain).ToSlice(&out)
	return out, nil
}
