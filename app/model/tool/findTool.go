package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
)

func (a ToolSQL) FindAll(ctx context.Context) ([]domain.Tool, error) {
	var json = make([]domain.Tool, 0)

	if err := a.db.FindAll(ctx, a.tableName, domain.Tool{}, &json); err != nil {
		return []domain.Tool{}, errors.Wrap(err, "error listing tools")
	}
	var tools = make([]domain.Tool, 0)

	for _, json := range json {
		var tool = domain.NewTool(
			json.ID,
			json.Name,
			json.Description,
			json.ToolCategoryId,
			json.MaterialID,
			json.Size,
			json.Weight,
			json.Durability,
			json.ToolUsage,
			json.Price,
			json.Maker,
			json.Recommend,
			json.EasyFishing,
			json.ToolImages,
		)
		tools = append(tools, tool)
	}
	return tools, nil
}

func (a ToolSQL) FindOne(ctx context.Context, id int) (domain.Tool, error) {
	var json = domain.Tool{}

	if err := a.db.FindOneTool(ctx, a.tableName, id, &json); err != nil {
		return domain.Tool{}, errors.Wrap(err, "error listing tools")
	}

	var tool = domain.NewTool(
		json.ID,
		json.Name,
		json.Description,
		json.ToolCategoryId,
		json.MaterialID,
		json.Size,
		json.Weight,
		json.Durability,
		json.ToolUsage,
		json.Price,
		json.Maker,
		json.Recommend,
		json.EasyFishing,
		json.ToolImages,
	)

	fmt.Println("")
	return tool, nil
}

func (ga *GormAdapter) FindAll(ctx context.Context, table string, query interface{}, result interface{}) error {
    return ga.DB.Table(table).Where(query).
		Preload("ToolImages").
		Order("id asc").
		Find(result).Error
}

func (ga *GormAdapter) FindOneTool(ctx context.Context, table string, tool_id int, result interface{}) error {
	return ga.DB.Table(table).Where("id = ?", tool_id).
		// Preload("ToolImages").
		First(result).Error
}
