package model

import (
	"storys-lab-fishing-api/app/domain"
)

func convertTools(tools []domain.Tool) []domain.Tool {
	var result []domain.Tool
	for _, t := range tools {
		result = append(result, domain.Tool{
			ID:        t.ID,
			Name:      t.Name,
			Description: t.Description,
			ToolImages: t.ToolImages,
		})
	}
	return result
}
