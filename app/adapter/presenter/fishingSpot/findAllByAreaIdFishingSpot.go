package presenter

import (
	// "time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/fishingSpot"
)

type findAllFishingSpotPresenter struct{}

func NewFindAllFishingSpotPresenter() usecase.FindAllFishingSpotPresenter {
	return findAllFishingSpotPresenter{}
}

func (a findAllFishingSpotPresenter) Output(fishingSpot []domain.FishingSpot) []domain.FishingSpot {
	var output = make([]domain.FishingSpot, 0)
	for _, fishingSpot := range fishingSpot {
		output = append(output, domain.FishingSpot{
			ID:             fishingSpot.ID,
			Name:           fishingSpot.Name,
			AreaId: 		fishingSpot.AreaId,
			Description:    fishingSpot.Description,
		})
	}
	return output
}

