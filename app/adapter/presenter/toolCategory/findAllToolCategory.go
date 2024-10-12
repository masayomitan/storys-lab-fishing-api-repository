package presenter

import (
	// "time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/toolCategory"
)

type findAllToolCategoryPresenter struct{}

func NewFindAllToolCategoryPresenter() usecase.FindAllToolCategoryPresenter {
	return findAllToolCategoryPresenter{}
}

func (a findAllToolCategoryPresenter) Output(toolCategories []domain.ToolCategory) []domain.ToolCategory {
	var output = make([]domain.ToolCategory, 0)
	for _, toolCategory := range toolCategories {
		output = append(output, domain.ToolCategory{
			ID:				toolCategory.ID,
			Name:			toolCategory.Name,
			Description: 	toolCategory.Description,
			Tools:			toolCategory.Tools,
		})
	}
	return output
}
