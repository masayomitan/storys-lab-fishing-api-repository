package action

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/area"
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
