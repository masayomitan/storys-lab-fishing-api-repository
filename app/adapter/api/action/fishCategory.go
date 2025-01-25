package action

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/adapter/validator"
	"storys-lab-fishing-api/app/usecase/fishCategory"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/utils"
)

type FishCategoryAdminAction struct {
	uc  usecase.FishCategoryAdminUseCase
	log logger.Logger
}

type FishCategoryCreateAdminAction struct {
	uc  usecase.FishCategoryAdminUseCase
	log logger.Logger
	val *validator.Validator
}

func NewFishCategoryAdminAction(uc usecase.FishCategoryAdminUseCase, log logger.Logger) FishCategoryAdminAction {
	return FishCategoryAdminAction{
		uc:  uc,
		log: log,
	}
}

func NewCreateFishCategoryAdminAction(uc usecase.FishCategoryAdminUseCase, log logger.Logger, val *validator.Validator) FishCategoryCreateAdminAction {
	return FishCategoryCreateAdminAction{
		uc:  uc,
		log: log,
		val: val,
	}
}

// func (t FindOneToolAction) FindOneAdmin(c *gin.Context) {
// 	const logKey = "find_one_Tool"
// 	fmt.Println("")
// 	output, err := t.uc.Execute(c.Request.Context(), utils.StrToInt(c.Param("id")))
// 	if err != nil {
// 		logging.NewError(
// 			t.log,
// 			err,
// 			logKey,
// 			http.StatusInternalServerError,
// 		).Log("error when returning the Tool")

// 		return
// 	}
// 	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning Tool")

// 	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
// }

func (t FishCategoryAdminAction) FindAllByAdmin(c *gin.Context) {
	const logKey = "find_one_fish_category"
	fmt.Println("")
	output, err := t.uc.FindAllExecuteByAdmin(c.Request.Context())
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the fish-category")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning fish-category")
	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}

// 魚カテゴリーデータ作成アクション
func (t FishCategoryCreateAdminAction) CreateByAdmin(c *gin.Context) {
	const logKey = "create_fish_category"
	fmt.Println("")

	var requestParam domain.FishCategory
	if err := c.ShouldBindJSON(&requestParam); err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusUnprocessableEntity,
		).Log("invalid request payload")
		response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
		return
	}
	fmt.Println(requestParam)

	if err := t.val.ValidateStruct(requestParam); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	output, err := t.uc.CreateExecuteByAdmin(c.Request.Context(), requestParam)
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the fish-category")
		response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning fish-category")
	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}

// 魚データ更新アクション
func (t FishCategoryCreateAdminAction) UpdateByAdmin(c *gin.Context, id int) {
	const logKey = "create_fish_category"
	fmt.Println("")

	var requestParam domain.FishCategory
	if err := c.ShouldBindJSON(&requestParam); err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusUnprocessableEntity,
		).Log("invalid request payload")
		response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
		return
	}
	fmt.Println(requestParam)

	if err := t.val.ValidateStruct(requestParam); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	output, err := t.uc.UpdateExecuteByAdmin(c.Request.Context(), requestParam, id)
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the fish-category")
		response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning fish-category")
	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}

func (t FishCategoryAdminAction) DeleteByAdmin(c *gin.Context) {
    const logKey = "delete_fish_category"
    fmt.Println("")

    // パスパラメータからIDを取得
    id := utils.StrToInt(c.Param("id"))
    if id == 0 {
        err := fmt.Errorf("id is required")
        logging.NewError(
            t.log,
            err,
            logKey,
            http.StatusBadRequest,
        ).Log("missing id in request")
        response.NewError(err, http.StatusBadRequest).Send(c.Writer)
        return
    }

    // UseCaseを呼び出して削除処理を実行
    err := t.uc.DeleteExecuteByAdmin(c.Request.Context(), id)
    if err != nil {
        logging.NewError(
            t.log,
            err,
            logKey,
            http.StatusInternalServerError,
        ).Log("error when deleting the fish-category")
        response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
        return
    }

    logging.NewInfo(t.log, logKey, http.StatusOK).Log("successfully deleted fish-category")
    response.NewSuccess(nil, http.StatusOK).Send(c.Writer)
}
