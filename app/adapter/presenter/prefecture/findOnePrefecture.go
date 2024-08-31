package presenter

import (
	// "time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/prefecture"
)

type findOnePrefPresenter struct{}

func NewFindOnePrefPresenter() usecase.FindOnePrefPresenter {
	return findOnePrefPresenter{}
}

func (a findOnePrefPresenter) Output(pref domain.Pref) domain.Pref {
	return domain.Pref{
		ID:             pref.ID,
		Name:           pref.Name,
		NameKana:       pref.NameKana,
		Areas:			pref.Areas,
	}
}
