package presenter

import (
	// "time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/tool"
)

type FindOneToolPresenter struct{}

func NewFindOneToolPresenter() usecase.FindOneToolPresenter {
	return FindOneToolPresenter{}
}

func (a FindOneToolPresenter) Output(tool domain.Tool) domain.Tool {
	return domain.Tool{
		ID:             tool.ID,
		Name:           tool.Name,
		Description:    tool.Description,
		ToolCategoryId: tool.ToolCategoryId,
        MaterialID:     tool.MaterialID,
        Size:           tool.Size,
        Weight:         tool.Weight,
        Durability:     tool.Durability,
        ToolUsage:      tool.ToolUsage,
        Price:          tool.Price,
        Maker:          tool.Maker,
        Recommend:      tool.Recommend,
        EasyFishing:    tool.EasyFishing,

	}
}

