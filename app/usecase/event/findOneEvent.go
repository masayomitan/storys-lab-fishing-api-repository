package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type (
	FindOneEventUseCase interface {
		Execute(context.Context, int) (domain.Event, error)
	}

	FindOneEventPresenter interface {
		Output(domain.Event) domain.Event
	}

	findOneEventInteractor struct {
		repo       repository.EventRepository
		presenter  FindOneEventPresenter
		ctxTimeout time.Duration
	}
)

func NewFindOneEventInteractor(
	repo repository.EventRepository,
	presenter FindOneEventPresenter,
	t time.Duration,
) FindOneEventUseCase {
	return findOneEventInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (t findOneEventInteractor) Execute(ctx context.Context, id int) (domain.Event, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	event, err := t.repo.FindOne(ctx, id)
	if err != nil {
		return t.presenter.Output(domain.Event{}), err
	}

	return t.presenter.Output(event), nil
}
