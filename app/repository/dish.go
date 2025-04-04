package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	DishRepository interface {
		// Create(context.Context, domain.Dish) (domain.Dish, error)
		FindOne(context.Context, int) (domain.Dish, error)
		Find(context.Context) ([]domain.Dish, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}

	DishAdminRepository interface {
		FindByAdmin(context.Context) ([]domain.Dish, error)
		FindOneByAdmin(context.Context, int) (domain.Dish, error)
		CreateByAdmin(context.Context, domain.Dish) (domain.Dish, error)
		UpdateByAdmin(context.Context, domain.Dish, int) (domain.Dish, error)
		DeleteByAdmin(context.Context, int) error
		// WithTransaction(context.Context, func(context.Context) error) error
	}
)
