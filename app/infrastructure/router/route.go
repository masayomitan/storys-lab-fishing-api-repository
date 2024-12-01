package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"storys-lab-fishing-api/app/adapter/api/action"
	"storys-lab-fishing-api/app/adapter/logger"
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

	// イベント
	r.GET("/events", g.buildFindAllArticleRoute())
	r.GET("/events/:id", g.buildFindOneArticleRoute())

	// ヘルスチェック
	r.GET("/health", g.healthCheck())

	// マイグレーション TODO 他に移行するべき
	r.GET("/migration", g.migration())
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
