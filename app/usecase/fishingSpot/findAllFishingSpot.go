package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type (
	FindAllFishingSpotUseCase interface {
		Execute(context.Context, int) ([]domain.FishingSpot, error)
	}

	FindAllFishingSpotPresenter interface {
		Output([]domain.FishingSpot) []domain.FishingSpot
	}

	findAllFishingSpotInteractor struct {
		repo       repository.FishingSpotRepository
		presenter  FindAllFishingSpotPresenter
		ctxTimeout time.Duration
	}
)

func NewFindAllFishingSpotInteractor(
	repo repository.FishingSpotRepository,
	presenter FindAllFishingSpotPresenter,
	t time.Duration,
) FindAllFishingSpotUseCase {
	return findAllFishingSpotInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (t findAllFishingSpotInteractor) Execute(ctx context.Context, id int) ([]domain.FishingSpot, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fishingSpot, err := t.repo.FindAllByAreaId(ctx, id)
	if err != nil {
		return t.presenter.Output([]domain.FishingSpot{}), err
	}

	return t.presenter.Output(fishingSpot), nil
}
