package action

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/fishingSpot"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/adapter/validator"
	"storys-lab-fishing-api/app/utils"
)

type FishingSpotAction struct {
	uc  usecase.FishingSpotUseCase
	log logger.Logger
}

type FishingSpotAdminAction struct {
	uc  usecase.FishingSpotAdminUseCase
	log logger.Logger
}


type MutationFishingSpotAdminAction struct {
	uc  usecase.FishingSpotAdminUseCase
	log logger.Logger
	val *validator.Validator
}


func NewFishingSpotAction(uc usecase.FishingSpotUseCase, log logger.Logger) FishingSpotAction {
	return FishingSpotAction{
		uc:  uc,
		log: log,
	}
}

func NewFishingSpotAdminAction(uc usecase.FishingSpotAdminUseCase, log logger.Logger) FishingSpotAdminAction {
	return FishingSpotAdminAction{
		uc:  uc,
		log: log,
	}
}

func NewMutationFishingSpotAdminAction(uc usecase.FishingSpotAdminUseCase, log logger.Logger, val *validator.Validator) MutationFishingSpotAdminAction {
	return MutationFishingSpotAdminAction{
		uc:  uc,
		log: log,
		val: val,
	}
}

func (t FishingSpotAction) FindOne(c *gin.Context) {
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


func (t FishingSpotAction) FindByAreaId(c *gin.Context) {
	const logKey = "find_all_fishing_spot_by_area_id"
	output, err := t.uc.FindByAreaIdExecute(c.Request.Context(), utils.StrToInt(c.Param("area_id")))
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


func (t FishingSpotAdminAction) FindByAdmin(c *gin.Context) {
	const logKey = "find_all_fishingSpots"

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

func (t FishingSpotAdminAction) FindOneByAdmin(c *gin.Context) {
	const logKey = "find_all_fishingSpots"

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

func (t MutationFishingSpotAdminAction) CreateByAdmin(c *gin.Context) {
	const logKey = "create_fishingSpot"
	fmt.Println("")

	var requestParam domain.FishingSpot
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
		).Log("error when returning the area")
		response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning area")
	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}

func (t MutationFishingSpotAdminAction) UpdateByAdmin(c *gin.Context) {
	const logKey = "create_area"
	fmt.Println("")

	id := utils.StrToInt(c.Param("id"))
	var requestParam domain.FishingSpot
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

func (t FishingSpotAdminAction) DeleteByAdmin(c *gin.Context) {
    const logKey = "delete_area"

	fmt.Println(c.Param("id"))
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
        ).Log("error when deleting the fishingSpot")
        response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
        return
    }

    logging.NewInfo(t.log, logKey, http.StatusOK).Log("successfully deleted fishingSpot")
    response.NewSuccess(map[string]interface{}{
		"status":  http.StatusOK,
		"message": "fishingSpot deleted successfully",
	}, http.StatusOK).Send(c.Writer)
}
