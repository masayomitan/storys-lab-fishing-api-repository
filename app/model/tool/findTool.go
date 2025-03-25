package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
)

func (a ToolSQL) Find(ctx context.Context) ([]domain.Tool, error) {
	var json = make([]domain.Tool, 0)

	if err := a.db.FindORM(ctx, a.tableName, domain.Tool{}, &json); err != nil {
		return []domain.Tool{}, errors.Wrap(err, "error listing tools")
	}
	var tools = make([]domain.Tool, 0)

	for _, json := range json {
		var tool = domain.NewTool(
			json.ID,
			json.Name,
			json.Description,
			json.ToolCategoryID,
			json.MaterialID,
			json.Size,
			json.Weight,
			json.Durability,
			json.ToolUsage,
			json.Price,
			json.Maker,
			json.Recommend,
			json.EasyFishing,
			json.Images,
		)
		tools = append(tools, tool)
	}
	return tools, nil
}

func (a ToolSQL) FindOne(ctx context.Context, id int) (domain.Tool, error) {
	var json = domain.Tool{}

	if err := a.db.FindOneORM(ctx, a.tableName, id, &json); err != nil {
		return domain.Tool{}, errors.Wrap(err, "error listing tools")
	}

	var tool = domain.NewTool(
		json.ID,
		json.Name,
		json.Description,
		json.ToolCategoryID,
		json.MaterialID,
		json.Size,
		json.Weight,
		json.Durability,
		json.ToolUsage,
		json.Price,
		json.Maker,
		json.Recommend,
		json.EasyFishing,
		json.Images,
	)

	fmt.Println("")
	return tool, nil
}

func (a ToolSQL) FindByAdmin(ctx context.Context) ([]domain.Tool, error) {
	var json = make([]domain.Tool, 0)

	if err := a.db.FindORM(ctx, a.tableName, domain.Tool{}, &json); err != nil {
		return []domain.Tool{}, errors.Wrap(err, "error listing tools")
	}
	var tools = make([]domain.Tool, 0)

	for _, json := range json {
		var tool = domain.NewTool(
			json.ID,
			json.Name,
			json.Description,
			json.ToolCategoryID,
			json.MaterialID,
			json.Size,
			json.Weight,
			json.Durability,
			json.ToolUsage,
			json.Price,
			json.Maker,
			json.Recommend,
			json.EasyFishing,
			json.Images,
		)
		tools = append(tools, tool)
	}
	return tools, nil
}

func (a ToolSQL) FindOneByAdmin(ctx context.Context, id int) (domain.Tool, error) {
	var json = domain.Tool{}

	if err := a.db.FindOneORM(ctx, a.tableName, id, &json); err != nil {
		return domain.Tool{}, errors.Wrap(err, "error listing tool_categories")
	}

	var tool = domain.NewTool(
		json.ID,
		json.Name,
		json.Description,
		json.ToolCategoryID,
		json.MaterialID,
		json.Size,
		json.Weight,
		json.Durability,
		json.ToolUsage,
		json.Price,
		json.Maker,
		json.Recommend,
		json.EasyFishing,
		json.Images,
	)

	fmt.Println("")
	return tool, nil
}

func (ga *GormAdapter) FindORM(ctx context.Context, table string, query interface{}, result interface{}) error {
    return ga.DB.Table(table).Where(query).
		Preload("Images").
		Order("id asc").
		Find(result).Error
}

func (ga *GormAdapter) FindOneORM(ctx context.Context, table string, tool_id int, result interface{}) error {
	return ga.DB.Table(table).
		Where("id = ?", tool_id).
		Preload("Images").
		First(result).Error
}
