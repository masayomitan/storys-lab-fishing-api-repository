package action

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/prefecture"
)

type FindOnePrefAction struct {
	uc  usecase.FindOnePrefUseCase
	log logger.Logger
}

func NewFindOnePrefAction(uc usecase.FindOnePrefUseCase, log logger.Logger) FindOnePrefAction {
	return FindOnePrefAction{
		uc:  uc,
		log: log,
	}
}

func (t FindOnePrefAction) FindOne(c *gin.Context) {
	const logKey = "find_one_pref"
	fmt.Println("")
	output, err := t.uc.Execute(c.Request.Context(), c.Param("id"))
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the pref")

		// response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning area")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}
