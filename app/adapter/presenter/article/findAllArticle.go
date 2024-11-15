package presenter

import (
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/article"
)

type findAllArticlePresenter struct{}

func NewFindAllArticlePresenter() usecase.FindAllArticlePresenter {
	return findAllArticlePresenter{}
}

func (a findAllArticlePresenter) Output(articles []domain.Article) []domain.Article {
	var output = make([]domain.Article, 0)
	for _, article := range articles {
		output = append(output, domain.Article{
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
			ArticleCategory: article.ArticleCategory,
		})
	}
	return output
}
