package model

import (
	"context"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
)

func (a ArticleSQL) FindAll(ctx context.Context) ([]domain.Article, error) {
	var articleJSON = make([]domain.Article, 0)

	if err := a.db.FindAll(ctx, a.tableName, domain.Article{}, &articleJSON); err != nil {
		return []domain.Article{}, errors.Wrap(err, "error listing areas")
	}

	var articles = make([]domain.Article, 0)

	for _, articleJSON := range articleJSON {
		var article = domain.NewArticle(
			articleJSON.ID,
			articleJSON.Title,
			articleJSON.SubTitle,
			articleJSON.InstructorId,
			articleJSON.AdminId,
			articleJSON.Description,
			articleJSON.IsDisplay,
			articleJSON.PublishedDateTime,
			articleJSON.CategoryId,
			articleJSON.ViewCount,
		)

		articles = append(articles, article)
	}

	return articles, nil
}

func (a ArticleSQL) FindAllByArticleCategoryId(ctx context.Context, category_id int) ([]domain.Article, error) {
	var articleJSON = make([]domain.Article, 0)

	if err := a.db.FindAll(ctx, a.tableName, domain.Article{}, &articleJSON); err != nil {
		return []domain.Article{}, errors.Wrap(err, "error listing areas")
	}

	var articles = make([]domain.Article, 0)

	for _, articleJSON := range articleJSON {
		var article = domain.NewArticle(
			articleJSON.ID,
			articleJSON.Title,
			articleJSON.SubTitle,
			articleJSON.InstructorId,
			articleJSON.AdminId,
			articleJSON.Description,
			articleJSON.IsDisplay,
			articleJSON.PublishedDateTime,
			articleJSON.CategoryId,
			articleJSON.ViewCount,
		)

		articles = append(articles, article)
	}

	return articles, nil
}

func (ga *GormAdapter) FindAll(ctx context.Context, table string, query interface{}, result interface{}) error {
    return ga.DB.Table(table).Where(query).
		Find(result).Error
}



