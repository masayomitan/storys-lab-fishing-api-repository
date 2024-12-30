package action

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/adapter/validator"
	"storys-lab-fishing-api/app/usecase/fishCategory"
	"storys-lab-fishing-api/app/domain"
	// "storys-lab-fishing-api/app/utils"
)

type FishCategoryAdminAction struct {
	uc  usecase.FishCategoryAdminUseCase
	log logger.Logger
}

type FishCategoryCreateAdminAction struct {
	uc  usecase.FishCategoryAdminUseCase
	log logger.Logger
	val *validator.Validator
}

func NewFishCategoryAdminAction(uc usecase.FishCategoryAdminUseCase, log logger.Logger) FishCategoryAdminAction {
	return FishCategoryAdminAction{
		uc:  uc,
		log: log,
	}
}

func NewFishCategoryCreateAdminAction(uc usecase.FishCategoryAdminUseCase, log logger.Logger, val *validator.Validator) FishCategoryCreateAdminAction {
	return FishCategoryCreateAdminAction{
		uc:  uc,
		log: log,
		val: val,
	}
}

// func (t FindOneToolAction) FindOneAdmin(c *gin.Context) {
// 	const logKey = "find_one_Tool"
// 	fmt.Println("")
// 	output, err := t.uc.Execute(c.Request.Context(), utils.StrToInt(c.Param("id")))
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

func (t FishCategoryAdminAction) FindAllByAdmin(c *gin.Context) {
	const logKey = "find_one_fish_category"
	fmt.Println("")
	output, err := t.uc.FindAllByAdmin(c.Request.Context())
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the fish-category")

		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning fish-category")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}

func (t FishCategoryCreateAdminAction) CreateByAdmin(c *gin.Context) {
	const logKey = "create_fish_category"
	fmt.Println("")

	var requestParam domain.FishCategory
	if err := c.ShouldBindJSON(&requestParam); err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusUnprocessableEntity,
		).Log("invalid request payload")
		response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
		return
	}
	fmt.Println(requestParam)

	if err := t.val.ValidateStruct(requestParam); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	output, err := t.uc.CreateByAdmin(c.Request.Context(), requestParam)
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the fish-category")
		response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning fish-category")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}
