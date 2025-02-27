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

	prefectureAdminPresenter "storys-lab-fishing-api/app/adapter/presenter/prefecture"
	prefectureAdminRepository "storys-lab-fishing-api/app/model/prefecture"
	prefectureAdminUsecase "storys-lab-fishing-api/app/usecase/prefecture"

	areaAdminPresenter "storys-lab-fishing-api/app/adapter/presenter/area"
	areaAdminRepository "storys-lab-fishing-api/app/model/area"
	areaAdminUsecase "storys-lab-fishing-api/app/usecase/area"

	imageAdminPresenter "storys-lab-fishing-api/app/adapter/presenter/image"
	imageAdminRepository "storys-lab-fishing-api/app/model/image"
	imageAdminUsecase "storys-lab-fishing-api/app/usecase/image"

	service "storys-lab-fishing-api/app/service"
)

// ******************** 
// 魚
// ********************
// 魚取得
func (g ginEngine) buildFindOneFishAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishAdminUsecase.NewFishAdminInteractor(
				fishAdminRepository.NewFishSQL(g.db),
				fishAdminPresenter.NewFishPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFishAdminAction(uc, g.log)
		)
		act.FindOneByAdmin(c)
	}
}


// 魚取得
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

// 魚作成
func (g ginEngine) buildCreateFishAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishAdminUsecase.NewFishAdminInteractor(
				fishAdminRepository.NewFishSQL(g.db),
				fishAdminPresenter.NewFishPresenter(),
				g.ctxTimeout,
			)
			act = action.NewMutationFishAdminAction(uc, g.log, &g.validator)
		)
		// 管理者向けに魚を作成するアクションを実行します。
		act.CreateByAdmin(c)
	}
}

// 魚更新
func (g ginEngine) buildUpdateFishAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishAdminUsecase.NewFishAdminInteractor(
				fishAdminRepository.NewFishSQL(g.db),
				fishAdminPresenter.NewFishPresenter(),
				g.ctxTimeout,
			)
			act = action.NewMutationFishAdminAction(uc, g.log, &g.validator)
		)
		act.UpdateByAdmin(c)
	}
}

// 魚削除
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
		act.DeleteByAdmin(c)
	}
}

// ******************** 
// 魚カテゴリー
// ********************
// 魚カテゴリー取得
func (g ginEngine) buildFindOneFishCategoriesAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishCategoriesAdminUsecase.NewFishCategoryAdminInteractor(
				fishCategoriesAdminRepository.NewFishCategorySQL(g.db),
				fishCategoriesAdminPresenter.NewFishCategoryPresenter(),
				g.ctxTimeout,
			)
			act = action.NewFishCategoryAdminAction(uc, g.log)
		)
		act.FindOneByAdmin(c)
	}
}

// すべての魚カテゴリー取得
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

// 魚カテゴリー作成
func (g ginEngine) buildCreateFishCategoryAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishCategoriesAdminUsecase.NewFishCategoryAdminInteractor(
				fishCategoriesAdminRepository.NewFishCategorySQL(g.db),
				fishCategoriesAdminPresenter.NewFishCategoryPresenter(),
				g.ctxTimeout,
			)
			act = action.NewMutationFishCategoryAdminAction(uc, g.log, &g.validator)
		)
		act.CreateByAdmin(c)
	}
}

// 魚カテゴリー更新
func (g ginEngine) buildUpdateFishCategoryAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = fishCategoriesAdminUsecase.NewFishCategoryAdminInteractor(
				fishCategoriesAdminRepository.NewFishCategorySQL(g.db),
				fishCategoriesAdminPresenter.NewFishCategoryPresenter(),
				g.ctxTimeout,
			)
			act = action.NewMutationFishCategoryAdminAction(uc, g.log, &g.validator)
		)
		act.UpdateByAdmin(c)
	}
}

// 魚カテゴリー削除
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
		act.DeleteByAdmin(c)
	}
}

// ******************** 
// 画像
// ********************
// 画像更新
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
		act.UploadByAdmin(c)
	}
}

// 画像取得
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
		act.FindAllByAdmin(c)
	}
}

// ******************** 
// 都道府県
// ********************
func (g ginEngine) buildFindPrefecturesAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = prefectureAdminUsecase.NewPrefectureAdminInteractor(
				prefectureAdminRepository.NewPrefectureSQL(g.db),
				prefectureAdminPresenter.NewPrefecturePresenter(),
				g.ctxTimeout,
			)
			act = action.NewPrefectureAdminAction(uc, g.log)
		)
		act.FindByAdmin(c)
	}
}

// ******************** 
// エリア
// ********************
func (g ginEngine) buildFindAreasAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = areaAdminUsecase.NewAreaAdminInteractor(
				areaAdminRepository.NewAreaSQL(g.db),
				areaAdminPresenter.NewAreaPresenter(),
				g.ctxTimeout,
			)
			act = action.NewAreaAdminAction(uc, g.log)
		)
		act.FindByAdmin(c)
	}
}

func (g ginEngine) buildFindOneAreaAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = areaAdminUsecase.NewAreaAdminInteractor(
				areaAdminRepository.NewAreaSQL(g.db),
				areaAdminPresenter.NewAreaPresenter(),
				g.ctxTimeout,
			)
			act = action.NewAreaAdminAction(uc, g.log)
		)
		act.FindOneByAdmin(c)
	}
}

func (g ginEngine) buildCreateAreaAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = areaAdminUsecase.NewAreaAdminInteractor(
				areaAdminRepository.NewAreaSQL(g.db),
				areaAdminPresenter.NewAreaPresenter(),
				g.ctxTimeout,
			)
			act = action.NewMutationAreaAdminAction(uc, g.log, &g.validator)
		)
		act.CreateByAdmin(c)
	}
}

func (g ginEngine) buildDeleteAreaAdminRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = areaAdminUsecase.NewAreaAdminInteractor(
				areaAdminRepository.NewAreaSQL(g.db),
				areaAdminPresenter.NewAreaPresenter(),
				g.ctxTimeout,
			)
			act = action.NewAreaAdminAction(uc, g.log)
		)
		act.DeleteByAdmin(c)
	}
}
