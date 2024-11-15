package model

import (
	"storys-lab-fishing-api/app/domain"
)

func convertArticle(a domain.ArticleCategory) domain.ArticleCategory {
	return domain.ArticleCategory {
		ID:        a.ID,
		Name:      a.Name,
	}
}
