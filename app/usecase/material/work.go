package usecase

import (
	"context"
	"time"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type MaterialUseCase interface {
	// FindExecute(ctx context.Context) ([]domain.Material, error)
}

type MaterialAdminUseCase interface {
	FindExecuteByAdmin(ctx context.Context) ([]domain.Material, error)
}

type MaterialPresenter interface {
	Present(domain.Material) domain.Material
	Presents([]domain.Material) []domain.Material
}

// type MaterialInteractor struct {
// 	repo       repository.MaterialRepository
// 	presenter  MaterialPresenter
// 	ctxTimeout time.Duration
// }

type materialAdminInteractor struct {
	repo       repository.MaterialAdminRepository
	presenter  MaterialPresenter
	ctxTimeout time.Duration
}

func NewMaterialAdminInteractor(
	repo repository.MaterialAdminRepository,
	presenter MaterialPresenter,
	timeout time.Duration,
) MaterialAdminUseCase {
	return &materialAdminInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}
