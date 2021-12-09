package repo

import (
	"context"
)

//go:generate mockgen -destination=../mocks/mrepo/aggre.go -package=mrepo . AggreRepo

type AggreRepo interface {
	MustGet(ctx context.Context, id int) (*domain.Aggre, error)
	MultiGet(ctx context.Context, id ...int) (domain.AggreList, error)
	Create(ctx context.Context, in *domain.Aggre) error
	Update(ctx context.Context, in *domain.Aggre) error
}
