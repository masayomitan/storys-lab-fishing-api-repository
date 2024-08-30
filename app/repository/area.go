package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	AreaRepository interface {
		Create(context.Context, domain.Area) (domain.Area, error)
		FindAll(context.Context) ([]domain.Area, error)
		// FindOne(context.Context, string) (domain.Area, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}
)
