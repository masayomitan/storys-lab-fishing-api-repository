package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	FishRepository interface {
		// Create(context.Context, domain.Fish) (domain.Fish, error)
		FindOne(context.Context, int) (domain.Fish, error)
		FindAll(context.Context) ([]domain.Fish, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}

	FishAdminRepository interface {
		FindOneByAdmin(context.Context, int) (domain.Fish, error)
		FindAllByAdmin(context.Context) ([]domain.Fish, error)
		CreateByAdmin(context.Context, domain.Fish) (domain.Fish, error)
		UpdateByAdmin(context.Context, domain.Fish, int) (domain.Fish, error)
		DeleteByAdmin(context.Context, int) error
		UpdateFishDishesByAdmin(ctx context.Context, fishID int, dishIDs []int) (domain.FishDishRelationResult, error)

		// WithTransaction(context.Context, func(context.Context) error) error
	}
)
