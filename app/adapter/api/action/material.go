package action

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/material"
	// "storys-lab-fishing-api/app/domain"
	// "storys-lab-fishing-api/app/adapter/validator"
	// "storys-lab-fishing-api/app/utils"
)

type MaterialAction struct {
	uc  usecase.MaterialUseCase
	log logger.Logger
}

type MaterialAdminAction struct {
	uc  usecase.MaterialAdminUseCase
	log logger.Logger
}

// func NewMaterialAction(uc usecase.MaterialUseCase, log logger.Logger) MaterialAction {
// 	return MaterialAction{
// 		uc:  uc,
// 		log: log,
// 	}
// }

func NewMaterialAdminAction(uc usecase.MaterialAdminUseCase, log logger.Logger) MaterialAdminAction {
	return MaterialAdminAction{
		uc:  uc,
		log: log,
	}
}

func (t MaterialAdminAction) FindByAdmin(c *gin.Context) {
	const logKey = "find_all_Materials"

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
