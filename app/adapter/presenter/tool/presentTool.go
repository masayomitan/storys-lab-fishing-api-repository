package presenter

import (
	// "time"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/tool"
)

type toolPresenter struct{}

func NewToolPresenter() usecase.ToolPresenter {
	return toolPresenter{}
}

func (a toolPresenter) Present(tool domain.Tool) domain.Tool {
	return domain.Tool{
		ID:             tool.ID,
		Name:           tool.Name,
		Description:    tool.Description,
		ToolCategoryID: tool.ToolCategoryID,
        MaterialID:     tool.MaterialID,
        Size:           tool.Size,
        Weight:         tool.Weight,
        ToolUsage:      tool.ToolUsage,
        Price:          tool.Price,
        Maker:          tool.Maker,
        Recommend:      tool.Recommend,
        EasyFishing:    tool.EasyFishing,
		Images:     	tool.Images,
	}
}


func (a toolPresenter) Presents(tools []domain.Tool) []domain.Tool {
	var output = make([]domain.Tool, 0)
	for _, tool := range tools {
		output = append(output, domain.Tool{
			ID:             tool.ID,
			Name:           tool.Name,
			Description:    tool.Description,
			ToolCategoryID: tool.ToolCategoryID,
			MaterialID:     tool.MaterialID,
			Size:           tool.Size,
			Weight:         tool.Weight,
			ToolUsage:      tool.ToolUsage,
			Price:          tool.Price,
			Maker:          tool.Maker,
			Recommend:      tool.Recommend,
			EasyFishing:    tool.EasyFishing,
			Images:     	tool.Images,
		})
	}
	return output
}
