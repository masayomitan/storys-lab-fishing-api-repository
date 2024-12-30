package usecase

import (
	"context"
	"time"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type FishCategoryAdminUseCase interface {
	FindOneByAdmin(ctx context.Context, id int) (domain.FishCategory, error)
	FindAllByAdmin(ctx context.Context) ([]domain.FishCategory, error)
	CreateByAdmin(ctx context.Context, requestParam domain.FishCategory)  (domain.FishCategory, error)
}

type FishCategoryPresenter interface {
	PresentOne(domain.FishCategory) domain.FishCategory
	PresentAll([]domain.FishCategory) []domain.FishCategory
}

type fishCategoryAdminInteractor struct {
	repo       repository.FishCategoryAdminRepository
	presenter  FishCategoryPresenter
	ctxTimeout time.Duration
}

func NewFishCategoryAdminInteractor(
	repo repository.FishCategoryAdminRepository,
	presenter FishCategoryPresenter,
	timeout time.Duration,
) FishCategoryAdminUseCase {
	return &fishCategoryAdminInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}

func (t fishCategoryAdminInteractor) FindOneByAdmin(ctx context.Context, id int) (domain.FishCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fish, err := t.repo.FindOneByAdmin(ctx, id)
	if err != nil {
		return t.presenter.PresentOne(domain.FishCategory{}), err
	}

	return t.presenter.PresentOne(fish), nil
}

func (t fishCategoryAdminInteractor) FindAllByAdmin(ctx context.Context) ([]domain.FishCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fishes, err := t.repo.FindAllByAdmin(ctx)
	if err != nil {
		return t.presenter.PresentAll([]domain.FishCategory{}), err
	}

	return t.presenter.PresentAll(fishes), nil
}

func (t fishCategoryAdminInteractor) CreateByAdmin(ctx context.Context, requestParam domain.FishCategory) (domain.FishCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fishCategory, err := t.repo.CreateByAdmin(ctx, requestParam)
	if err != nil {
		return t.presenter.PresentOne(domain.FishCategory{}), err
	}

	return t.presenter.PresentOne(fishCategory), nil
}
