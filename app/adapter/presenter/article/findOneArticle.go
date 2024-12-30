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
		ID: 				article.ID,
		Title:      	 	article.Title,
		SubTitle:  		  	article.SubTitle,
		InstructorID: 		article.InstructorID,
		AdminID:     		article.AdminID,
		Description: 		article.Description,
		IsDisplay:   		article.IsDisplay,
		PublishedDatetime: 	article.PublishedDatetime,
		ArticleCategoryID:  article.ArticleCategoryID,
		ViewCount:   		article.ViewCount,
		ArticleImages: 		article.ArticleImages,
		Instructor:         article.Instructor,
		Admin:              article.Admin,
	}
}

