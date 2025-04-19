package action

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/dish"
	"storys-lab-fishing-api/app/adapter/validator"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/utils"
)

type DishAction struct {
	uc  usecase.DishUseCase
	log logger.Logger
}
type DishAdminAction struct {
	uc  usecase.DishAdminUseCase
	log logger.Logger
}

// MutationFishAdminAction handles fish creation and update actions for admin app.
type MutationDishAdminAction struct {
	uc  usecase.DishAdminUseCase
	log logger.Logger
	val *validator.Validator
}

func NewDishAction(uc usecase.DishUseCase, log logger.Logger) DishAction {
	return DishAction{
		uc:  uc,
		log: log,
	}
}

func NewDishAdminAction(uc usecase.DishAdminUseCase, log logger.Logger) DishAdminAction {
	return DishAdminAction{
		uc:  uc,
		log: log,
	}
}

func NewMutationDishAdminAction(uc usecase.DishAdminUseCase, log logger.Logger, val *validator.Validator) MutationDishAdminAction {
	return MutationDishAdminAction{
		uc:  uc,
		log: log,
		val: val,
	}
}

func (t DishAction) FindOne(c *gin.Context) {
	const logKey = "find_one_Dish"
	fmt.Println("")
	output, err := t.uc.FindOneExecute(c.Request.Context(), utils.StrToInt(c.Param("id")))
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the Dish")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning Dish")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}


func (t DishAction) FindAll(c *gin.Context) {
	const logKey = "find_all_Dishs"

	output, err := t.uc.FindExecute(c.Request.Context())
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the Dish list")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning Dish list")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}

// FindOneByAdmin retrieves a list of all fishes for admin app.
func (t DishAdminAction) FindOneByAdmin(c *gin.Context) {
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
func (t DishAdminAction) FindByAdmin(c *gin.Context) {
	const logKey = "find_all_fishes"

	output, err := t.uc.FindExecuteByAdmin(c.Request.Context())
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
func (t MutationDishAdminAction) CreateByAdmin(c *gin.Context) {
	const logKey = "create_fish"
	fmt.Println("")

	var requestParam domain.Dish
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
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when return")
	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}

// UpdateByAdmin creates a new fish entry for admin app.
//
// The request payload is validated before Updating the fish entry.
func (t MutationDishAdminAction) UpdateByAdmin(c *gin.Context) {
	const logKey = "create_fish"
	fmt.Println("")

	id := utils.StrToInt(c.Param("id"))
	var requestParam domain.Dish
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


func (t DishAdminAction) DeleteByAdmin(c *gin.Context) {
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
		"message": "Dish deleted successfully",
	}, http.StatusOK).Send(c.Writer)
}
