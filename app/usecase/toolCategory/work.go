package usecase

import (
	"context"
	"time"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type ToolCategoryUseCase interface {
	FindOneExecute(ctx context.Context, id int) (domain.ToolCategory, error)
	FindExecute(ctx context.Context) ([]domain.ToolCategory, error)
	// Create(ctx context.Context, requestParam domain.Tool)  (domain.Tool, error)
}

type ToolCategoryAdminUseCase interface {
	FindOneExecuteByAdmin(ctx context.Context, id int) (domain.ToolCategory, error)
	FindAllExecuteByAdmin(ctx context.Context) ([]domain.ToolCategory, error)
	CreateExecuteByAdmin(ctx context.Context, requestParam domain.ToolCategory)  (domain.ToolCategory, error)
	UpdateExecuteByAdmin(ctx context.Context, requestParam domain.ToolCategory, id int)  (domain.ToolCategory, error)
	DeleteExecuteByAdmin(ctx context.Context, id int) error
}

type ToolCategoryPresenter interface {
	Present(domain.ToolCategory) domain.ToolCategory
	Presents([]domain.ToolCategory) []domain.ToolCategory
}

type toolCategoryInteractor struct {
	repo       repository.ToolCategoryRepository
	presenter  ToolCategoryPresenter
	ctxTimeout time.Duration
}


type toolCategoryAdminInteractor struct {
	repo       repository.ToolCategoryAdminRepository
	presenter  ToolCategoryPresenter
	ctxTimeout time.Duration
}

func NewToolCategoryInteractor(
	repo repository.ToolCategoryRepository,
	presenter ToolCategoryPresenter,
	timeout time.Duration,
) ToolCategoryUseCase {
	return &toolCategoryInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}


func NewToolCategoryAdminInteractor(
	repo repository.ToolCategoryAdminRepository,
	presenter ToolCategoryPresenter,
	timeout time.Duration,
) ToolCategoryAdminUseCase {
	return &toolCategoryAdminInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}
