package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
)

type (
	// FindOneFishUseCase input port
	FindOneFishUseCase interface {
		Execute(context.Context, string) (domain.FishStruct, error)
	}

	// FindOneFishPresenter output port
	FindOneFishPresenter interface {
		Output(domain.Fish) domain.FishStruct
	}

	findOneFishInteractor struct {
		repo       domain.FishRepository
		presenter  FindOneFishPresenter
		ctxTimeout time.Duration
	}
)

// NewFindOneFishInteractor creates new findOneFishInteractor with its dependencies
func NewFindOneFishInteractor(
	repo domain.FishRepository,
	presenter FindOneFishPresenter,
	t time.Duration,
) FindOneFishUseCase {
	return findOneFishInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (t findOneFishInteractor) Execute(ctx context.Context, id string) (domain.FishStruct, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	Fish, err := t.repo.FindOne(ctx, id)
	if err != nil {
		return t.presenter.Output(domain.Fish{}), err
	}

	return t.presenter.Output(Fish), nil
}
