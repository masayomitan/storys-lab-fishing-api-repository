package usecase

import (
	"context"
	"time"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type DishUseCase interface {
	FindOneExecute(ctx context.Context, id int) (domain.Fish, error)
	FindExecute(ctx context.Context) ([]domain.Fish, error)
	// Create(ctx context.Context, requestParam domain.Fish)  (domain.Fish, error)
}

type DishAdminUseCase interface {
	FindOneExecuteByAdmin(ctx context.Context, id int) (domain.Dish, error)
	FindExecuteByAdmin(ctx context.Context) ([]domain.Dish, error)
	CreateExecuteByAdmin(ctx context.Context, requestParam domain.Dish)  (domain.Dish, error)
	UpdateExecuteByAdmin(ctx context.Context, requestParam domain.Dish, id int)  (domain.Dish, error)
	DeleteExecuteByAdmin(ctx context.Context, id int) error
}

type DishPresenter interface {
	Present(domain.Dish) domain.Dish
	Presents([]domain.Dish) []domain.Dish
}

type dishInteractor struct {
	repo       repository.DishRepository
	presenter  DishPresenter
	ctxTimeout time.Duration
}

type dishAdminInteractor struct {
	repo       repository.DishAdminRepository
	presenter  DishPresenter
	ctxTimeout time.Duration
}

func NewDishAdminInteractor(
	repo repository.DishAdminRepository,
	presenter DishPresenter,
	timeout time.Duration,
) DishAdminUseCase {
	return &dishAdminInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}
