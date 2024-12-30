package presenter

import (
	// "time"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/fishCategory"
)

type FishCategoryPresenter struct{}

// コンストラクタ
func NewFishCategoryPresenter() usecase.FishCategoryPresenter {
	return FishCategoryPresenter{}
}

func (a FishCategoryPresenter) PresentOne(FishCategory domain.FishCategory) domain.FishCategory {
	return domain.FishCategory{
		ID:             FishCategory.ID,
		Name:           FishCategory.Name,
		Description:    FishCategory.Description,
	}
}

func (a FishCategoryPresenter) PresentAll(fishCategories []domain.FishCategory) []domain.FishCategory {
	var output = make([]domain.FishCategory, 0)
	for _, fishCategory := range fishCategories {
		output = append(output, domain.FishCategory{
			ID:             fishCategory.ID,
			Name:           fishCategory.Name,
			Description:    fishCategory.Description,
		})
	}
	return output
}

