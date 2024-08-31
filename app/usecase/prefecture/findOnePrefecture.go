package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type (
	FindOnePrefUseCase interface {
		Execute(context.Context, string) (domain.Pref, error)
	}

	FindOnePrefPresenter interface {
		Output(domain.Pref) domain.Pref
	}

	findOnePrefInteractor struct {
		repo       repository.PrefRepository
		presenter  FindOnePrefPresenter
		ctxTimeout time.Duration
	}
)

func NewFindOnePrefInteractor(
	repo repository.PrefRepository,
	presenter FindOnePrefPresenter,
	t time.Duration,
) FindOnePrefUseCase {
	return findOnePrefInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (t findOnePrefInteractor) Execute(ctx context.Context, id string) (domain.Pref, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	pref, err := t.repo.FindOne(ctx, id)
	if err != nil {
		return t.presenter.Output(domain.Pref{}), err
	}

	return t.presenter.Output(pref), nil
}
