package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"storys-lab-fishing-api/app/adapter/api/action"
	"storys-lab-fishing-api/app/adapter/logger"
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
)

type ginEngine struct {
	router *gin.Engine
	log    logger.Logger
	db     *gorm.DB
	// validator  validator.Validator
	port       Port
	ctxTimeout time.Duration
}

func (g ginEngine) setAppHandlers(r *gin.Engine) {

	r.Static("/public", "./public")

	// 魚
	r.GET("/fishes", g.buildFindAllFishAction())
	r.GET("/fishes/:id", g.buildFindOneFishAction())

	// 都道府県
	r.GET("/prefectures/:id", g.buildFindOnePrefAction())

	// エリア
	r.GET("/areas/:id", g.buildFindOneAreaAction())

	// 釣り場
	r.GET("/fishing-spots/:id", g.buildFindOneFishingSpotAction())
	r.GET("/fishing-spots/area/:area_id", g.buildFindAllFishingSpotByAreaIdAction())

	// 道具
	r.GET("/tools/:id", g.buildFindOneToolAction())

	// 道具種別
	r.GET("/tool-categories", g.buildFindAllToolCategoryAction())

	// 記事
	r.GET("/articles", g.buildFindAllArticleAction())


	// ヘルスチェック
	r.GET("/health", g.healthCheck())

	// マイグレーション TODO 他に移行するべき
	r.GET("/migration", g.migration())
}

func (g ginEngine) buildFindOneFishAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishUsecase.NewFindOneFishInteractor(
				fishRepository.NewFishOneSQL(g.db),
				fishPresenter.NewFindOneFishPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindOneFishAction(uc, g.log)
		)
		act.FindOne(c)
	}
}

func (g ginEngine) buildFindAllFishAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishUsecase.NewFindAllFishInteractor(
				fishRepository.NewFishSQL(g.db),
				fishPresenter.NewFindAllFishPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindAllFishAction(uc, g.log)
		)
		act.FindAll(c)
	}
}

func (g ginEngine) buildFindOneAreaAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = areaUsecase.NewFindOneAreaInteractor(
				areaRepository.NewAreaSQL(g.db),
				areaPresenter.NewFindOneAreaPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindOneAreaAction(uc, g.log)
		)
		act.FindOne(c)
	}
}

func (g ginEngine) buildFindOnePrefAction() gin.HandlerFunc {
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

func (g ginEngine) buildFindOneFishingSpotAction() gin.HandlerFunc {
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

func (g ginEngine) buildFindAllFishingSpotByAreaIdAction() gin.HandlerFunc {
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

func (g ginEngine) buildFindOneToolAction() gin.HandlerFunc {
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

func (g ginEngine) buildFindAllToolCategoryAction() gin.HandlerFunc {
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

func (g ginEngine) healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c)
	}
}

func (g ginEngine) migration() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.Migration(c)
	}
}
