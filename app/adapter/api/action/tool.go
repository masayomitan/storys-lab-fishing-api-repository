package action

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/tool"
	"storys-lab-fishing-api/app/adapter/validator"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/utils"
)

type ToolAction struct {
	uc  usecase.ToolUseCase
	log logger.Logger
}
type ToolAdminAction struct {
	uc  usecase.ToolAdminUseCase
	log logger.Logger
}

// MutationFishAdminAction handles fish creation and update actions for admin app.
type MutationToolAdminAction struct {
	uc  usecase.ToolAdminUseCase
	log logger.Logger
	val *validator.Validator
}

func NewToolAction(uc usecase.ToolUseCase, log logger.Logger) ToolAction {
	return ToolAction{
		uc:  uc,
		log: log,
	}
}

func NewToolAdminAction(uc usecase.ToolAdminUseCase, log logger.Logger) ToolAdminAction {
	return ToolAdminAction{
		uc:  uc,
		log: log,
	}
}

func NewMutationToolAdminAction(uc usecase.ToolAdminUseCase, log logger.Logger, val *validator.Validator) MutationToolAdminAction {
	return MutationToolAdminAction{
		uc:  uc,
		log: log,
		val: val,
	}
}

func (t ToolAction) FindOne(c *gin.Context) {
	const logKey = "find_one_Tool"
	fmt.Println("")
	output, err := t.uc.FindOneExecute(c.Request.Context(), utils.StrToInt(c.Param("id")))
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the Tool")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning Tool")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}


func (t ToolAction) FindAll(c *gin.Context) {
	const logKey = "find_all_Tools"

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
func (t ToolAdminAction) FindOneByAdmin(c *gin.Context) {
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
func (t ToolAdminAction) FindByAdmin(c *gin.Context) {
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
func (t MutationToolAdminAction) CreateByAdmin(c *gin.Context) {
	const logKey = "create_fish"
	fmt.Println("")

	var requestParam domain.Tool
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
		).Log("error when returning")
		response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when retur")
	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}

// UpdateByAdmin creates a new fish entry for admin app.
//
// The request payload is validated before Updating the fish entry.
func (t MutationToolAdminAction) UpdateByAdmin(c *gin.Context) {
	const logKey = "create_fish"
	fmt.Println("")

	id := utils.StrToInt(c.Param("id"))
	var requestParam domain.Tool
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
		).Log("error when returning")
		response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when return")
	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}


func (t ToolAdminAction) DeleteByAdmin(c *gin.Context) {
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
        ).Log("error when deleting")
        response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
        return
    }

    logging.NewInfo(t.log, logKey, http.StatusOK).Log("successfully deleted")
    response.NewSuccess(map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Tool deleted successfully",
	}, http.StatusOK).Send(c.Writer)
}
