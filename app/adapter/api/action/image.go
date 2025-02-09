package action

import (
	"net/http"
	// "fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/adapter/api/logging"
	"storys-lab-fishing-api/app/adapter/api/response"
	"storys-lab-fishing-api/app/adapter/logger"
	"storys-lab-fishing-api/app/usecase/image"
	"storys-lab-fishing-api/app/adapter/validator"
	// "storys-lab-fishing-api/app/domain"
	// "storys-lab-fishing-api/app/utils"
)

type ImageAdminAction struct {
	uc  usecase.ImageAdminUseCase
	log logger.Logger
}

// MutationImageAdminAction handles Image creation and update actions for admin app.
type MutationImageAdminAction struct {
	uc  usecase.ImageAdminUseCase
	log logger.Logger
	val *validator.Validator
}

// NewImageAdminAction find new instance of ImageAdminAction.
//
// Args:
// - uc: Use case interface for Image admin actions.
// - log: Logger for recording events.
func NewImageAdminAction(uc usecase.ImageAdminUseCase, log logger.Logger) ImageAdminAction {
	return ImageAdminAction{
		uc:  uc,
		log: log,
	}
}

// NewImageCreateAdminAction creates a new instance of MutationImageAdminAction.
//
// Args:
// - uc: Use case interface for Image admin actions.
// - log: Logger for recording events.
// - val: Validator for input validation.
func NewUploadImageAdminAction(uc usecase.ImageAdminUseCase, log logger.Logger, val *validator.Validator) MutationImageAdminAction {
	return MutationImageAdminAction{
		uc:  uc,
		log: log,
		val: val,
	}
}

func (t ImageAdminAction) FindAllByAdmin(c *gin.Context) {
	const logKey = "find_all_fishes"

	imageType := c.Query("type")

	output, err := t.uc.FindAllExecuteByAdmin(c.Request.Context(), imageType)
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

// UpdateByAdmin creates a new Image entry for admin app.
//
// The request payload is validated before creating the Image entry.
func (t MutationImageAdminAction) UploadByAdmin(c *gin.Context) {
	const logKey = "upload_image"

	// Parse the multipart form
	form, err := c.MultipartForm()
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("failed to parse multipart form")
		response.NewError(err, http.StatusBadRequest).Send(c.Writer)
		return
	}

	// Retrieve files from the form
	files := form.File["files"]
	names := form.Value["names"]

	if len(files) == 0 {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("no files found in the request")
		response.NewError(err, http.StatusBadRequest).Send(c.Writer)
		return
	}

	// Retrieve image_type from the form
	imageTypeStr := form.Value["image_type"]
	imageType, _ := strconv.Atoi(imageTypeStr[0])

	var imagePayloads []usecase.ImageUploadPayload

	for i, fileHeader := range files {
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

		name := ""
		if i < len(names) {
			name = names[i]
		}

		imagePayloads = append(imagePayloads, usecase.ImageUploadPayload{
			File: file,
			Name: name,
		})
	}

	requestPayload := struct {
		Images    []usecase.ImageUploadPayload
		ImageType int
	}{
		Images:    imagePayloads,
		ImageType: imageType,
	}

	// Call the use case
	output, err := t.uc.UploadExecuteByAdmin(c.Request.Context(), requestPayload)
	if err != nil {
		logging.NewError(
			t.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when processing the image upload")
		response.NewError(err, http.StatusInternalServerError).Send(c.Writer)
		return
	}

	logging.NewInfo(t.log, logKey, http.StatusOK).Log("successfully uploaded images")
	response.NewSuccess(output, http.StatusOK).Send(c.Writer)
}

