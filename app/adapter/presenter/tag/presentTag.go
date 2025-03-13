package presenter

import (
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/tag"
)

type tagPresenter struct{}

func NewTagPresenter() usecase.TagPresenter {
	return tagPresenter{}
}

func (a tagPresenter) Present(fishingSpot domain.Tag) domain.Tag {
	return domain.Tag{
		ID:             			fishingSpot.ID,
		Name:           			fishingSpot.Name,
		CreatedAt:					fishingSpot.CreatedAt,
		UpdatedAt:					fishingSpot.UpdatedAt,
	}
}

func (a tagPresenter) Presents(fishingSpot []domain.Tag) []domain.Tag {
	var output = make([]domain.Tag, 0)
	for _, fishingSpot := range fishingSpot {
		output = append(output, domain.Tag{
			ID:             			fishingSpot.ID,
			Name:           			fishingSpot.Name,
			CreatedAt:					fishingSpot.CreatedAt,
			UpdatedAt:					fishingSpot.UpdatedAt,
		})
	}
	return output
}
