package router

import (
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	"storys-lab-fishing-api/app/adapter/api/action"

	// fishAdminPresenter "storys-lab-fishing-api/app/adapter/presenter/fish"
	// fishAdminRepository "storys-lab-fishing-api/app/model/fish"
	// fishAdminUsecase "storys-lab-fishing-api/app/usecase/fish"

	fishCategoriesAdminPresenter "storys-lab-fishing-api/app/adapter/presenter/fishCategory"
	fishCategoriesAdminRepository "storys-lab-fishing-api/app/model/fishCategory"
	fishCategoriesAdminUsecase "storys-lab-fishing-api/app/usecase/fishCategory"
)

// func (g ginEngine) buildFindAllFishesAdminRoute() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var (
// 			uc = fishAdminUsecase.NewFishAdminInteractor(
// 				fishAdminRepository.NewFishSQL(g.db),
// 				fishAdminPresenter.NewFishPresenter(),
// 				g.ctxTimeout,
// 			)
// 			act = action.NewFindFishCategoryAdminAction(uc, g.log)
// 		)
// 		act.FindAllByAdmin(c)
// 	}
// }

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
			act = action.NewFishCategoryCreateAdminAction(uc, g.log, &g.validator)
		)
		act.CreateByAdmin(c)
	}
}
