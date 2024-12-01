package action

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/event"
	"storys-lab-fishing-api/app/utils"
)

type FindOneEventAction struct {
	uc  usecase.FindOneEventUseCase
	log logger.Logger
}

type FindAllEventAction struct {
	uc  usecase.FindAllEventUseCase
	log logger.Logger
}

func NewFindOneEventAction(uc usecase.FindOneEventUseCase, log logger.Logger) FindOneEventAction {
	return FindOneEventAction{
		uc:  uc,
		log: log,
	}
}

func NewFindAllEventAction(uc usecase.FindAllEventUseCase, log logger.Logger) FindAllEventAction {
	return FindAllEventAction{
		uc:  uc,
		log: log,
	}
}

func (t FindOneEventAction) FindOne(c *gin.Context) {
	const logKey = "find_one_Event"
	fmt.Println("")
	output, err := t.uc.Execute(c.Request.Context(), utils.StrToInt(c.Param("id")))
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the Event")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning Event")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}


func (t FindAllEventAction) FindAll(c *gin.Context) {
	const logKey = "find_all_Events"

	output, err := t.uc.Execute(c.Request.Context())
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the Event list")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning Event list")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}
