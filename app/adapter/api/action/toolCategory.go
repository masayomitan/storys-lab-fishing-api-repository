package action

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/toolCategory"
	"storys-lab-fishing-api/app/adapter/validator"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/utils"
)

type ToolCategoryAction struct {
	uc  usecase.ToolCategoryUseCase
	log logger.Logger
}


type ToolCategoryAdminAction struct {
	uc  usecase.ToolCategoryAdminUseCase
	log logger.Logger
}

// MutationFishAdminAction handles fish creation and update actions for admin app.
type MutationToolCategoryAdminAction struct {
	uc  usecase.ToolCategoryAdminUseCase
	log logger.Logger
	val *validator.Validator
}

func NewToolCategoryAction(uc usecase.ToolCategoryUseCase, log logger.Logger) ToolCategoryAction {
	return ToolCategoryAction{
		uc:  uc,
		log: log,
	}
}

func NewToolCategoryAdminAction(uc usecase.ToolCategoryAdminUseCase, log logger.Logger) ToolCategoryAdminAction {
	return ToolCategoryAdminAction{
		uc:  uc,
		log: log,
	}
}

// NewFishCreateAdminAction creates a new instance of MutationFishAdminAction.
//
// Args:
// - uc: Use case interface for fish admin actions.
// - log: Logger for recording events.
// - val: Validator for input validation.
func NewMutationToolCategoryAdminAction(uc usecase.ToolCategoryAdminUseCase, log logger.Logger, val *validator.Validator) MutationToolCategoryAdminAction {
	return MutationToolCategoryAdminAction{
		uc:  uc,
		log: log,
		val: val,
	}
}

// func (t FindOneToolCategoryAction) FindOne(c *gin.Context) {
// 	const logKey = "find_one_Tool"
// 	fmt.Println("")
// 	output, err := t.uc.Execute(c.Request.Context(), c.Param("id"))
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


func (t ToolCategoryAction) Find(c *gin.Context) {
	const logKey = "find_all_tool_categories"
	fmt.Println("")
	output, err := t.uc.FindExecute(c.Request.Context())
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the Tool list")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning Tool list")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}


// FindOneByAdmin retrieves a list of all fishes for admin app.
func (t ToolCategoryAdminAction) FindOneByAdmin(c *gin.Context) {
	const logKey = "find_all_fishes"

	output, err := t.uc.FindOneExecuteByAdmin(c.Request.Context(), utils.StrToInt(c.Param("id")))
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the fish list")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning fish list")
	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}

// FindAllByAdmin retrieves a list of all fishes for admin app.
func (t ToolCategoryAdminAction) FindAllByAdmin(c *gin.Context) {
	const logKey = "find_all_fishes"

	output, err := t.uc.FindAllExecuteByAdmin(c.Request.Context())
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the fish list")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning fish list")
	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}

// CreateByAdmin creates a new fish entry for admin app.
//
// The request payload is validated before creating the fish entry.
func (t MutationToolCategoryAdminAction) CreateByAdmin(c *gin.Context) {
	const logKey = "create_fish"
	fmt.Println("")

	var requestParam domain.ToolCategory
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

// UpdateByAdmin creates a new fish entry for admin app.
//
// The request payload is validated before Updating the fish entry.
func (t MutationToolCategoryAdminAction) UpdateByAdmin(c *gin.Context) {
	const logKey = "create_fish"
	fmt.Println("")

	id := utils.StrToInt(c.Param("id"))
	var requestParam domain.ToolCategory
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


func (t ToolCategoryAdminAction) DeleteByAdmin(c *gin.Context) {
    const logKey = "delete_fish"
	fmt.Println(c.Param("id"))
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
        ).Log("error when deleting the fish")
        response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
        return
    }

    logging.NewInfo(t.log, logKey, http.StatusOK).Log("successfully deleted fish")
    response.NewSuccess(map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Fish deleted successfully",
	}, http.StatusOK).Send(c.Writer)
}
