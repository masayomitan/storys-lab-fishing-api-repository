package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	ImageAdminRepository interface {
		// FindOneByAdmin(context.Context, int) (domain.Image, error)
		// FindAllByAdmin(context.Context) ([]domain.Image, error)
		UpdateByAdmin(context.Context, domain.Image) (domain.Image, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}
)
