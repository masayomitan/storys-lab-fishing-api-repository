package model

import (
	"context"
	"fmt"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
)

func (a ToolCategorySQL) Find(ctx context.Context) ([]domain.ToolCategory, error) {
	var json = make([]domain.ToolCategory, 0)

	if err := a.db.FindORM(ctx, a.tableName, domain.ToolCategory{}, &json); err != nil {
		return []domain.ToolCategory{}, errors.Wrap(err, "error listing tool_categories")
	}

	var toolCategories = make([]domain.ToolCategory, 0)

	for _, json := range json {
		var toolCategory = domain.NewToolCategory(
			json.ID,
			json.Name,
			json.Description,
			convertTools(json.Tools),
		)

		toolCategories = append(toolCategories, toolCategory)
	}
	fmt.Println(toolCategories)
	return toolCategories, nil
}

func (a ToolCategorySQL) FindOne(ctx context.Context, id int) (domain.ToolCategory, error) {
	var json = domain.ToolCategory{}

	if err := a.db.FindOneORM(ctx, a.tableName, id, &json); err != nil {
		return domain.ToolCategory{}, errors.Wrap(err, "error listing tool_categories")
	}

	var toolCategory = domain.NewToolCategory(
		json.ID,
		json.Name,
		json.Description,
		convertTools(json.Tools),
	)

	return toolCategory, nil
}

func (a ToolCategorySQL) FindByAdmin(ctx context.Context) ([]domain.ToolCategory, error) {
	var json = make([]domain.ToolCategory, 0)

	if err := a.db.FindORM(ctx, a.tableName, domain.ToolCategory{}, &json); err != nil {
		return []domain.ToolCategory{}, errors.Wrap(err, "error listing tool_categories")
	}

	var toolCategories = make([]domain.ToolCategory, 0)

	for _, json := range json {
		var toolCategory = domain.NewToolCategory(
			json.ID,
			json.Name,
			json.Description,
			convertTools(json.Tools),
		)

		toolCategories = append(toolCategories, toolCategory)
	}
	fmt.Println(toolCategories)
	return toolCategories, nil
}

func (a ToolCategorySQL) FindOneByAdmin(ctx context.Context, id int) (domain.ToolCategory, error) {
	var json = domain.ToolCategory{}

	if err := a.db.FindOneORM(ctx, a.tableName, id, &json); err != nil {
		return domain.ToolCategory{}, errors.Wrap(err, "error listing tool_categories")
	}

	var toolCategory = domain.NewToolCategory(
		json.ID,
		json.Name,
		json.Description,
		convertTools(json.Tools),
	)

	return toolCategory, nil
}

func (ga *GormAdapter) FindORM(ctx context.Context, table string, query interface{}, result interface{}) error {
    return ga.DB.Table(table).
		Where(query).
		Where("deleted_at IS NULL").
		Preload("Tools", "deleted_at IS NULL").
		Preload("Tools.Images").
		Find(result).Error
}

func (ga *GormAdapter) FindOneORM(ctx context.Context, table string, tool_category_id int, result interface{}) error {
	return ga.DB.Table(table).
		Where("id = ?", tool_category_id).
		Where("deleted_at IS NULL").
		Preload("Tools", "deleted_at IS NULL").
		Preload("Tools.Images").
		First(result).Error
}
