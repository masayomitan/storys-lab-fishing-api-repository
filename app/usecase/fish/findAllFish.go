package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type (
	// FindAllFishUseCase input port
	FindAllFishUseCase interface {
		Execute(context.Context) ([]domain.Fish, error)
	}

	// FindAllFishPresenter output port
	FindAllFishPresenter interface {
		Output([]domain.Fish) []domain.Fish
	}

	findAllFishInteractor struct {
		repo       repository.FishRepository
		presenter  FindAllFishPresenter
		ctxTimeout time.Duration
	}
)

// NewFindAllFishInteractor creates new findAllFishInteractor with its dependencies
func NewFindAllFishInteractor(
	repo repository.FishRepository,
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
func (t findAllFishInteractor) Execute(ctx context.Context) ([]domain.Fish, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	Fishes, err := t.repo.FindAll(ctx)
	if err != nil {
		return t.presenter.Output([]domain.Fish{}), err
	}

	return t.presenter.Output(Fishes), nil
}
