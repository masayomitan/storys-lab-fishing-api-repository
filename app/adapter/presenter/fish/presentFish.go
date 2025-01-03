package presenter

import (
	// "time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/fish"
)

type findFishPresenter struct{}

func NewFishPresenter() usecase.FishPresenter {
	return findFishPresenter{}
}

func (a findFishPresenter) PresentOne(fish domain.Fish) domain.Fish {
	return domain.Fish{
		ID:             		fish.ID,
		Name:           		fish.Name,
		FamilyName:     		fish.FamilyName,
		ScientificName: 		fish.ScientificName,
		FishCategoryID: 		fish.FishCategoryID,
		Description:    		fish.Description,
		Length:    				fish.Length,
		Weight:    				fish.Weight,
		Habitat:    			fish.Habitat,
		DepthRange:    			fish.DepthRange,
		WaterTemperatureRange: 	fish.WaterTemperatureRange,
		ConservationStatus: 	fish.ConservationStatus,
		FishCategory: 			fish.FishCategory,
		FishingMethods: 		fish.FishingMethods,
		Dishes: 				fish.Dishes,
		FishImages: 			fish.FishImages,
	}
}

func (a findFishPresenter) PresentAll(fishes []domain.Fish) []domain.Fish {
	var output = make([]domain.Fish, 0)
	for _, fish := range fishes {
		output = append(output, domain.Fish{
			ID:				fish.ID,
			Name: 			fish.Name,
			FamilyName: 	fish.FamilyName,
			ScientificName: fish.ScientificName,
			FishCategoryID: fish.FishCategoryID,
			Description: 	fish.Description,
			FishImages: 	fish.FishImages,
		})
	}
	return output
}
