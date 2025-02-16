package usecase

import (
	"context"
	"time"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type FishUseCase interface {
	FindOneExecute(ctx context.Context, id int) (domain.Fish, error)
	FindAllExecute(ctx context.Context) ([]domain.Fish, error)
	// Create(ctx context.Context, requestParam domain.Fish)  (domain.Fish, error)
}

type FishAdminUseCase interface {
	FindOneExecuteByAdmin(ctx context.Context, id int) (domain.Fish, error)
	FindAllExecuteByAdmin(ctx context.Context) ([]domain.Fish, error)
	CreateExecuteByAdmin(ctx context.Context, requestParam domain.Fish)  (domain.Fish, error)
	UpdateExecuteByAdmin(ctx context.Context, requestParam domain.Fish, id int)  (domain.Fish, error)
	DeleteExecuteByAdmin(ctx context.Context, id int) error
}

type FishPresenter interface {
	PresentOne(domain.Fish) domain.Fish
	PresentAll([]domain.Fish) []domain.Fish
}

type fishInteractor struct {
	repo       repository.FishRepository
	presenter  FishPresenter
	ctxTimeout time.Duration
}

type fishAdminInteractor struct {
	repo       repository.FishAdminRepository
	presenter  FishPresenter
	ctxTimeout time.Duration
}

func NewFishInteractor(
	repo repository.FishRepository,
	presenter FishPresenter,
	t time.Duration,
) FishUseCase {
	return &fishInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func NewFishAdminInteractor(
	repo repository.FishAdminRepository,
	presenter FishPresenter,
	timeout time.Duration,
) FishAdminUseCase {
	return &fishAdminInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}
