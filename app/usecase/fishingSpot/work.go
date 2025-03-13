package usecase

import (
	"context"
	"time"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)


type FishingSpotUseCase interface {
	FindOneExecute(ctx context.Context, id int) (domain.FishingSpot, error)
	FindByAreaIdExecute(ctx context.Context, area_id int) ([]domain.FishingSpot, error)
	FindAllExecute(ctx context.Context) ([]domain.FishingSpot, error)
	// Create(ctx context.Context, requestParam domain.Fish)  (domain.Fish, error)
}

type FishingSpotAdminUseCase interface {
	FindExecuteByAdmin(ctx context.Context) ([]domain.FishingSpot, error)
	FindOneExecuteByAdmin(ctx context.Context, id int) (domain.FishingSpot, error)
	CreateExecuteByAdmin(ctx context.Context, requestParam domain.FishingSpot)  (domain.FishingSpot, error)
	UpdateExecuteByAdmin(ctx context.Context, requestParam domain.FishingSpot, id int)  (domain.FishingSpot, error)
	DeleteExecuteByAdmin(ctx context.Context, id int) error
}

type FishingSpotPresenter interface {
	PresentOne(domain.FishingSpot) domain.FishingSpot
	PresentAll([]domain.FishingSpot) []domain.FishingSpot
}

type FishingSpotInteractor struct {
	repo       repository.FishingSpotRepository
	presenter  FishingSpotPresenter
	ctxTimeout time.Duration
}
type FishingSpotAdminInteractor struct {
	repo       repository.FishingSpotAdminRepository
	presenter  FishingSpotPresenter
	ctxTimeout time.Duration
}

func NewFishingSpotInteractor(
	repo 		repository.FishingSpotRepository,
	presenter 	FishingSpotPresenter,
	timeout 	time.Duration,
) FishingSpotUseCase {
	return &FishingSpotInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}

func NewFishingSpotAdminInteractor(
	repo 		repository.FishingSpotAdminRepository,
	presenter 	FishingSpotPresenter,
	timeout 	time.Duration,
) FishingSpotAdminUseCase {
	return &FishingSpotAdminInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}
