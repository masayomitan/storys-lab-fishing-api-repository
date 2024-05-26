package presenter

import (
	// "time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/fish"
)

type findOneFishPresenter struct{}

func NewFindOneFishPresenter() usecase.FindOneFishPresenter {
	return findOneFishPresenter{}
}

func (a findOneFishPresenter) Output(fish domain.Fish) domain.FishStruct {
	return domain.FishStruct{
		ID:             fish.ID().String(),
		Name:           fish.Name(),
		FamilyName:     fish.FamilyName(),
		ScientificName: fish.ScientificName(),
		FishCategoryId: fish.FishCategoryId(),
		Description:    fish.Description(),
	}
}
