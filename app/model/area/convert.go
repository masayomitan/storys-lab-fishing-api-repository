package model


import (
	// "gorm.io/gorm"
	"storys-lab-fishing-api/app/domain"
)

func convertFishingSpots(fishingSpot []domain.FishingSpot) []domain.FishingSpot {
	var result []domain.FishingSpot
	for _, f := range fishingSpot {
		result = append(result, domain.FishingSpot{
			ID: 			f.ID,
			Name: 			f.Name,
			ImageUrl:		f.ImageUrl,
			Description: 	f.Description,
			AreaId: 		f.AreaId,
			Tags: 			convertTags(f.Tags),
		})
	}
	return result
}

func convertTide(tide []domain.Tide) []domain.Tide {
	var result []domain.Tide
	for _, t := range tide {
		result = append(result, domain.Tide{
			ID: t.ID,
			AreaId:			 t.AreaId,
			PrefectureId:    t.PrefectureId,
		})
	}
	return result
}

func convertTags(tags []domain.Tag) []domain.Tag {
	var result []domain.Tag
	for _, t := range tags {
		result = append(result, domain.Tag{
			ID:        t.ID,
			Name:      t.Name,
		})
	}
	return result
}
