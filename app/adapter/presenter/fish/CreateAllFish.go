package presenter

import (
	// "time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/fish"
)

type createFishPresenter struct{}

func NewCreateFishPresenter() usecase.CreateFishPresenter {
	return createFishPresenter{}
}

func (a createFishPresenter) Output(fish domain.Fish) domain.Fish {
	return domain.Fish{
		ID: fish.ID,
		Name: fish.Name,
		FamilyName: fish.FamilyName,
		ScientificName: fish.ScientificName,
		FishCategoryId: fish.FishCategoryId,
		Description: fish.Description,
	}
}
