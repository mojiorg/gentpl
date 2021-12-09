package domain

type Aggre struct {
	Id         int // 系统主键
	TenantId   int `validate:"required"` // 店铺ID
	CompanyId  int `validate:"required"` // 企业ID
	Status     int //
	Ct         int `validate:"required"`
	Mt         int `validate:"required"`
	CreatorId  int // 创建用户ID
	ModifierId int // 修改用户ID
}
