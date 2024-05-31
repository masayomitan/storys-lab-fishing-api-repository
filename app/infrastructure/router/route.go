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

	r.GET("/fishes", g.buildFindAllFishAction())
	r.GET("/fishes/:id", g.buildFindOneFishAction())

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
