package presenter

import (
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/article"
)

type findOneArticlePresenter struct{}

func NewFindOneArticlePresenter() usecase.FindOneArticlePresenter {
	return findOneArticlePresenter{}
}

func (a findOneArticlePresenter) Output(article domain.Article) domain.Article {
	return domain.Article {
		ID:          article.ID,
		Title:       article.Title,
		SubTitle:    article.SubTitle,
		InstructorId: article.InstructorId,
		AdminId:     article.AdminId,
		Description: article.Description,
		IsDisplay:   article.IsDisplay,
		PublishedDateTime: article.PublishedDateTime,
		ArticleCategoryId:  article.ArticleCategoryId,
		ViewCount:   article.ViewCount,
	}
}

