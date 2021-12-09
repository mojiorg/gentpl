package model

import (
	"baibao/settlement/pkg/domain"
	"github.com/jinzhu/copier"
	"github.com/quexer/utee"
)

type Aggre struct {
	Id         int
	TenantId   int
	CompanyId  int
	Status     int
	Ct         int
	Mt         int
	CreatorId  int
	ModifierId int
}

func (Aggre) ModelToDomain(x *Aggre) *domain.Aggre {
	return x.ToDomain()
}

func (Aggre) New(x *domain.Aggre) *Aggre {
	out := &Aggre{}
	utee.Chk(copier.Copy(out, x))
	return out
}

func (p *Aggre) ToDomain() *domain.Aggre {
	out := &domain.Aggre{}
	utee.Chk(copier.Copy(out, p))
	return out
}
