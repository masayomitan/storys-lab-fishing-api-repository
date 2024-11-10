package router

import (
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	"storys-lab-fishing-api/app/adapter/api/action"

	articlePresenter "storys-lab-fishing-api/app/adapter/presenter/article"
	articleRepository "storys-lab-fishing-api/app/model/article"
	articleUsecase "storys-lab-fishing-api/app/usecase/article"
)

func (g ginEngine) buildFindAllArticleAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = articleUsecase.NewFindAllArticleInteractor(
				articleRepository.NewArticleSQL(g.db),
				articlePresenter.NewFindAllArticlePresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindAllArticleAction(uc, g.log)
		)
		act.FindAll(c)
	}
}
