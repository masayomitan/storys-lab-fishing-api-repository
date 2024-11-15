package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type (
	FindOneArticleUseCase interface {
		Execute(context.Context, int) (domain.Article, error)
	}

	FindOneArticlePresenter interface {
		Output(domain.Article) domain.Article
	}

	findOneArticleInteractor struct {
		repo       repository.ArticleRepository
		presenter  FindOneArticlePresenter
		ctxTimeout time.Duration
	}
)

func NewFindOneArticleInteractor(
	repo repository.ArticleRepository,
	presenter FindOneArticlePresenter,
	t time.Duration,
) FindOneArticleUseCase {
	return findOneArticleInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (t findOneArticleInteractor) Execute(ctx context.Context, id int) (domain.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	article, err := t.repo.FindOne(ctx, id)
	if err != nil {
		return t.presenter.Output(domain.Article{}), err
	}

	return t.presenter.Output(article), nil
}
