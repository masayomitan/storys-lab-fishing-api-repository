package action

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/toolCategory"
)

// type FindOneToolCategoryAction struct {
// 	uc  usecase.FindOneToolCategoryUseCase
// 	log logger.Logger
// }

type FindAllToolCategoryAction struct {
	uc  usecase.FindAllToolCategoryUseCase
	log logger.Logger
}

// func NewFindOneToolCategoryAction(uc usecase.FindOneToolCategoryUseCase, log logger.Logger) FindOneToolAction {
// 	return FindOneToolCategoryAction{
// 		uc:  uc,
// 		log: log,
// 	}
// }

func NewFindAllToolCategoryAction(uc usecase.FindAllToolCategoryUseCase, log logger.Logger) FindAllToolCategoryAction {
	return FindAllToolCategoryAction{
		uc:  uc,
		log: log,
	}
}

// func (t FindOneToolCategoryAction) FindOne(c *gin.Context) {
// 	const logKey = "find_one_Tool"
// 	fmt.Println("")
// 	output, err := t.uc.Execute(c.Request.Context(), c.Param("id"))
// 	if err != nil {
// 		logging.NewError(
// 			t.log,
// 			err,
// 			logKey,
// 			http.StatusInternalServerError,
// 		).Log("error when returning the Tool")

// 		return
// 	}
// 	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning Tool")

// 	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
// }


func (t FindAllToolCategoryAction) FindAll(c *gin.Context) {
	const logKey = "find_all_tool_categories"
	fmt.Println("")
	output, err := t.uc.Execute(c.Request.Context())
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the Tool list")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning Tool list")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}
