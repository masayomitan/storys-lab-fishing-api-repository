package presenter

import (
	// "time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase"
)

type createFishPresenter struct{}

func NewCreateFishPresenter() usecase.CreateFishPresenter {
	return createFishPresenter{}
}

func (a createFishPresenter) Output(fish domain.Fish) usecase.CreateFishOutput {
	return usecase.CreateFishOutput{
		ID:        fish.ID().String(),
		Name:      fish.FishName(),
	}
}


