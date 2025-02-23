package action

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/prefecture"
	"storys-lab-fishing-api/app/utils"
)

type PrefectureAction struct {
	uc  usecase.PrefectureUseCase
	log logger.Logger
}

type PrefectureAdminAction struct {
	uc  usecase.PrefectureAdminUseCase
	log logger.Logger
}

func NewPrefectureAction(uc usecase.PrefectureUseCase, log logger.Logger) PrefectureAction {
	return PrefectureAction{
		uc:  uc,
		log: log,
	}
}

func NewPrefectureAdminAction(uc usecase.PrefectureAdminUseCase, log logger.Logger) PrefectureAdminAction {
	return PrefectureAdminAction{
		uc:  uc,
		log: log,
	}
}


func (t PrefectureAction) FindOne(c *gin.Context) {
	const logKey = "find_one_pref"
	fmt.Println("")
	output, err := t.uc.FindOneExecute(c.Request.Context(), utils.StrToInt(c.Param("id")))
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the pref")

		response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning area")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}


func (t PrefectureAdminAction) FindByAdmin(c *gin.Context) {
	const logKey = "find_one_pref"
	fmt.Println("")
	output, err := t.uc.FindExecuteByAdmin(c.Request.Context())
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the pref")

		response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning area")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}
