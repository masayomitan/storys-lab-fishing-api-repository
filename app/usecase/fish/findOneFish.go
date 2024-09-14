package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type (
	FindOneFishUseCase interface {
		Execute(context.Context, string) (domain.Fish, error)
	}

	FindOneFishPresenter interface {
		Output(domain.Fish) domain.Fish
	}

	findOneFishInteractor struct {
		repo       repository.FishRepository
		presenter  FindOneFishPresenter
		ctxTimeout time.Duration
	}
)

func NewFindOneFishInteractor(
	repo repository.FishRepository,
	presenter FindOneFishPresenter,
	t time.Duration,
) FindOneFishUseCase {
	return findOneFishInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (t findOneFishInteractor) Execute(ctx context.Context, id string) (domain.Fish, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fish, err := t.repo.FindOne(ctx, id)
	if err != nil {
		return t.presenter.Output(domain.Fish{}), err
	}

	return t.presenter.Output(fish), nil
}
