package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	FishRepository interface {
		Create(context.Context, domain.Fish) (domain.Fish, error)
		FindAll(context.Context) ([]domain.Fish, error)
		FindOne(context.Context, string) (domain.Fish, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}
)
