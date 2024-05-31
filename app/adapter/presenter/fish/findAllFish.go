package presenter

import (
	// "time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/fish"
)

type findAllFishPresenter struct{}

func NewFindAllFishPresenter() usecase.FindAllFishPresenter {
	return findAllFishPresenter{}
}

func (a findAllFishPresenter) Output(fishes []domain.Fish) []domain.FishStruct {
	var output = make([]domain.FishStruct, 0)
	for _, fish := range fishes {
		output = append(output, domain.FishStruct{
			ID:fish.ID(),
			Name: fish.Name(),
			FamilyName: fish.FamilyName(),
			ScientificName: fish.ScientificName(),
			FishCategoryId: fish.FishCategoryId(),
			Description: fish.Description(),
		})
	}
	return output
}
