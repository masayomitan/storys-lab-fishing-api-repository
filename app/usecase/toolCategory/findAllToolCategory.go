package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type (
	FindAllToolCategoryUseCase interface {
		Execute(context.Context) ([]domain.ToolCategory, error)
	}

	FindAllToolCategoryPresenter interface {
		Output([]domain.ToolCategory) []domain.ToolCategory
	}

	findAllToolCategoryInteractor struct {
		repo       repository.ToolCategoryRepository
		presenter  FindAllToolCategoryPresenter
		ctxTimeout time.Duration
	}
)

func NewFindAllToolCategoryInteractor(
	repo repository.ToolCategoryRepository,
	presenter FindAllToolCategoryPresenter,
	t time.Duration,
) FindAllToolCategoryUseCase {
	return findAllToolCategoryInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (t findAllToolCategoryInteractor) Execute(ctx context.Context) ([]domain.ToolCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	ToolCategory, err := t.repo.FindAll(ctx)
	if err != nil {
		return t.presenter.Output([]domain.ToolCategory{}), err
	}

	return t.presenter.Output(ToolCategory), nil
}
