package presenter

import (
	// "time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/area"
)

type findOneAreaPresenter struct{}

func NewFindOneAreaPresenter() usecase.FindOneAreaPresenter {
	return findOneAreaPresenter{}
}

func (a findOneAreaPresenter) Output(area domain.Area) domain.Area {
	return domain.Area{
		ID: area.ID,
		Name: area.Name,
		Description: area.Description,
		PrefectureId: area.PrefectureId,
		FishingSpots: area.FishingSpots,
	}
}
