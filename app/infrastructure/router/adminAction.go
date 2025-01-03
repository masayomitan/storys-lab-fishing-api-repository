package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	// "gorm.io/gorm"
	"storys-lab-fishing-api/app/adapter/api/action"

	fishAdminPresenter "storys-lab-fishing-api/app/adapter/presenter/fish"
	fishAdminRepository "storys-lab-fishing-api/app/model/fish"
	fishAdminUsecase "storys-lab-fishing-api/app/usecase/fish"

	fishCategoriesAdminPresenter "storys-lab-fishing-api/app/adapter/presenter/fishCategory"
	fishCategoriesAdminRepository "storys-lab-fishing-api/app/model/fishCategory"
	fishCategoriesAdminUsecase "storys-lab-fishing-api/app/usecase/fishCategory"

	imageAdminPresenter "storys-lab-fishing-api/app/adapter/presenter/image"
	imageAdminRepository "storys-lab-fishing-api/app/model/image"
	imageAdminUsecase "storys-lab-fishing-api/app/usecase/image"


	service "storys-lab-fishing-api/app/service"
)

func (g ginEngine) buildFindAllFishesAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishAdminUsecase.NewFishAdminInteractor(
				fishAdminRepository.NewFishSQL(g.db),
				fishAdminPresenter.NewFishPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFishAdminAction(uc, g.log)
		)
		act.FindAllByAdmin(c)
	}
}

func (g ginEngine) buildCreateFishAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishAdminUsecase.NewFishAdminInteractor(
				fishAdminRepository.NewFishSQL(g.db),
				fishAdminPresenter.NewFishPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFishCreateAdminAction(uc, g.log, &g.validator)
		)
		act.CreateByAdmin(c)
	}
}

func (g ginEngine) buildFindAllFishCategoriesAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishCategoriesAdminUsecase.NewFishCategoryAdminInteractor(
				fishCategoriesAdminRepository.NewFishCategorySQL(g.db),
				fishCategoriesAdminPresenter.NewFishCategoryPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFishCategoryAdminAction(uc, g.log)
		)
		act.FindAllByAdmin(c)
	}
}

func (g ginEngine) buildCreateFishCategoryAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishCategoriesAdminUsecase.NewFishCategoryAdminInteractor(
				fishCategoriesAdminRepository.NewFishCategorySQL(g.db),
				fishCategoriesAdminPresenter.NewFishCategoryPresenter(),
				g.ctxTimeout,
			)
			act = action.NewCreateFishCategoryAdminAction(uc, g.log, &g.validator)
		)
		act.CreateByAdmin(c)
	}
}


func (g ginEngine) buildUpdateImageAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		// AWS Configをロード
		cfg, err := awsConfig.LoadDefaultConfig(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load AWS config"})
			return
		}

		// S3クライアントを作成
		s3Client := s3.NewFromConfig(cfg)

		// S3アップローダーを初期化
		bucketName := "your-s3-bucket-name" // 必要に応じて環境変数や設定ファイルから取得

		var (
			uc = imageAdminUsecase.NewImageAdminInteractor(
				imageAdminRepository.NewImageSQL(g.db),
				imageAdminPresenter.NewImagePresenter(),
				service.NewS3Uploader(s3Client, bucketName),
				g.ctxTimeout,
			)
			act = action.NewUpdateImageAdminAction(uc, g.log, &g.validator)
		)

		// アクションを実行
		act.UpdateByAdmin(c)
	}
}

