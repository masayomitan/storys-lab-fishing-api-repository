package router

import (
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	"storys-lab-fishing-api/app/adapter/api/action"

	toolPresenter "storys-lab-fishing-api/app/adapter/presenter/tool"
	toolRepository "storys-lab-fishing-api/app/model/tool"
	toolUsecase "storys-lab-fishing-api/app/usecase/tool"

	toolCategoryPresenter "storys-lab-fishing-api/app/adapter/presenter/toolCategory"
	toolCategoryRepository "storys-lab-fishing-api/app/model/toolCategory"
	toolCategoryUsecase "storys-lab-fishing-api/app/usecase/toolCategory"

	articlePresenter "storys-lab-fishing-api/app/adapter/presenter/article"
	articleRepository "storys-lab-fishing-api/app/model/article"
	articleUsecase "storys-lab-fishing-api/app/usecase/article"
	
)
func (g ginEngine) buildFindAllToolRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = toolUsecase.NewFindAllToolInteractor(
				toolRepository.NewToolSQL(g.db),
				toolPresenter.NewFindAllToolPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindAllToolAction(uc, g.log)
		)
		act.FindAll(c)
	}
}

func (g ginEngine) buildFindOneToolRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = toolUsecase.NewFindOneToolInteractor(
				toolRepository.NewToolSQL(g.db),
				toolPresenter.NewFindOneToolPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindOneToolAction(uc, g.log)
		)
		act.FindOne(c)
	}
}

func (g ginEngine) buildFindAllToolCategoryRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = toolCategoryUsecase.NewFindAllToolCategoryInteractor(
				toolCategoryRepository.NewToolCategorySQL(g.db),
				toolCategoryPresenter.NewFindAllToolCategoryPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindAllToolCategoryAction(uc, g.log)
		)
		act.FindAll(c)
	}
}

func (g ginEngine) buildFindOneArticleRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = articleUsecase.NewFindOneArticleInteractor(
				articleRepository.NewArticleSQL(g.db),
				articlePresenter.NewFindOneArticlePresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindOneArticleAction(uc, g.log)
		)
		act.FindOne(c)
	}
}

func (g ginEngine) buildFindAllArticleRoute() gin.HandlerFunc {
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

func (g ginEngine) buildFindAllArticleByArticleCategoryIdRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = articleUsecase.NewFindAllArticleInteractor(
				articleRepository.NewArticleSQL(g.db),
				articlePresenter.NewFindAllArticlePresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindAllArticleByArticleCategoryIdAction(uc, g.log)
		)
		act.FindAllByArticleCategoryID(c)
	}
}
