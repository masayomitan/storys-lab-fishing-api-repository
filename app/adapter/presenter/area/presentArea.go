package presenter

import (
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/area"
)

type findAreaPresenter struct{}

func NewAreaPresenter() usecase.AreaPresenter {
	return findAreaPresenter{}
}

func (a findAreaPresenter) PresentOne(area domain.Area) domain.Area {
	return domain.Area{
		ID: 			area.ID,
		Name: 			area.Name,
		Description: 	area.Description,
		PrefectureId: 	area.PrefectureId,
		FishingSpots: 	area.FishingSpots,
		CreatedAt:		area.CreatedAt,
		UpdatedAt:		area.UpdatedAt,
		Tides: 			area.Tides,
		Images:			area.Images,
	}
}

func (a findAreaPresenter) PresentAll(areas []domain.Area) []domain.Area {
	var output = make([]domain.Area, 0)
	for _, area := range areas {
		output = append(output, domain.Area{
			ID:				area.ID,
			Name:			area.Name,
			Description: 	area.Description,
			PrefectureId: 	area.PrefectureId,
			CreatedAt:		area.CreatedAt,
			UpdatedAt:		area.UpdatedAt,
			Images:			area.Images,
		})
	}
	return output
}
