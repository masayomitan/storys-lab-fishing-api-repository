package usecase

import (
	"context"
	"time"
	"fmt"
	"os"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
)

func (t ImageAdminInteractor) UploadExecuteByAdmin(ctx context.Context, requestPayload struct {
    Images    []ImageUploadPayload
    ImageType int
}) ([]domain.Image, error) {
    ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
    defer cancel()

    if len(requestPayload.Images) == 0 {
        return nil, errors.New("no images provided")
    }

    var path string
    switch requestPayload.ImageType {
    case 1:
        path = "images/fishes"
    case 2:
        path = "images/areas"
    default:
        path = "images/others"
    }
    uploadedImages := make([]domain.Image, 0, len(requestPayload.Images))

    for _, img := range requestPayload.Images {
        // Generate a unique file name
        fileName := fmt.Sprintf("%s/%d-%s", path, time.Now().UnixNano(), "uploaded_image.jpg")

        // Upload image to S3
        s3URL, err := t.s3Uploader.UploadFile(ctx, img.File, fileName)
        if err != nil {
            return nil, errors.Wrap(err, "failed to upload image to S3")
        }
        imageURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", os.Getenv("AWS_S3_BUCKET"), fileName)

        // Create image object with details
        image := domain.Image{
            Name:      img.Name,
            ImageType: requestPayload.ImageType,
            S3Url:     s3URL,
            ImageUrl:  imageURL,
        }

        // Save the image details in the database
        savedImage, err := t.repo.UploadByAdmin(ctx, image)
        if err != nil {
            return nil, errors.Wrap(err, "failed to save image to the database")
        }
        uploadedImages = append(uploadedImages, t.presenter.PresentOne(savedImage))
    }
    return uploadedImages, nil
}



