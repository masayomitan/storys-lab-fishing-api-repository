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
	// r.GET("/fishing-spots/:id", g.buildFindAllAreaAction())

	r.GET("/health", g.healthCheck())

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
