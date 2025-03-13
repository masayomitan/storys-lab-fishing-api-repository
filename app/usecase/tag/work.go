package usecase

import (
	"context"
	"time"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type TagUseCase interface {
	// FindExecute(ctx context.Context) ([]domain.Tag, error)
}

type TagAdminUseCase interface {
	FindExecuteByAdmin(ctx context.Context) ([]domain.Tag, error)
}

type TagPresenter interface {
	Present(domain.Tag) domain.Tag
	Presents([]domain.Tag) []domain.Tag
}

// type tagInteractor struct {
// 	repo       repository.TagRepository
// 	presenter  TagPresenter
// 	ctxTimeout time.Duration
// }

type tagAdminInteractor struct {
	repo       repository.TagAdminRepository
	presenter  TagPresenter
	ctxTimeout time.Duration
}

func NewTagAdminInteractor(
	repo repository.TagAdminRepository,
	presenter TagPresenter,
	timeout time.Duration,
) TagAdminUseCase {
	return &tagAdminInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}
