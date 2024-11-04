package action

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/tool"
	"storys-lab-fishing-api/app/utils"
)

type FindOneToolAction struct {
	uc  usecase.FindOneToolUseCase
	log logger.Logger
}

// type FindAllToolAction struct {
// 	uc  usecase.FindAllToolUseCase
// 	log logger.Logger
// }

func NewFindOneToolAction(uc usecase.FindOneToolUseCase, log logger.Logger) FindOneToolAction {
	return FindOneToolAction{
		uc:  uc,
		log: log,
	}
}

// func NewFindAllToolAction(uc usecase.FindAllToolUseCase, log logger.Logger) FindAllToolAction {
// 	return FindAllToolAction{
// 		uc:  uc,
// 		log: log,
// 	}
// }

func (t FindOneToolAction) FindOne(c *gin.Context) {
	const logKey = "find_one_Tool"
	fmt.Println("")
	output, err := t.uc.Execute(c.Request.Context(), utils.StrToInt(c.Param("area_id")))
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the Tool")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning Tool")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}


// func (t FindAllToolAction) FindAll(c *gin.Context) {
// 	const logKey = "find_all_Tool"

// 	output, err := t.uc.Execute(c.Request.Context())
// 	if err != nil {
// 		logging.NewError(
// 			t.log,
// 			err,
// 			logKey,
// 			http.StatusInternalServerError,
// 		).Log("error when returning the Tool list")

// 		return
// 	}
// 	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning Tool list")

// 	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
// }
