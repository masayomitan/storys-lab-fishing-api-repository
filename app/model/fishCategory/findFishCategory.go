package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/utils"
)

func (a FishCategorySQL) FindAllByAdmin(ctx context.Context) ([]domain.FishCategory, error) {
	var json = make([]domain.FishCategory, 0)

	if err := a.db.FindAll(ctx, a.tableName, domain.FishCategory{}, &json); err != nil {
		return []domain.FishCategory{}, errors.Wrap(err, "error listing tools")
	}
	var fishCategories = make([]domain.FishCategory, 0)

	for _, json := range json {
		var fishCategory = domain.NewFishCategory(
			json.ID,
			json.Name,
			json.Description,
			json.CreatedAt,
			json.UpdatedAt,
		)
		fishCategories = append(fishCategories, fishCategory)
	}
	return fishCategories, nil
}

func (a FishCategorySQL) FindOneByAdmin(ctx context.Context, id int) (domain.FishCategory, error) {
	var json = domain.FishCategory{}

	if err := a.db.FindOneTool(ctx, a.tableName, id, &json); err != nil {
		return domain.FishCategory{}, errors.Wrap(err, "error listing tools")
	}

	var fishCategory = domain.NewFishCategory(
		json.ID,
		json.Name,
		json.Description,
		json.CreatedAt,
		json.UpdatedAt,
	)

	fmt.Println("")
	return fishCategory, nil
}

func (a FishCategorySQL) CreateByAdmin(ctx context.Context, requestParam domain.FishCategory) (domain.FishCategory, error) {
	utils.SetTimestamps(&requestParam)

	if err := a.db.Store(ctx, a.tableName, &requestParam); err != nil {
		return domain.FishCategory{}, errors.Wrap(err, "error creating fish-category")
	}

	fmt.Println("")
	return requestParam, nil
}

func (ga *GormAdapter) FindAll(ctx context.Context, table string, query interface{}, result interface{}) error {
    return ga.DB.Table(table).Where(query).
		Order("id asc").
		Find(result).Error
}

func (ga *GormAdapter) FindOneTool(ctx context.Context, table string, tool_id int, result interface{}) error {
	return ga.DB.Table(table).Where("id = ?", tool_id).
		First(result).Error
}


func (ga *GormAdapter) Store(ctx context.Context, table string, entity interface{}) error {
    return ga.DB.Table(table).Create(entity).Error
}
