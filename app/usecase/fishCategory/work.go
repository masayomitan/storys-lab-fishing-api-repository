package usecase

import (
	"context"
	"time"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type FishCategoryAdminUseCase interface {
	FindOneExecuteByAdmin(ctx context.Context, id int) (domain.FishCategory, error)
	FindAllExecuteByAdmin(ctx context.Context) ([]domain.FishCategory, error)
	CreateExecuteByAdmin(ctx context.Context, requestParam domain.FishCategory)  (domain.FishCategory, error)
	UpdateExecuteByAdmin(ctx context.Context, requestParam domain.FishCategory, id int)  (domain.FishCategory, error)
	DeleteExecuteByAdmin(ctx context.Context, id int) error
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
