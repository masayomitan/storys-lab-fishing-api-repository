package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	FishCategoryAdminRepository interface {
		FindOneByAdmin(context.Context, int) (domain.FishCategory, error)
		FindAllByAdmin(context.Context) ([]domain.FishCategory, error)
		CreateByAdmin(context.Context, domain.FishCategory) (domain.FishCategory, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}
)
