package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type (
	FindAllAreaUseCase interface {
		Execute(context.Context) ([]domain.Area, error)
	}

	// FindAllAreaPresenter output port
	FindAllAreaPresenter interface {
		Output([]domain.Area) []domain.Area
	}

	findAllAreaInteractor struct {
		repo       repository.AreaRepository
		presenter  FindAllAreaPresenter
		ctxTimeout time.Duration
	}
)

func NewFindAllAreaInteractor(
	repo repository.AreaRepository,
	presenter FindAllAreaPresenter,
	t time.Duration,
) FindAllAreaUseCase {
	return findAllAreaInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (t findAllAreaInteractor) Execute(ctx context.Context) ([]domain.Area, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	Areas, err := t.repo.FindAll(ctx)
	if err != nil {
		return t.presenter.Output([]domain.Area{}), err
	}

	return t.presenter.Output(Areas), nil
}
