package model

import (
	"context"
	"fmt"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
)

func (a ToolCategorySQL) FindAll(ctx context.Context) ([]domain.ToolCategory, error) {
	var json = make([]domain.ToolCategory, 0)

	if err := a.db.FindAll(ctx, a.collectionName, domain.ToolCategory{}, &json); err != nil {
		return []domain.ToolCategory{}, errors.Wrap(err, "error listing tool_categories")
	}

	var tools = make([]domain.ToolCategory, 0)

	for _, json := range json {
		var fish = domain.NewToolCategory(
			json.ID,
			json.Name,
			json.Description,
			convertTools(json.Tools),
		)

		tools = append(tools, fish)
	}
	fmt.Println(tools)
	return tools, nil
}

func (ga *GormAdapter) FindAll(ctx context.Context, table string, query interface{}, result interface{}) error {
    return ga.DB.Table(table).Where(query).
		Preload("Tools").
		Find(result).Error
}
