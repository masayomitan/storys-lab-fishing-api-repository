package usecase

import (
	"context"
	"time"
	"io"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
	"storys-lab-fishing-api/app/service"
)

type ImageUploadPayload struct {
	File io.Reader
	Name string
}

type ImageAdminUseCase interface {
	// FindOneExecuteByAdmin(ctx context.Context, id int) (domain.Image, error)
	FindAllExecuteByAdmin(ctx context.Context, typeID string) ([]domain.Image, error)
	UploadExecuteByAdmin(ctx context.Context, requestPayload struct {
		Images    []ImageUploadPayload
		ImageType int
	})  ([]domain.Image, error)
}

type ImagePresenter interface {
	PresentOne(domain.Image) domain.Image
	PresentAll([]domain.Image) []domain.Image
}

type ImageAdminInteractor struct {
	repo       repository.ImageAdminRepository
	presenter  ImagePresenter
	s3Uploader service.S3Uploader
	ctxTimeout time.Duration
}


func NewImageAdminInteractor(
	repo repository.ImageAdminRepository,
	presenter ImagePresenter,
	timeout time.Duration,
) ImageAdminUseCase {
	return &ImageAdminInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}

func NewUploadImageAdminInteractor(
	repo repository.ImageAdminRepository,
	presenter ImagePresenter,
	s3Uploader service.S3Uploader,
	timeout time.Duration,
) ImageAdminUseCase {
	return &ImageAdminInteractor{
		repo:       repo,
		presenter:  presenter,
		s3Uploader: s3Uploader,
		ctxTimeout: timeout,
	}
}
