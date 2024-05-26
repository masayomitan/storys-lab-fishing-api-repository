package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
)

type (
	// FindAllFishUseCase input port
	FindAllFishUseCase interface {
		Execute(context.Context) ([]domain.FishStruct, error)
	}

	// FindAllFishPresenter output port
	FindAllFishPresenter interface {
		Output([]domain.Fish) []domain.FishStruct
	}

	findAllFishInteractor struct {
		repo       domain.FishRepository
		presenter  FindAllFishPresenter
		ctxTimeout time.Duration
	}
)

// NewFindAllFishInteractor creates new findAllFishInteractor with its dependencies
func NewFindAllFishInteractor(
	repo domain.FishRepository,
	presenter FindAllFishPresenter,
	t time.Duration,
) FindAllFishUseCase {
	return findAllFishInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (t findAllFishInteractor) Execute(ctx context.Context) ([]domain.FishStruct, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	Fishes, err := t.repo.FindAll(ctx)
	if err != nil {
		return t.presenter.Output([]domain.Fish{}), err
	}

	return t.presenter.Output(Fishes), nil
}
