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
			PrefectureId: a.PrefectureId,
		})
	}
	return result
}
