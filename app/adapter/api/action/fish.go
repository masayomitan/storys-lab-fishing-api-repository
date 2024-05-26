package action

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/fish"
)

type FindOneFishAction struct {
	uc  usecase.FindOneFishUseCase
	log logger.Logger
}

type FindAllFishAction struct {
	uc  usecase.FindAllFishUseCase
	log logger.Logger
}

func NewFindOneFishAction(uc usecase.FindOneFishUseCase, log logger.Logger) FindOneFishAction {
	return FindOneFishAction{
		uc:  uc,
		log: log,
	}
}

func NewFindAllFishAction(uc usecase.FindAllFishUseCase, log logger.Logger) FindAllFishAction {
	return FindAllFishAction{
		uc:  uc,
		log: log,
	}
}

func (t FindOneFishAction) FindOne(c *gin.Context) {
	const logKey = "find_one_fish"
	fmt.Println("")
	output, err := t.uc.Execute(c.Request.Context(), c.Param("id"))
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the fish")

		// response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning fish")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}


func (t FindAllFishAction) FindAll(c *gin.Context) {
	const logKey = "find_all_fish"

	output, err := t.uc.Execute(c.Request.Context())
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the fish list")

		// response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning fish list")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}
