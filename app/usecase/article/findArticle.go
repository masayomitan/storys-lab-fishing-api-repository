package usecase

import (
	"context"
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type (
	FindAllArticleUseCase interface {
		Execute(context.Context) ([]domain.Article, error)
		ExecuteSecond(context.Context, int) ([]domain.Article, error)
	}

	FindAllArticleByArticleCategoryIDUseCase interface {
		ExecuteSecond(context.Context, int) ([]domain.Article, error)
	}

	FindAllArticlePresenter interface {
		Output([]domain.Article) []domain.Article
	}

	findAllArticleInteractor struct {
		repo       repository.ArticleRepository
		presenter  FindAllArticlePresenter
		ctxTimeout time.Duration
	}
)

func NewFindAllArticleInteractor(
	repo repository.ArticleRepository,
	presenter FindAllArticlePresenter,
	t time.Duration,
) FindAllArticleUseCase {
	return findAllArticleInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (t findAllArticleInteractor) Execute(ctx context.Context) ([]domain.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	articles, err := t.repo.FindAll(ctx)
	if err != nil {
		return t.presenter.Output([]domain.Article{}), err
	}

	return t.presenter.Output(articles), nil
}

func (t findAllArticleInteractor) ExecuteSecond(ctx context.Context, id int) ([]domain.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	articles, err := t.repo.FindAllByArticleCategoryId(ctx, id)
	if err != nil {
		return t.presenter.Output([]domain.Article{}), err
	}

	return t.presenter.Output(articles), nil
}
