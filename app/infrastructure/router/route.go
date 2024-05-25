package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/action"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/adapter/presenter"
	"storys-lab-fishing-api/app/adapter/repository"
	"storys-lab-fishing-api/app/usecase"
)

type ginEngine struct {
	router *gin.Engine
	log    logger.Logger
	db     repository.DBMethods
	// validator  validator.Validator
	port       Port
	ctxTimeout time.Duration
}

func (g ginEngine) setAppHandlers(r *gin.Engine) {

	r.GET("/fishes", g.buildFindAllFishAction())

	r.GET("/health", g.healthCheck())

	r.GET("/migration", g.migration())
}

func (g ginEngine) buildFindAllFishAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = usecase.NewFindAllFishInteractor(
				repository.NewFishSQL(g.db),
				presenter.NewFindAllFishPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFindAllFishAction(uc, g.log)
		)

		act.Execute(c.Writer, c.Request)
	}
}

func (g ginEngine) healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}


func (g ginEngine) migration() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.Migration(c.Writer, c.Request)
	}
}
