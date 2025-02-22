package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	AreaRepository interface {
		// Create(context.Context, domain.Area) (domain.Area, error)
		FindAll(context.Context) ([]domain.Area, error)
		FindOne(context.Context, int) (domain.Area, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}

	AreaAdminRepository interface {
		FindByAdmin(context.Context) ([]domain.Area, error)
		FindOneByAdmin(context.Context, int) (domain.Area, error)
		CreateByAdmin(context.Context, domain.Area) (domain.Area, error)
		// UpdateByAdmin(context.Context, domain.Area, int) (domain.Area, error)
		// DeleteByAdmin(context.Context, int) error
		// WithTransaction(context.Context, func(context.Context) error) error
	}
)
