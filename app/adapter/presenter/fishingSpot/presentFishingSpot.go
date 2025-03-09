package presenter

import (
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/fishingSpot"
)

type fishingSpotPresenter struct{}

func NewFishingSpotPresenter() usecase.FishingSpotPresenter {
	return fishingSpotPresenter{}
}

func (a fishingSpotPresenter) PresentOne(fishingSpot domain.FishingSpot) domain.FishingSpot {
	return domain.FishingSpot{
		ID:             			fishingSpot.ID,
		Name:           			fishingSpot.Name,
		AreaID: 					fishingSpot.AreaID,
		Description:    			fishingSpot.Description,
		RecommendedFishingMethods: 	fishingSpot.RecommendedFishingMethods,
		CreatedAt:					fishingSpot.CreatedAt,
		UpdatedAt:					fishingSpot.UpdatedAt,
		Area: 						fishingSpot.Area,
		Fishes:						fishingSpot.Fishes,
		Tags: 						fishingSpot.Tags,
		Images: 					fishingSpot.Images,
	}
}

func (a fishingSpotPresenter) PresentAll(fishingSpot []domain.FishingSpot) []domain.FishingSpot {
	var output = make([]domain.FishingSpot, 0)
	for _, fishingSpot := range fishingSpot {
		output = append(output, domain.FishingSpot{
			ID:             			fishingSpot.ID,
			Name:           			fishingSpot.Name,
			AreaID: 					fishingSpot.AreaID,
			Description:    			fishingSpot.Description,
			RecommendedFishingMethods:	fishingSpot.RecommendedFishingMethods,
			CreatedAt:					fishingSpot.CreatedAt,
			UpdatedAt:					fishingSpot.UpdatedAt,
			Area: 						fishingSpot.Area,
			Fishes:						fishingSpot.Fishes,
			Tags: 						fishingSpot.Tags,
			Images: 					fishingSpot.Images,
		})
	}
	return output
}
