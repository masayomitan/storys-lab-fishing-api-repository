package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type (
	FindAllFishUseCase interface {
		Execute(context.Context) ([]domain.Fish, error)
	}

	FindAllFishPresenter interface {
		Output([]domain.Fish) []domain.Fish
	}

	findAllFishInteractor struct {
		repo       repository.FishRepository
		presenter  FindAllFishPresenter
		ctxTimeout time.Duration
	}
)

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

func (t findAllFishInteractor) Execute(ctx context.Context) ([]domain.Fish, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fishes, err := t.repo.FindAll(ctx)
	if err != nil {
		return t.presenter.Output([]domain.Fish{}), err
	}

	return t.presenter.Output(fishes), nil
}
