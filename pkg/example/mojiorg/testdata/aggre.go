package testdata

func Aggre(id ...int) *domain.Aggre {
	obj := &domain.Aggre{
		Id:         0,
		TenantId:   2,
		CompanyId:  3,
		Status:     4,
		Ct:         5,
		Mt:         6,
		CreatorId:  7,
		ModifierId: 8,
	}
	if len(id) > 0 {
		obj.Id = id[0]
	}
	return obj
}
