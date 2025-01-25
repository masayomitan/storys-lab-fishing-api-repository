package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
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


// 管理者用の「すべての魚を取得」ルート
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
		// 管理者向けにすべての魚を取得するアクションを実行します。
		act.FindAllByAdmin(c)
	}
}

// 管理者用の「魚を作成」ルート
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
		// 管理者向けに魚を作成するアクションを実行します。
		act.CreateByAdmin(c)
	}
}

// 管理者用の「魚を削除」ルート
func (g ginEngine) buildDeleteFishAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishAdminUsecase.NewFishAdminInteractor(
				fishAdminRepository.NewFishSQL(g.db),
				fishAdminPresenter.NewFishPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFishAdminAction(uc, g.log)
		)
		// 管理者向けに魚を削除するアクションを実行します。
		act.DeleteByAdmin(c)
	}
}

// 管理者用の「すべての魚カテゴリを取得」ルート
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
		// 管理者向けにすべての魚カテゴリを取得するアクションを実行します。
		act.FindAllByAdmin(c)
	}
}

// 管理者用の「魚カテゴリを作成」ルート
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
		// 管理者向けに魚カテゴリを作成するアクションを実行します。
		act.CreateByAdmin(c)
	}
}

//管理者用の「魚カテゴリを更新」ルート
// func (g ginEngine) buildUpdateFishCategoryAdminRoute() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var (
// 			uc = fishCategoriesAdminUsecase.NewFishCategoryAdminInteractor(
// 				fishCategoriesAdminRepository.NewFishCategorySQL(g.db),
// 				fishCategoriesAdminPresenter.NewFishCategoryPresenter(),
// 				g.ctxTimeout,
// 			)
// 			act = action.NewUpdateFishCategoryAdminAction(uc, g.log, &g.validator)
// 		)
// 		// 管理者向けに魚カテゴリを更新するアクションを実行します。
// 		act.UpdateByAdmin(c)
// 	}
// }

// 管理者用の「魚カテゴリを削除」ルート
func (g ginEngine) buildDeleteFishCategoryAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishCategoriesAdminUsecase.NewFishCategoryAdminInteractor(
				fishCategoriesAdminRepository.NewFishCategorySQL(g.db),
				fishCategoriesAdminPresenter.NewFishCategoryPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFishCategoryAdminAction(uc, g.log)
		)
		// 管理者向けに魚カテゴリを削除するアクションを実行します。
		act.DeleteByAdmin(c)
	}
}

// 管理者用の「画像を更新」ルート
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
		bucketName := os.Getenv("AWS_S3_BUCKET")

		var (
			uc = imageAdminUsecase.NewUploadImageAdminInteractor(
				imageAdminRepository.NewImageSQL(g.db),
				imageAdminPresenter.NewImagePresenter(),
				service.NewS3Uploader(s3Client, bucketName),
				g.ctxTimeout,
			)
			act = action.NewUploadImageAdminAction(uc, g.log, &g.validator)
		)

		// 管理者向けに画像を更新するアクションを実行します。
		act.UploadByAdmin(c)
	}
}

// 管理者用の「すべての画像を取得」ルート
func (g ginEngine) buildFindAllImagesAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = imageAdminUsecase.NewImageAdminInteractor(
				imageAdminRepository.NewImageSQL(g.db),
				imageAdminPresenter.NewImagePresenter(),
				g.ctxTimeout,
			)
			act = action.NewImageAdminAction(uc, g.log)
		)
		// 管理者向けにすべての画像を取得するアクションを実行します。
		act.FindAllByAdmin(c)
	}
}
