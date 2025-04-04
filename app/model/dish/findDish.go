package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
)

func (a DishSQL) Find(ctx context.Context) ([]domain.Dish, error) {
	var json = make([]domain.Dish, 0)

	if err := a.db.FindORM(ctx, a.tableName, domain.Dish{}, &json); err != nil {
		return []domain.Dish{}, errors.Wrap(err, "error listing tools")
	}
	var dishes = make([]domain.Dish, 0)

	for _, json := range json {
		var dish = domain.NewDish(
			json.ID,
			json.Name,
			json.Description,
			json.Ingredients,
			json.Kind,
			json.Level,
			json.CreatedAt,
			json.UpdatedAt,
			json.Images,
		)
		dishes = append(dishes, dish)
	}
	return dishes, nil
}

func (a DishSQL) FindOne(ctx context.Context, id int) (domain.Dish, error) {
	var json = domain.Dish{}

	if err := a.db.FindOneORM(ctx, a.tableName, id, &json); err != nil {
		return domain.Dish{}, errors.Wrap(err, "error listing tools")
	}

	var dish = domain.NewDish(
		json.ID,
		json.Name,
		json.Description,
		json.Ingredients,
		json.Kind,
		json.Level,
		json.CreatedAt,
		json.UpdatedAt,
		json.Images,
	)

	return dish, nil
}

func (a DishSQL) FindByAdmin(ctx context.Context) ([]domain.Dish, error) {
	var json = make([]domain.Dish, 0)

	if err := a.db.FindORM(ctx, a.tableName, domain.Dish{}, &json); err != nil {
		return []domain.Dish{}, errors.Wrap(err, "error listing tools")
	}
	var dishes = make([]domain.Dish, 0)

	for _, json := range json {
		var dish = domain.NewDish(
			json.ID,
			json.Name,
			json.Description,
			json.Ingredients,
			json.Kind,
			json.Level,
			json.CreatedAt,
			json.UpdatedAt,
			json.Images,
		)
		dishes = append(dishes, dish)
	}
	return dishes, nil
}

func (a DishSQL) FindOneByAdmin(ctx context.Context, id int) (domain.Dish, error) {
	var json = domain.Dish{}

	if err := a.db.FindOneORM(ctx, a.tableName, id, &json); err != nil {
		return domain.Dish{}, errors.Wrap(err, "error listing tools")
	}

	var dish = domain.NewDish(
		json.ID,
		json.Name,
		json.Description,
		json.Ingredients,
		json.Kind,
		json.Level,
		json.CreatedAt,
		json.UpdatedAt,
		json.Images,
	)

	fmt.Println("")
	return dish, nil
}

func (ga *GormAdapter) FindORM(ctx context.Context, table string, query interface{}, result interface{}) error {
    return ga.DB.Table(table).Where(query).
		Where(query).
		Where("deleted_at IS NULL").
		Preload("Images").
		Order("id asc").
		Find(result).Error
}

func (ga *GormAdapter) FindOneORM(ctx context.Context, table string, tool_id int, result interface{}) error {
	return ga.DB.Table(table).
		Where("deleted_at IS NULL").
		Where("id = ?", tool_id).
		Preload("Images").
		First(result).Error
}
