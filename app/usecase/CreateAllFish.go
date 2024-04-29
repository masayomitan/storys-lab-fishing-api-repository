package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
)

type (
	// CreateAccountUseCase input port
	CreateFishUseCase interface {
		Execute(context.Context, CreateFishInput) (CreateFishOutput, error)
	}

	// CreateFishInput input data
	CreateFishInput struct {
		ID string `json:"id"`
		Name string `json:"name"`
		FamilyName string `json:"family_name"`
		ScientificName string `json:"scientific_name"`
		FishCategory int `json:"fish_category"`
		Description string `json:"description"`
	}

	// CreateFishPresenter output port
	CreateFishPresenter interface {
		Output(domain.Fish) CreateFishOutput
	}

	// CreateFishOutput output data
	CreateFishOutput struct {
		ID string `json:"id"`
		Name string `json:"name"`
		FamilyName string `json:"family_name"`
		ScientificName string `json:"scientific_name"`
		FishCategory int `json:"fish_category"`
		Description string `json:"description"`
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
func (interactor createFishInteractor) Execute(ctx context.Context, input CreateFishInput) (CreateFishOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, interactor.ctxTimeout)
	defer cancel()

	var fish = domain.NewFish(
		domain.FishID(domain.NewUUID()),
		input.Name,
		input.FamilyName,
		input.ScientificName,
		input.FishCategory,
		input.Description,
	)

	fish, err := interactor.repo.Create(ctx, fish)
	if err != nil {
		return interactor.presenter.Output(domain.Fish{}), err
	}

	return interactor.presenter.Output(fish), nil
}
