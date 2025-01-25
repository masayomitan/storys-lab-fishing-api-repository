package usecase

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

// func (t ImageAdminInteractor) FindOneExecuteByAdmin(ctx context.Context, id int) (domain.Image, error) {
// 	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
// 	defer cancel()

// 	image, err := t.repo.FindOneByAdmin(ctx, id)
// 	if err != nil {
// 		return t.presenter.PresentOne(domain.Image{}), err
// 	}

// 	return t.presenter.PresentOne(image), nil
// }

func (t ImageAdminInteractor) FindAllExecuteByAdmin(ctx context.Context, imageType string) ([]domain.Image, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	query := map[string]interface{}{}
	if imageType != "" {
		query["image_type"] = imageType
	}

	images, err := t.repo.FindAllByAdmin(ctx, query)
	if err != nil {
		return t.presenter.PresentAll([]domain.Image{}), err
	}

	return t.presenter.PresentAll(images), nil
}
