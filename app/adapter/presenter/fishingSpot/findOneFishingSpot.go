package presenter

import (
	// "time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/fishingSpot"
)

type findOneFishingSpotPresenter struct{}

func NewFindOneFishingSpotPresenter() usecase.FindOneFishingSpotPresenter {
	return findOneFishingSpotPresenter{}
}

func (a findOneFishingSpotPresenter) Output(fishingSpot domain.FishingSpot) domain.FishingSpot {
	return domain.FishingSpot{
		ID:             fishingSpot.ID,
		Name:           fishingSpot.Name,
		ImageUrl:       fishingSpot.ImageUrl,
		AreaId: 		fishingSpot.AreaId,
		Description:    fishingSpot.Description,

		Area: fishingSpot.Area,
		Fishes: fishingSpot.Fishes,
		Tags: fishingSpot.Tags,
	}
}
