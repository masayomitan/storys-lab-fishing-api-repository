package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
)

type (
	// FindOneFishUseCase input port
	FindOneFishUseCase interface {
		Execute(context.Context, string) (FindOneFishOutput, error)
	}

	// FindOneFishPresenter output port
	FindOneFishPresenter interface {
		Output(domain.Fish) FindOneFishOutput
	}

	// FindOneFishOutput output data
	FindOneFishOutput struct {
		ID string  `json:"id"`
		Name string  `json:"name"`
		FamilyName string `json:"family_name"`
		ScientificName string `json:"scientific_name"`
		FishCategoryId int `json:"fish_category_id"`
		Description string `json:"description"`
	}

	findOneFishInteractor struct {
		repo       domain.FishRepository
		presenter  FindOneFishPresenter
		ctxTimeout time.Duration
	}
)

// NewFindOneFishInteractor creates new findOneFishInteractor with its dependencies
func NewFindOneFishInteractor(
	repo domain.FishRepository,
	presenter FindOneFishPresenter,
	t time.Duration,
) FindOneFishUseCase {
	return findOneFishInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (t findOneFishInteractor) Execute(ctx context.Context, id string) (FindOneFishOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	Fish, err := t.repo.FindOne(ctx, id)
	if err != nil {
		return t.presenter.Output(domain.Fish{}), err
	}

	return t.presenter.Output(Fish), nil
}
