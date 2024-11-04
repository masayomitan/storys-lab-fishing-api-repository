package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
)

func (a ToolSQL) FindOne(ctx context.Context, id int) (domain.Tool, error) {
	var json = domain.Tool{}

	if err := a.db.FindOneTool(ctx, a.tableName, id, &json); err != nil {
		return domain.Tool{}, errors.Wrap(err, "error listing fishes")
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
	)

	fmt.Println("tool")
	fmt.Println(tool)

	return tool, nil
}

func (ga *GormAdapter) FindOneTool(ctx context.Context, table string, area_id int, result interface{}) error {
	return ga.DB.Table(table).Where("id = ?", area_id).
		Find(result).Error
}
