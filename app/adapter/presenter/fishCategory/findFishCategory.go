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

func (a FishCategoryPresenter) PresentOne(fishCategory domain.FishCategory) domain.FishCategory {
	return domain.FishCategory{
		ID:             fishCategory.ID,
		Name:           fishCategory.Name,
		EnglishName:	fishCategory.EnglishName,
		FamilyName:		fishCategory.FamilyName,
		Description:    fishCategory.Description,
		CreatedAt:		fishCategory.CreatedAt,
		UpdatedAt:		fishCategory.UpdatedAt,
	}
}

func (a FishCategoryPresenter) PresentAll(fishCategories []domain.FishCategory) []domain.FishCategory {
	var output = make([]domain.FishCategory, 0)
	for _, fishCategory := range fishCategories {
		output = append(output, domain.FishCategory{
			ID:             fishCategory.ID,
			Name:           fishCategory.Name,
			EnglishName:	fishCategory.EnglishName,
			FamilyName:		fishCategory.FamilyName,
			Description:    fishCategory.Description,
			CreatedAt:		fishCategory.CreatedAt,
			UpdatedAt:		fishCategory.UpdatedAt,
		})
	}
	return output
}

