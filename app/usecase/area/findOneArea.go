package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type (
	// FindOneAreaUseCase input port
	FindOneAreaUseCase interface {
		Execute(context.Context, string) (domain.Area, error)
	}

	FindOneAreaPresenter interface {
		Output(domain.Area) domain.Area
	}

	findOneAreaInteractor struct {
		repo       repository.AreaRepository
		presenter  FindOneAreaPresenter
		ctxTimeout time.Duration
	}
)

func NewFindOneAreaInteractor(
	repo repository.AreaRepository,
	presenter FindOneAreaPresenter,
	t time.Duration,
) FindOneAreaUseCase {
	return findOneAreaInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (t findOneAreaInteractor) Execute(ctx context.Context, id string) (domain.Area, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	Area, err := t.repo.FindOne(ctx, id)
	if err != nil {
		return t.presenter.Output(domain.Area{}), err
	}

	return t.presenter.Output(Area), nil
}
