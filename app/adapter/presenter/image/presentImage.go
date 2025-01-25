package presenter

import (
	// "time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/image"
)

type ImagePresenter struct{}

func NewImagePresenter() usecase.ImagePresenter {
	return ImagePresenter{}
}

func (a ImagePresenter) PresentOne(image domain.Image) domain.Image {
	return domain.Image{
		ID:			image.ID,
		Name:   	image.Name,
		ImageUrl:	image.ImageUrl,
		S3Url:     	image.S3Url,
		ImageType:	image.ImageType,
	}
}

func (a ImagePresenter) PresentAll(images []domain.Image) []domain.Image {
	var output = make([]domain.Image, 0)
	for _, images := range images {
		output = append(output, domain.Image{
			ID:			images.ID,
			Name:   	images.Name,
			ImageUrl:	images.ImageUrl,
			S3Url:     	images.S3Url,
			ImageType:	images.ImageType,
		})
	}
	return output
}
