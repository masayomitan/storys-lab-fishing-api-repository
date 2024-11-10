package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	ArticleRepository interface {
		// Create(context.Context, domain.Article) (domain.Article, error)
		FindAll(context.Context) ([]domain.Article, error)
		FindAllByArticleCategoryId(context.Context, int) ([]domain.Article, error)
		// FindOne(context.Context, int) (domain.Article, error)
	}
)
