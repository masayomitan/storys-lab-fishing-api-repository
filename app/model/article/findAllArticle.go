package model

import (
	"context"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
)

func (a ArticleSQL) FindAll(ctx context.Context) ([]domain.Article, error) {
	var json = make([]domain.Article, 0)

	if err := a.db.FindAll(ctx, a.tableName, domain.Article{}, &json); err != nil {
		return []domain.Article{}, errors.Wrap(err, "error listing areas")
	}

	var articles = make([]domain.Article, 0)

	for _, json := range json {
		var article = domain.NewArticle(
			json.ID,
			json.Title,
			json.SubTitle,
			json.InstructorId,
			json.AdminId,
			json.Description,
			json.IsDisplay,
			json.PublishedDatetime,
			json.ArticleCategoryId,
			json.ViewCount,
			json.ArticleCategory,
			json.ArticleImages,
		)

		articles = append(articles, article)
	}

	return articles, nil
}

func (a ArticleSQL) FindAllByArticleCategoryId(ctx context.Context, category_id int) ([]domain.Article, error) {
	var json = make([]domain.Article, 0)

	if err := a.db.FindAllByArticleCategoryId(ctx, a.tableName, category_id, &json); err != nil {
		return []domain.Article{}, errors.Wrap(err, "error listing areas")
	}

	var articles = make([]domain.Article, 0)

	for _, json := range json {
		var article = domain.NewArticle(
			json.ID,
			json.Title,
			json.SubTitle,
			json.InstructorId,
			json.AdminId,
			json.Description,
			json.IsDisplay,
			json.PublishedDatetime,
			json.ArticleCategoryId,
			json.ViewCount,
			convertArticle(json.ArticleCategory),
			json.ArticleImages,
		)
		articles = append(articles, article)
	}

	return articles, nil
}

func (ga *GormAdapter) FindAll(ctx context.Context, table string, query interface{}, result interface{}) error {
    return ga.DB.Table(table).Where(query).
		Preload("ArticleImages").
		Order("id asc").
		Find(result).Error
}

func (ga *GormAdapter) FindAllByArticleCategoryId(ctx context.Context, table string, article_category_id int, result interface{}) error {
	return ga.DB.Table(table).Where("article_category_id = ?", article_category_id).
		Preload("ArticleCategory").
		Preload("ArticleImages").
		Order("id asc").
		Find(result).Error
}
