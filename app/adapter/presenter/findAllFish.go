package presenter

import (
	"time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase"
)

type createAccountPresenter struct{}

func NewCreateAccountPresenter() usecase.CreateFishPresenter {
	return createAccountPresenter{}
}

func (a createAccountPresenter) Output(fish domain.Fish) usecase.FindAllFishOutput {
	return usecase.FindAllFishOutput{
		ID:        fish.ID().String(),
		Name:      fish.Name(),
	}
}
