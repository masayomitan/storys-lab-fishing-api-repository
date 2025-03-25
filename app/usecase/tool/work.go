package usecase

import (
	"context"
	"time"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type ToolUseCase interface {
	FindOneExecute(ctx context.Context, id int) (domain.Tool, error)
	FindExecute(ctx context.Context) ([]domain.Tool, error)
	// Create(ctx context.Context, requestParam domain.Tool)  (domain.Tool, error)
}

type ToolAdminUseCase interface {
	FindOneExecuteByAdmin(ctx context.Context, id int) (domain.Tool, error)
	FindAllExecuteByAdmin(ctx context.Context) ([]domain.Tool, error)
	CreateExecuteByAdmin(ctx context.Context, requestParam domain.Tool)  (domain.Tool, error)
	UpdateExecuteByAdmin(ctx context.Context, requestParam domain.Tool, id int)  (domain.Tool, error)
	DeleteExecuteByAdmin(ctx context.Context, id int) error
}

type ToolPresenter interface {
	Present(domain.Tool) domain.Tool
	Presents([]domain.Tool) []domain.Tool
}

type toolInteractor struct {
	repo       repository.ToolRepository
	presenter  ToolPresenter
	ctxTimeout time.Duration
}

type toolAdminInteractor struct {
	repo       repository.ToolAdminRepository
	presenter  ToolPresenter
	ctxTimeout time.Duration
}

func NewToolInteractor(
	repo repository.ToolRepository,
	presenter ToolPresenter,
	timeout time.Duration,
) ToolUseCase {
	return &toolInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}

func NewToolAdminInteractor(
	repo repository.ToolAdminRepository,
	presenter ToolPresenter,
	timeout time.Duration,
) ToolAdminUseCase {
	return &toolAdminInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}
