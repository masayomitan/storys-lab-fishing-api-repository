package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type (
	// CreateAccountUseCase input port
	CreateAreaUseCase interface {
		Execute(context.Context, domain.Area) (domain.Area, error)
	}

	// CreateAreaPresenter output port
	CreateAreaPresenter interface {
		Output(domain.Area) domain.Area
	}

	createAreaInteractor struct {
		repo       repository.AreaRepository
		presenter  CreateAreaPresenter
		ctxTimeout time.Duration
	}
)

// NewCreateAreaInteractor creates new createAreaInteractor with its dependencies
func NewCreateAreaInteractor(
	repo repository.AreaRepository,
	presenter CreateAreaPresenter,
	t time.Duration,
) CreateAreaUseCase {
	return createAreaInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (interactor createAreaInteractor) Execute(ctx context.Context, input domain.Area) (domain.Area, error) {
	ctx, cancel := context.WithTimeout(ctx, interactor.ctxTimeout)
	defer cancel()

	var area = domain.NewArea(
		domain.NewUUID(),
		input.Name,
		input.Description,
		input.PrefectureId,
		input.FishingSpots,
		input.Tides,
	)

	area, err := interactor.repo.Create(ctx, area)
	if err != nil {
		return interactor.presenter.Output(domain.Area{}), err
	}

	return interactor.presenter.Output(area), nil
}
