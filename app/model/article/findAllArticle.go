package model

import (
	"fmt"
	"context"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
)

func (a ArticleSQL) FindAll(ctx context.Context) ([]domain.Article, error) {
	var json = make([]domain.Article, 0)

	// まず記事一覧を取得
	if err := a.db.FindAll(ctx, a.tableName, domain.Article{}, &json); err != nil {
		return []domain.Article{}, errors.Wrap(err, "error listing articles")
	}

	// 各記事についてInstructorとAdminをロード
	for i := range json {
		article := &json[i]

		// Instructor のロード
		if article.InstructorID != 0 {
			instructor := domain.Instructor{}
			if err := a.db.DB.Table("instructors").Where("id = ?", article.InstructorID).First(&instructor).Error; err != nil {
				return nil, errors.Wrap(err, "error loading instructor")
			}
			article.Instructor = instructor
		}

		// Admin のロード
		if article.AdminID != 0 {
			admin := domain.Admin{}
			if err := a.db.DB.Table("admins").Where("id = ?", article.AdminID).First(&admin).Error; err != nil {
				return nil, errors.Wrap(err, "error loading admin")
			}
			article.Admin = admin
		}
	}

	// 新しい構造体に変換
	var articles = make([]domain.Article, len(json))
	for i, json := range json {
		articles[i] = domain.NewArticle(
			json.ID,
			json.Title,
			json.SubTitle,
			json.InstructorID,
			json.AdminID,
			json.Description,
			json.IsDisplay,
			json.PublishedDatetime,
			json.ArticleCategoryID,
			json.ViewCount,
			json.ArticleCategory,
			json.ArticleImages,
			json.Instructor,
			json.Admin,
		)
	}
    fmt.Println(articles)
	return articles, nil
}

func (a ArticleSQL) FindAllByArticleCategoryId(ctx context.Context, category_id int) ([]domain.Article, error) {
	var json = make([]domain.Article, 0)

	// カテゴリIDに基づいて記事一覧を取得
	if err := a.db.FindAllByArticleCategoryId(ctx, a.tableName, category_id, &json); err != nil {
		return []domain.Article{}, errors.Wrap(err, "error listing articles")
	}

	// 各記事についてInstructorとAdminをロード
	for i := range json {
		article := &json[i]

		// Instructor のロード
		if article.InstructorID != 0 {
			instructor := domain.Instructor{}
			if err := a.db.DB.Table("instructors").Where("id = ?", article.InstructorID).First(&instructor).Error; err != nil {
				return nil, errors.Wrap(err, "error loading instructor")
			}
			article.Instructor = instructor
		}

		// Admin のロード
		if article.AdminID != 0 {
			admin := domain.Admin{}
			if err := a.db.DB.Table("admins").Where("id = ?", article.AdminID).First(&admin).Error; err != nil {
				return nil, errors.Wrap(err, "error loading admin")
			}
			article.Admin = admin
		}
	}

	// 新しい構造体に変換
	var articles = make([]domain.Article, len(json))
	for i, json := range json {
		articles[i] = domain.NewArticle(
			json.ID,
			json.Title,
			json.SubTitle,
			json.InstructorID,
			json.AdminID,
			json.Description,
			json.IsDisplay,
			json.PublishedDatetime,
			json.ArticleCategoryID,
			json.ViewCount,
			convertArticle(json.ArticleCategory),
			json.ArticleImages,
			json.Instructor,
			json.Admin,
		)
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
