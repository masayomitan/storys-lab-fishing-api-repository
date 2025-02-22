package router

import (
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	"storys-lab-fishing-api/app/adapter/api/action"

	fishPresenter "storys-lab-fishing-api/app/adapter/presenter/fish"
	fishRepository "storys-lab-fishing-api/app/model/fish"
	fishUsecase "storys-lab-fishing-api/app/usecase/fish"

	areaPresenter "storys-lab-fishing-api/app/adapter/presenter/area"
	areaRepository "storys-lab-fishing-api/app/model/area"
	areaUsecase"storys-lab-fishing-api/app/usecase/area"

	prefPresenter "storys-lab-fishing-api/app/adapter/presenter/prefecture"
	prefRepository "storys-lab-fishing-api/app/model/prefecture"
	prefUsecase"storys-lab-fishing-api/app/usecase/prefecture"

	fishingSpotPresenter "storys-lab-fishing-api/app/adapter/presenter/fishingSpot"
	fishingSpotRepository "storys-lab-fishing-api/app/model/fishingSpot"
	fishingSpotUsecase "storys-lab-fishing-api/app/usecase/fishingSpot"

	toolPresenter "storys-lab-fishing-api/app/adapter/presenter/tool"
	toolRepository "storys-lab-fishing-api/app/model/tool"
	toolUsecase "storys-lab-fishing-api/app/usecase/tool"

	toolCategoryPresenter "storys-lab-fishing-api/app/adapter/presenter/toolCategory"
	toolCategoryRepository "storys-lab-fishing-api/app/model/toolCategory"
	toolCategoryUsecase "storys-lab-fishing-api/app/usecase/toolCategory"

	articlePresenter "storys-lab-fishing-api/app/adapter/presenter/article"
	articleRepository "storys-lab-fishing-api/app/model/article"
	articleUsecase "storys-lab-fishing-api/app/usecase/article"

	eventPresenter "storys-lab-fishing-api/app/adapter/presenter/event"
	eventRepository "storys-lab-fishing-api/app/model/event"
	eventUsecase "storys-lab-fishing-api/app/usecase/event"
	
)

func (g ginEngine) buildFindOneFishRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishUsecase.NewFishInteractor(
				fishRepository.NewFishOneSQL(g.db),
				fishPresenter.NewFishPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFishAction(uc, g.log)
		)
		act.FindOne(c)
	}
}

func (g ginEngine) buildFindAllFishRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishUsecase.NewFishInteractor(
				fishRepository.NewFishSQL(g.db),
				fishPresenter.NewFishPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFishAction(uc, g.log)
		)
		act.FindAll(c)
	}
}

func (g ginEngine) buildFindOneAreaRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = areaUsecase.NewAreaInteractor(
				areaRepository.NewAreaSQL(g.db),
				areaPresenter.NewAreaPresenter(),
				g.ctxTimeout,
			)
			act = action.NewAreaAction(uc, g.log)
		)
		act.FindOne(c)
	}
}

func (g ginEngine) buildFindOnePrefRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = prefUsecase.NewFindOnePrefInteractor(
				prefRepository.NewPrefSQL(g.db),
				prefPresenter.NewFindOnePrefPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindOnePrefAction(uc, g.log)
		)
		act.FindOne(c)
	}
}

func (g ginEngine) buildFindOneFishingSpotRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishingSpotUsecase.NewFindOneFishingSpotInteractor(
				fishingSpotRepository.NewFishingSpotSQL(g.db),
				fishingSpotPresenter.NewFindOneFishingSpotPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindOneFishingSpotAction(uc, g.log)
		)
		act.FindOne(c)
	}
}

func (g ginEngine) buildFindAllFishingSpotByAreaIdRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishingSpotUsecase.NewFindAllFishingSpotInteractor(
				fishingSpotRepository.NewFishingSpotSQL(g.db),
				fishingSpotPresenter.NewFindAllFishingSpotPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindAllFishingSpotAction(uc, g.log)
		)
		act.FindAllByAreaId(c)
	}
}

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

func (g ginEngine) buildFindOneEventRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = eventUsecase.NewFindOneEventInteractor(
				eventRepository.NewEventSQL(g.db),
				eventPresenter.NewFindOneEventPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindOneEventAction(uc, g.log)
		)
		act.FindOne(c)
	}
}

func (g ginEngine) buildFindAllEventRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = eventUsecase.NewFindAllEventInteractor(
				eventRepository.NewEventSQL(g.db),
				eventPresenter.NewFindAllEventPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindAllEventAction(uc, g.log)
		)
		act.FindAll(c)
	}
}

