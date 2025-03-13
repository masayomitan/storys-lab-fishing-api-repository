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
			Description: 	f.Description,
			AreaID: 		f.AreaID,
			Tags: 			convertTags(f.Tags),
			Images: 		convertImages(f.Images),
		})
	}
	return result
}

func convertTide(tide []domain.Tide) []domain.Tide {
	var result []domain.Tide
	for _, t := range tide {
		result = append(result, domain.Tide{
			ID: t.ID,
			AreaID:			 t.AreaID,
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

func convertImages(methods []domain.Image) []domain.Image {
	var result []domain.Image
	for _, i := range methods {
		result = append(result, domain.Image{
			ID: 		i.ID,
			Name: 		i.Name,
			ImageUrl: 	i.ImageUrl,
			S3Url: 		i.S3Url,
		})
	}
	return result
}
