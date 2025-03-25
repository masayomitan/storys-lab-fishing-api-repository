package presenter

import (
	// "time"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/toolCategory"
)

type ToolCategoryPresenter struct{}

// コンストラクタ
func NewToolCategoryPresenter() usecase.ToolCategoryPresenter {
	return ToolCategoryPresenter{}
}

func (a ToolCategoryPresenter) Present(toolCategory domain.ToolCategory) domain.ToolCategory {
	return domain.ToolCategory{
		ID:             toolCategory.ID,
		Name:           toolCategory.Name,
		Description:    toolCategory.Description,
		CreatedAt:		toolCategory.CreatedAt,
		UpdatedAt:		toolCategory.UpdatedAt,
	}
}

func (a ToolCategoryPresenter) Presents(toolCategories []domain.ToolCategory) []domain.ToolCategory {
	var output = make([]domain.ToolCategory, 0)
	for _, toolCategory := range toolCategories {
		output = append(output, domain.ToolCategory{
			ID:             toolCategory.ID,
			Name:           toolCategory.Name,
			Description:    toolCategory.Description,
			CreatedAt:		toolCategory.CreatedAt,
			UpdatedAt:		toolCategory.UpdatedAt,
		})
	}
	return output
}

