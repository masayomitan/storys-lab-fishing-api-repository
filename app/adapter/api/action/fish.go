package action

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/fish"
	"storys-lab-fishing-api/app/adapter/validator"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/utils"
)

// FishAction handles fish-related actions for end-users.
type FishAction struct {
	uc  usecase.FishUseCase
	log logger.Logger
}

// FishAdminAction handles fish-related actions for admin app.
type FishAdminAction struct {
	uc  usecase.FishAdminUseCase
	log logger.Logger
}

// MutationFishAdminAction handles fish creation and update actions for admin app.
type MutationFishAdminAction struct {
	uc  usecase.FishAdminUseCase
	log logger.Logger
	val *validator.Validator
}

// NewFishAction creates a new instance of FishAction.
//
// Args:
// - uc: Use case interface for fish actions.
// - log: Logger for recording events.
func NewFishAction(uc usecase.FishUseCase, log logger.Logger) FishAction {
	return FishAction{
		uc:  uc,
		log: log,
	}
}

// NewFishAdminAction creates a new instance of FishAdminAction.
//
// Args:
// - uc: Use case interface for fish admin actions.
// - log: Logger for recording events.
func NewFishAdminAction(uc usecase.FishAdminUseCase, log logger.Logger) FishAdminAction {
	return FishAdminAction{
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
func NewMutationFishAdminAction(uc usecase.FishAdminUseCase, log logger.Logger, val *validator.Validator) MutationFishAdminAction {
	return MutationFishAdminAction{
		uc:  uc,
		log: log,
		val: val,
	}
}

// FindOne retrieves details of a specific fish.
//
// The fish ID is provided as a path parameter.
func (t FishAction) FindOne(c *gin.Context) {
	const logKey = "find_one_fish"
	fmt.Println("")
	output, err := t.uc.FindOneExecute(c.Request.Context(), utils.StrToInt(c.Param("id")))
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the fish")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning fish")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}

// FindAll retrieves a list of all fishes.
func (t FishAction) FindAll(c *gin.Context) {
	const logKey = "find_all_fishes"

	output, err := t.uc.FindAllExecute(c.Request.Context())
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

// FindOneByAdmin retrieves a list of all fishes for admin app.
func (t FishAdminAction) FindOneByAdmin(c *gin.Context) {
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
func (t FishAdminAction) FindAllByAdmin(c *gin.Context) {
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
func (t MutationFishAdminAction) CreateByAdmin(c *gin.Context) {
	const logKey = "create_fish"
	fmt.Println("")

	var requestParam domain.Fish
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
func (t MutationFishAdminAction) UpdateByAdmin(c *gin.Context) {
	const logKey = "create_fish"
	fmt.Println("")

	id := utils.StrToInt(c.Param("id"))
	var requestParam domain.Fish
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


func (t FishAdminAction) DeleteByAdmin(c *gin.Context) {
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

func (t FishAdminAction) UpdateFishDishes(c *gin.Context) {
	const logKey = "update_fish_dishes"

	id := utils.StrToInt(c.Param("id"))
	if id == 0 {
		err := fmt.Errorf("id is required")
		logging.NewError(t.log, err, logKey, http.StatusBadRequest).Log("missing id in request")
		response.NewError(err, http.StatusBadRequest).Send(c.Writer)
		return
	}

	var req domain.UpdateFishDishesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logging.NewError(t.log, err, logKey, http.StatusBadRequest).Log("invalid dish_ids format")
		response.NewError(err, http.StatusBadRequest).Send(c.Writer)
		return
	}

	result, err := t.uc.UpdateFishDishesExecute(c.Request.Context(), id, req.DishIDs)
	if err != nil {
		logging.NewError(t.log, err, logKey, http.StatusInternalServerError).Log("failed to update fish dishes relation")
		response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
		return
	}

	logging.NewInfo(t.log, logKey, http.StatusOK).Log("successfully updated fish dishes")
	response.NewSuccess(result, http.StatusOK).Send(c.Writer)
}


