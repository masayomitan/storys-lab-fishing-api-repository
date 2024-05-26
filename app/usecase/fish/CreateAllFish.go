package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
)

type (
	// CreateAccountUseCase input port
	CreateFishUseCase interface {
		Execute(context.Context, domain.FishStruct) (domain.FishStruct, error)
	}

	// CreateFishPresenter output port
	CreateFishPresenter interface {
		Output(domain.Fish) domain.FishStruct
	}

	createFishInteractor struct {
		repo       domain.FishRepository
		presenter  CreateFishPresenter
		ctxTimeout time.Duration
	}
)

// NewCreateFishInteractor creates new createFishInteractor with its dependencies
func NewCreateFishInteractor(
	repo domain.FishRepository,
	presenter CreateFishPresenter,
	t time.Duration,
) CreateFishUseCase {
	return createFishInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (interactor createFishInteractor) Execute(ctx context.Context, input domain.FishStruct) (domain.FishStruct, error) {
	ctx, cancel := context.WithTimeout(ctx, interactor.ctxTimeout)
	defer cancel()

	var fish = domain.NewFish(
		domain.FishID(domain.NewUUID()),
		input.Name,
		input.FamilyName,
		input.ScientificName,
		input.FishCategoryId,
		input.Description,
		input.Length,
		input.Weight,
		input.Habitat,
		input.Depth_range,
		input.Water_temperature_range,
		input.Conservation_status,
	)

	fish, err := interactor.repo.Create(ctx, fish)
	if err != nil {
		return interactor.presenter.Output(domain.Fish{}), err
	}

	return interactor.presenter.Output(fish), nil
}
