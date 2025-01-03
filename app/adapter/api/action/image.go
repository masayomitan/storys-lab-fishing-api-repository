package action

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/image"
	"storys-lab-fishing-api/app/adapter/validator"
	"storys-lab-fishing-api/app/domain"
	// "storys-lab-fishing-api/app/utils"
)

// MutationImageAdminAction handles Image creation and update actions for admin app.
type MutationImageAdminAction struct {
	uc  usecase.ImageAdminUseCase
	log logger.Logger
	val *validator.Validator
}

// NewImageCreateAdminAction creates a new instance of MutationImageAdminAction.
//
// Args:
// - uc: Use case interface for Image admin actions.
// - log: Logger for recording events.
// - val: Validator for input validation.
func NewUpdateImageAdminAction(uc usecase.ImageAdminUseCase, log logger.Logger, val *validator.Validator) MutationImageAdminAction {
	return MutationImageAdminAction{
		uc:  uc,
		log: log,
		val: val,
	}
}

// UpdateByAdmin creates a new Image entry for admin app.
//
// The request payload is validated before creating the Image entry.
func (t MutationImageAdminAction) UpdateByAdmin(c *gin.Context) {
	const logKey = "update_image"
	fmt.Println("")

	var requestParam domain.Image
	fmt.Println(requestParam)

	fileHeader, err := c.FormFile("file")
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("failed to get file from request")
		response.NewError(err, http.StatusBadRequest).Send(c.Writer)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("failed to open file")
		response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
		return
	}
	defer file.Close()

	output, err := t.uc.UpdateExecuteByAdmin(c.Request.Context(), file, requestParam)
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning the image")
		response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
		return
	}
	logging.NewInfo(t.log, logKey, http.StatusOK).Log("success when returning image")

	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}
