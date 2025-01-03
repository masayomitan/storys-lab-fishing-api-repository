package usecase

import (
	"context"
	"time"
	"io"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
	"storys-lab-fishing-api/app/service"
)

type ImageAdminUseCase interface {
	// FindOneExecuteByAdmin(ctx context.Context, id int) (domain.Image, error)
	// FindAllExecuteByAdmin(ctx context.Context) ([]domain.Image, error)
	UpdateExecuteByAdmin(ctx context.Context, file io.Reader, requestParam domain.Image)  (domain.Image, error)
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


// func (t ImageAdminInteractor) FindOneExecuteByAdmin(ctx context.Context, id int) (domain.Image, error) {
// 	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
// 	defer cancel()

// 	image, err := t.repo.FindOneByAdmin(ctx, id)
// 	if err != nil {
// 		return t.presenter.PresentOne(domain.Image{}), err
// 	}

// 	return t.presenter.PresentOne(image), nil
// }

// func (t ImageAdminInteractor) FindAllExecuteByAdmin(ctx context.Context) ([]domain.Image, error) {
// 	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
// 	defer cancel()

// 	images, err := t.repo.FindAllByAdmin(ctx)
// 	if err != nil {
// 		return t.presenter.PresentAll([]domain.Image{}), err
// 	}

// 	return t.presenter.PresentAll(images), nil
// }

func (t ImageAdminInteractor) UpdateExecuteByAdmin(ctx context.Context, file io.Reader, requestParam domain.Image) (domain.Image, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	// S3に画像をアップロード
	s3URL, err := t.s3Uploader.UploadFile(ctx, file, requestParam.ImageUrl)
	if err != nil {
		return domain.Image{}, errors.Wrap(err, "failed to upload image to S3")
	}

	image, err := t.repo.UpdateByAdmin(ctx, requestParam)

	// アップロードしたURLをデータにセット
	image.ImageUrl = s3URL
	if err != nil {
		return t.presenter.PresentOne(domain.Image{}), err
	}

	return t.presenter.PresentOne(image), nil
}
