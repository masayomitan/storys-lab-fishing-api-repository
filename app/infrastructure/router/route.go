package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"storys-lab-fishing-api/app/adapter/api/action"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/adapter/validator"
)

type ginEngine struct {
	router *gin.Engine
	log    logger.Logger
	db     *gorm.DB
	validator  validator.Validator
	port       Port
	ctxTimeout time.Duration
}

func (g ginEngine) setAppHandlers(r *gin.Engine) {
	// `/api` プレフィックスを追加
	api := r.Group("/api")

	// 静的ファイル
	api.Static("/public", "./public")

	// 魚
	api.GET("/fishes", g.buildFindAllFishRoute())
	api.GET("/fishes/:id", g.buildFindOneFishRoute())

	// 都道府県
	api.GET("/prefectures/:id", g.buildFindOnePrefRoute())

	// エリア
	api.GET("/areas/:id", g.buildFindOneAreaRoute())

	// 釣り場
	api.GET("/fishing-spots/:id", g.buildFindOneFishingSpotRoute())
	api.GET("/fishing-spots/area/:area_id", g.buildFindAllFishingSpotByAreaIdRoute())
	// todo
	// api.GET("/areas/:area_id/fishing-spots", g.buildFindAllFishingSpotByAreaIdRoute())

	// 道具
	api.GET("/tools", g.buildFindAllToolRoute())
	api.GET("/tools/:id", g.buildFindOneToolRoute())
	api.GET("/tool-categories/:id/tools", g.buildFindAllToolCategoryRoute())

	// 道具種別
	api.GET("/tool-categories", g.buildFindAllToolCategoryRoute())

	// 記事
	api.GET("/articles", g.buildFindAllArticleRoute())
	api.GET("/articles/:id", g.buildFindOneArticleRoute())
	api.GET("/article-categories/:id", g.buildFindAllArticleByArticleCategoryIdRoute())

	// イベント
	api.GET("/events", g.buildFindAllEventRoute())
	api.GET("/events/:id", g.buildFindOneEventRoute())

	// ヘルスチェック
	api.GET("/health", g.healthCheck())

	// マイグレーション TODO 他に移行するべき
	api.GET("/migration", g.migration())

	//////// 管理側
	mc := api.Group("/admin")

	// /mc/fishes サブグループ
	// fishes := mc.Group("/fishes")
	// fishes.GET("/", g.buildFindAllFishesAdminRoute())
	// fishes.GET("/create", g.buildCreateFishRoute())
	// fishes.GET("/update/:id", g.buildUpdateFishRoute())

	// /mc/fish-categories サブグループ
	fishCategories := mc.Group("/fish-categories")
	fishCategories.GET("/", g.buildFindAllFishCategoriesAdminRoute())
	fishCategories.POST("/create", g.buildCreateFishCategoryAdminRoute())
	// fishCategories.GET("/update/:id", g.buildUpdateFishCategoryRoute())
	
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
