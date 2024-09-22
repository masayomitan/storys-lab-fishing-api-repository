package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	FishingSpotRepository interface {
		// Create(context.Context, domain.FishingSpot) (domain.FishingSpot, error)
		FindAllByAreaId(context.Context, int) ([]domain.FishingSpot, error)
		FindOne(context.Context, string) (domain.FishingSpot, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}
)
