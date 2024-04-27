package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
)

type (
	// FindAllFishUseCase input port
	FindAllFishUseCase interface {
		Execute(context.Context) ([]FindAllFishOutput, error)
	}

	// FindAllFishPresenter output port
	FindAllFishPresenter interface {
		Output([]domain.Fish) []FindAllFishOutput
	}

	// FindAllFishOutput output data
	FindAllFishOutput struct {
		ID        string  `json:"id"`
		Name      string  `json:"name"`
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
func (t findAllFishInteractor) Execute(ctx context.Context) ([]FindAllFishOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	Fishes, err := t.repo.FindAll(ctx)
	if err != nil {
		return t.presenter.Output([]domain.Fish{}), err
	}

	return t.presenter.Output(Fishes), nil
}
