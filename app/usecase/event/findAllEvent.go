package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type (
	FindAllEventUseCase interface {
		Execute(context.Context) ([]domain.Event, error)
	}

	FindAllEventPresenter interface {
		Output([]domain.Event) []domain.Event
	}

	findAllEventInteractor struct {
		repo       repository.EventRepository
		presenter  FindAllEventPresenter
		ctxTimeout time.Duration
	}
)

func NewFindAllEventInteractor(
	repo repository.EventRepository,
	presenter FindAllEventPresenter,
	t time.Duration,
) FindAllEventUseCase {
	return findAllEventInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (t findAllEventInteractor) Execute(ctx context.Context) ([]domain.Event, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	Events, err := t.repo.FindAll(ctx)
	if err != nil {
		return t.presenter.Output([]domain.Event{}), err
	}

	return t.presenter.Output(Events), nil
}

