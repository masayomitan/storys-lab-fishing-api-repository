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
		ImageUrl:	image.ImageUrl,
		Type:		image.Type,
		S3Url:     	image.S3Url,
		Sort: 		image.Sort,
		IsMain: 	image.IsMain,
	}
}

func (a ImagePresenter) PresentAll(images []domain.Image) []domain.Image {
	var output = make([]domain.Image, 0)
	for _, images := range images {
		output = append(output, domain.Image{
			ID:			images.ID,
			ImageUrl:	images.ImageUrl,
			S3Url:     	images.S3Url,
			Type:     	images.Type,
			Sort: 		images.Sort,
			IsMain:		images.IsMain,
		})
	}
	return output
}
