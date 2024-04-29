package action

import (
	"net/http"

	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase"
)

type FindAllFishAction struct {
	uc  usecase.FindAllFishUseCase
	log logger.Logger
}

func NewFindAllFishAction(uc usecase.FindAllFishUseCase, log logger.Logger) FindAllFishAction {
	return FindAllFishAction{
		uc:  uc,
		log: log,
	}
}

func (t FindAllFishAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "find_all_fish"

	output, err := t.uc.Execute(r.Context())
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

	response.NewSuccess(output, http.StatusOK).Send(w)
}