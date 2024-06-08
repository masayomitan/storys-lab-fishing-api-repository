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

func (a findOneFishPresenter) Output(fish domain.Fish) domain.Fish {
	return domain.Fish{
		ID:             fish.ID,
		Name:           fish.Name,
		FamilyName:     fish.FamilyName,
		ScientificName: fish.ScientificName,
		FishCategoryId: fish.FishCategoryId,
		Description:    fish.Description,
		Length:    		fish.Length,
		Weight:    		fish.Weight,
		Habitat:    	fish.Habitat,
		DepthRange:    	fish.DepthRange,
		WaterTemperatureRange: fish.WaterTemperatureRange,
		ConservationStatus: fish.ConservationStatus,
		FishCategory: fish.FishCategory,
		FishingMethods: fish.FishingMethods,
		Dishes: fish.Dishes,
		FishImages: fish.FishImages,
	}
}
