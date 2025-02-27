package model

import (
	"storys-lab-fishing-api/app/domain"
)

func convertAreas(areas []domain.Area) []domain.Area {
	var result []domain.Area
	for _, a := range areas {
		result = append(result, domain.Area{
			ID:   a.ID,
			Name: a.Name,
			Description: a.Description,
			PrefectureID: a.PrefectureID,
			Images: convertImages(a.Images),
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
