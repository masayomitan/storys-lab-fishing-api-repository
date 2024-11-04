package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type (
	FindOneToolUseCase interface {
		Execute(context.Context, int) (domain.Tool, error)
	}

	FindOneToolPresenter interface {
		Output(domain.Tool) domain.Tool
	}

	findOneToolInteractor struct {
		repo       repository.ToolRepository
		presenter  FindOneToolPresenter
		ctxTimeout time.Duration
	}
)

func NewFindOneToolInteractor(
	repo repository.ToolRepository,
	presenter FindOneToolPresenter,
	t time.Duration,
) FindOneToolUseCase {
	return findOneToolInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (t findOneToolInteractor) Execute(ctx context.Context, id int) (domain.Tool, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	tool, err := t.repo.FindOne(ctx, id)
	if err != nil {
		return t.presenter.Output(domain.Tool{}), err
	}

	return t.presenter.Output(tool), nil
}
