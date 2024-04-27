package presenter

import (
	// "time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase"
)

type findAllFishPresenter struct{}

func NewFindAllFishPresenter() usecase.FindAllFishPresenter {
	return findAllFishPresenter{}
}

func (a findAllFishPresenter) Output(fishes []domain.Fish) []usecase.FindAllFishOutput {
	var output = make([]usecase.FindAllFishOutput, 0)
	for _, fish := range fishes {
		output = append(output, usecase.FindAllFishOutput{
			ID:        fish.ID().String(),
			Name:      fish.Name(),
		})
	}
	return output
}
