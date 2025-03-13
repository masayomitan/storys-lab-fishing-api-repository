package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (

	FishingSpotRepository interface {
		Find(context.Context) ([]domain.FishingSpot, error)
		FindByAreaId(context.Context, int) ([]domain.FishingSpot, error)
		FindOne(context.Context, int) (domain.FishingSpot, error)
	}

	FishingSpotAdminRepository interface {
		FindByAdmin(context.Context) ([]domain.FishingSpot, error)
		FindOneByAdmin(context.Context, int) (domain.FishingSpot, error)
		CreateByAdmin(context.Context, domain.FishingSpot) (domain.FishingSpot, error)
		UpdateByAdmin(context.Context, domain.FishingSpot, int) (domain.FishingSpot, error)
		DeleteByAdmin(context.Context, int) error
		// WithTransaction(context.Context, func(context.Context) error) error
	}
)
