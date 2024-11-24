package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type (
	FindAllToolUseCase interface {
		Execute(context.Context) ([]domain.Tool, error)
	}

	FindAllToolPresenter interface {
		Output([]domain.Tool) []domain.Tool
	}

	findAllToolInteractor struct {
		repo       repository.ToolRepository
		presenter  FindAllToolPresenter
		ctxTimeout time.Duration
	}
)

func NewFindAllToolInteractor(
	repo repository.ToolRepository,
	presenter FindAllToolPresenter,
	t time.Duration,
) FindAllToolUseCase {
	return findAllToolInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (t findAllToolInteractor) Execute(ctx context.Context) ([]domain.Tool, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	tools, err := t.repo.FindAll(ctx)
	if err != nil {
		return t.presenter.Output([]domain.Tool{}), err
	}

	return t.presenter.Output(tools), nil
}
