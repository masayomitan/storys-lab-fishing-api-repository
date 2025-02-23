package action

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/area"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/adapter/validator"
	"storys-lab-fishing-api/app/utils"
)

type AreaAction struct {
	uc  usecase.AreaUseCase
	log logger.Logger
}

type AreaAdminAction struct {
	uc  usecase.AreaAdminUseCase
	log logger.Logger
}

type MutationAreaAdminAction struct {
	uc  usecase.AreaAdminUseCase
	log logger.Logger
	val *validator.Validator
}

func NewAreaAction(uc usecase.AreaUseCase, log logger.Logger) AreaAction {
	return AreaAction{
		uc:  uc,
		log: log,
	}
}

func NewAreaAdminAction(uc usecase.AreaAdminUseCase, log logger.Logger) AreaAdminAction {
	return AreaAdminAction{
		uc:  uc,
		log: log,
	}
}

func NewMutationAreaAdminAction(uc usecase.AreaAdminUseCase, log logger.Logger, val *validator.Validator) MutationAreaAdminAction {
	return MutationAreaAdminAction{
		uc:  uc,
		log: log,
		val: val,
	}
}

func (t AreaAction) Find(c *gin.Context) {
	const logKey = "find_all_area"

	output, err := t.uc.FindAllExecute(c.Request.Context())
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the area list")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning area list")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}

func (t AreaAction) FindOne(c *gin.Context) {
	const logKey = "find_one_area"
	fmt.Println("")
	output, err := t.uc.FindOneExecute(c.Request.Context(), utils.StrToInt(c.Param("id")))
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the area")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning area")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}


func (t AreaAdminAction) FindByAdmin(c *gin.Context) {
	const logKey = "find_all_areas"

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

func (t AreaAdminAction) FindOneByAdmin(c *gin.Context) {
	const logKey = "find_all_areas"

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

func (t MutationAreaAdminAction) CreateByAdmin(c *gin.Context) {
	const logKey = "create_area"
	fmt.Println("")

	var requestParam domain.Area
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
