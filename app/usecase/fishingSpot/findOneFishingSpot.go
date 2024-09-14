package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type (
	FindOneFishingSpotUseCase interface {
		Execute(context.Context, string) (domain.FishingSpot, error)
	}

	FindOneFishingSpotPresenter interface {
		Output(domain.FishingSpot) domain.FishingSpot
	}

	findOneFishingSpotInteractor struct {
		repo       repository.FishingSpotRepository
		presenter  FindOneFishingSpotPresenter
		ctxTimeout time.Duration
	}
)

func NewFindOneFishingSpotInteractor(
	repo repository.FishingSpotRepository,
	presenter FindOneFishingSpotPresenter,
	t time.Duration,
) FindOneFishingSpotUseCase {
	return findOneFishingSpotInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (t findOneFishingSpotInteractor) Execute(ctx context.Context, id string) (domain.FishingSpot, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fishingSpot, err := t.repo.FindOne(ctx, id)
	if err != nil {
		return t.presenter.Output(domain.FishingSpot{}), err
	}

	return t.presenter.Output(fishingSpot), nil
}
