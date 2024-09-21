package action

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/area"
)

type FindOneAreaAction struct {
	uc  usecase.FindOneAreaUseCase
	log logger.Logger
}

type FindAllAreaAction struct {
	uc  usecase.FindAllAreaUseCase
	log logger.Logger
}

func NewFindOneAreaAction(uc usecase.FindOneAreaUseCase, log logger.Logger) FindOneAreaAction {
	return FindOneAreaAction{
		uc:  uc,
		log: log,
	}
}

func NewFindAllAreaAction(uc usecase.FindAllAreaUseCase, log logger.Logger) FindAllAreaAction {
	return FindAllAreaAction{
		uc:  uc,
		log: log,
	}
}

func (t FindOneAreaAction) FindOne(c *gin.Context) {
	const logKey = "find_one_area"
	fmt.Println("")
	output, err := t.uc.Execute(c.Request.Context(), c.Param("id"))
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


func (t FindAllAreaAction) FindAll(c *gin.Context) {
	const logKey = "find_all_area"

	output, err := t.uc.Execute(c.Request.Context())
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
