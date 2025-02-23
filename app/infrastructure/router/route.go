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

			// 管理側
			admin := api.Group("/admin")
			{
				// /admin/fishes サブグループ
				fishes := admin.Group("/fishes")
				fishes.GET("/", g.buildFindAllFishesAdminRoute())
				fishes.GET("/:id", g.buildFindOneFishAdminRoute())
				fishes.POST("/create", g.buildCreateFishAdminRoute())
				fishes.PUT("/update/:id", g.buildUpdateFishAdminRoute())
				fishes.DELETE("/delete/:id", g.buildDeleteFishAdminRoute())

				// /admin/fish-categories サブグループ
				fishCategories := admin.Group("/fish-categories")
				fishCategories.GET("/", g.buildFindAllFishCategoriesAdminRoute())
				fishCategories.GET("/:id", g.buildFindOneFishCategoriesAdminRoute())
				fishCategories.POST("/create", g.buildCreateFishCategoryAdminRoute())
				fishCategories.PUT("/update/:id", g.buildUpdateFishCategoryAdminRoute())
				fishCategories.DELETE("/delete/:id", g.buildDeleteFishCategoryAdminRoute())

				// 都道府県
				prefectures := admin.Group("/prefectures")
				prefectures.GET("/", g.buildFindPrefecturesAdminRoute())
				// api.GET("/prefectures/:id", g.buildFindOnePrefRoute())

				areas := admin.Group("/areas")
				areas.GET("/", g.buildFindAreasAdminRoute())
				areas.GET("/:id", g.buildFindOneAreaAdminRoute())
				areas.POST("/create", g.buildCreateAreaAdminRoute())

				// /admin/fish-categories サブグループ
				images := admin.Group("/images")
				images.GET("/", g.buildFindAllImagesAdminRoute())
				images.POST("/upload", g.buildUpdateImageAdminRoute())
				// images.PUT("/upload/:id", g.buildUpdateFishCategoryRoute())
				// images.DELETE("/delete", g.buildDeleteImageAdminRoute())
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
