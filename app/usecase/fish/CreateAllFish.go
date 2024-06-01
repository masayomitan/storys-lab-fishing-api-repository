package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type (
	// CreateAccountUseCase input port
	CreateFishUseCase interface {
		Execute(context.Context, domain.Fish) (domain.Fish, error)
	}

	// CreateFishPresenter output port
	CreateFishPresenter interface {
		Output(domain.Fish) domain.Fish
	}

	createFishInteractor struct {
		repo       repository.FishRepository
		presenter  CreateFishPresenter
		ctxTimeout time.Duration
	}
)

// NewCreateFishInteractor creates new createFishInteractor with its dependencies
func NewCreateFishInteractor(
	repo repository.FishRepository,
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
func (interactor createFishInteractor) Execute(ctx context.Context, input domain.Fish) (domain.Fish, error) {
	ctx, cancel := context.WithTimeout(ctx, interactor.ctxTimeout)
	defer cancel()

	var fish = domain.NewFish(
		domain.NewUUID(),
		input.Name,
		input.FamilyName,
		input.ScientificName,
		input.FishCategoryId,
		input.Description,
		input.Length,
		input.Weight,
		input.Habitat,
		input.DepthRange,
		input.WaterTemperatureRange,
		input.ConservationStatus,
		input.FishCategory,
		input.FishingMethods,
		input.Dishes,
		
	)

	fish, err := interactor.repo.Create(ctx, fish)
	if err != nil {
		return interactor.presenter.Output(domain.Fish{}), err
	}

	return interactor.presenter.Output(fish), nil
}
