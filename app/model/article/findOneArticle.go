package model

import (
	"context"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
	"fmt"
)

func (a ArticleSQL) FindOne(ctx context.Context, id int) (domain.Article, error) {
	var json = domain.Article{}

	if err := a.db.FindOne(ctx, a.tableName, id, &json); err != nil {
		return domain.Article{}, errors.Wrap(err, "error listing areas")
	}

	// Instructor のロード
	if json.InstructorID != 0 {
		instructor := domain.Instructor{}
		if err := a.db.DB.Table("instructors").Where("id = ?", json.InstructorID).First(&instructor).Error; err != nil {
			return domain.Article{}, errors.Wrap(err, "error loading instructor")
		}
		json.Instructor = instructor
	}

	// Admin のロード
	if json.AdminID != 0 {
		admin := domain.Admin{}
		if err := a.db.DB.Table("admins").Where("id = ?", json.AdminID).First(&admin).Error; err != nil {
			return domain.Article{}, errors.Wrap(err, "error loading admin")
		}
		json.Admin = admin
	}

	var article = domain.NewArticle(
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
	fmt.Printf("PublishedDateTime: %v\n", article.PublishedDatetime)

	return article, nil
}

func (ga *GormAdapter) FindOne(ctx context.Context, table string, id int, result interface{}) error {
	return ga.DB.Table(table).Where("id = ?", id).
		Preload("ArticleImages").
		First(result).Error
}

