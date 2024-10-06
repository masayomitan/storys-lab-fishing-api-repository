package presenter

import (
	// "time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/area"
)

type findAllAreaPresenter struct{}

func NewFindAllAreaPresenter() usecase.FindAllAreaPresenter {
	return findAllAreaPresenter{}
}

func (a findAllAreaPresenter) Output(areas []domain.Area) []domain.Area {
	var output = make([]domain.Area, 0)
	for _, area := range areas {
		output = append(output, domain.Area{
			ID:				area.ID,
			Name:			area.Name,
			Description: 	area.Description,
			PrefectureId: 	area.PrefectureId,
			ImageUrl:	 	area.ImageUrl,
		})
	}
	return output
}
