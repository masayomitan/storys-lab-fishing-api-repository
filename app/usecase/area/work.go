package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type AreaUseCase interface {
	FindOneExecute(ctx context.Context, id int) (domain.Area, error)
	FindAllExecute(ctx context.Context) ([]domain.Area, error)
}

type AreaAdminUseCase interface {
	FindExecuteByAdmin(ctx context.Context) ([]domain.Area, error)
	FindOneExecuteByAdmin(ctx context.Context, id int) (domain.Area, error)
	CreateExecuteByAdmin(ctx context.Context, requestParam domain.Area)  (domain.Area, error)
	UpdateExecuteByAdmin(ctx context.Context, requestParam domain.Area, id int)  (domain.Area, error)
	DeleteExecuteByAdmin(ctx context.Context, id int) error
}

type AreaPresenter interface {
	PresentOne(domain.Area) domain.Area
	PresentAll([]domain.Area) []domain.Area
}

type areaInteractor struct {
	repo       repository.AreaRepository
	presenter  AreaPresenter
	ctxTimeout time.Duration
}

type areaAdminInteractor struct {
	repo       repository.AreaAdminRepository
	presenter  AreaPresenter
	ctxTimeout time.Duration
}

func NewAreaInteractor(
	repo repository.AreaRepository,
	presenter AreaPresenter,
	t time.Duration,
) AreaUseCase {
	return areaInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func NewAreaAdminInteractor(
	repo repository.AreaAdminRepository,
	presenter AreaPresenter,
	timeout time.Duration,
) AreaAdminUseCase {
	return &areaAdminInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}
