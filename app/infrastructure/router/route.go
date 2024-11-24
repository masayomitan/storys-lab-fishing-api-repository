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
	r.GET("/fishes", g.buildFindAllFishRoute())
	r.GET("/fishes/:id", g.buildFindOneFishRoute())

	// 都道府県
	r.GET("/prefectures/:id", g.buildFindOnePrefRoute())

	// エリア
	r.GET("/areas/:id", g.buildFindOneAreaRoute())

	// 釣り場
	r.GET("/fishing-spots/:id", g.buildFindOneFishingSpotRoute())
	r.GET("/fishing-spots/area/:area_id", g.buildFindAllFishingSpotByAreaIdRoute())
	// todo
	// r.GET("/areas/:area_id/fishing-spots", g.buildFindAllFishingSpotByAreaIdRoute())

	// 道具
	r.GET("/tools", g.buildFindAllToolRoute())
	r.GET("/tools/:id", g.buildFindOneToolRoute())
	r.GET("/tool-categories/:id/tools", g.buildFindAllToolCategoryRoute())

	// 道具種別
	r.GET("/tool-categories", g.buildFindAllToolCategoryRoute())

	// 記事
	r.GET("/articles", g.buildFindAllArticleRoute())
	r.GET("/articles/:id", g.buildFindOneArticleRoute())
	r.GET("/article-categories/:id", g.buildFindAllArticleByArticleCategoryIdRoute())

	// ヘルスチェック
	r.GET("/health", g.healthCheck())

	// マイグレーション TODO 他に移行するべき
	r.GET("/migration", g.migration())
}

func (g ginEngine) buildFindOneFishRoute() gin.HandlerFunc {
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

func (g ginEngine) buildFindAllFishRoute() gin.HandlerFunc {
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

func (g ginEngine) buildFindOneAreaRoute() gin.HandlerFunc {
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
