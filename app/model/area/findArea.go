package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
)


func (a AreaSQL) FindOne(ctx context.Context, id int) (domain.Area, error) {
	var json = domain.Area{}
	fmt.Println(id)

	if err := a.db.FindOneAreaORM(ctx, a.tableName, id, &json); err != nil {
		return domain.Area{}, errors.Wrap(err, "error listing fishes")
	}

	var area = domain.NewArea(
		json.ID,
		json.Name,
		json.Description,
		json.PrefectureID,
		json.CreatedAt,
		json.UpdatedAt,
		convertFishingSpots(json.FishingSpots),
		convertTide(json.Tides),
		convertImages(json.Images),
	)
	fmt.Println(area)

	return area, nil
}


func (a AreaSQL) FindAll(ctx context.Context) ([]domain.Area, error) {
	var json = make([]domain.Area, 0)
	var areas = make([]domain.Area, 0)
	if err := a.db.FindByPrefectureIdORM(ctx, a.tableName, domain.Area{}, &json); err != nil {
		return []domain.Area{}, errors.Wrap(err, "error listing areas")
	}

	for _, json := range json {
		var area = domain.NewArea(
			json.ID,
			json.Name,
			json.Description,
			json.PrefectureID,
			json.CreatedAt,
			json.UpdatedAt,
			convertFishingSpots(json.FishingSpots),
			convertTide(json.Tides),
			convertImages(json.Images),
		)
		areas = append(areas, area)
	}

	return areas, nil
}

func (a AreaSQL) FindByAdmin(ctx context.Context) ([]domain.Area, error) {
	var json = make([]domain.Area, 0)
	var areas = make([]domain.Area, 0)

	if err := a.db.FindORM(ctx, a.tableName, &json); err != nil {
		return []domain.Area{}, errors.Wrap(err, "error listing areas")
	}

	for _, json := range json {
		var area = domain.NewArea(
			json.ID,
			json.Name,
			json.Description,
			json.PrefectureID,
			json.CreatedAt,
			json.UpdatedAt,
			convertFishingSpots(json.FishingSpots),
			convertTide(json.Tides),
			convertImages(json.Images),
		)
		areas = append(areas, area)
	}

	return areas, nil
}

func (a AreaSQL) FindOneByAdmin(ctx context.Context, id int) (domain.Area, error) {
	var json = domain.Area{}
	fmt.Println(id)

	if err := a.db.FindOneAreaORM(ctx, a.tableName, id, &json); err != nil {
		return domain.Area{}, errors.Wrap(err, "error listing fishes")
	}

	var area = domain.NewArea(
		json.ID,
		json.Name,
		json.Description,
		json.PrefectureID,
		json.CreatedAt,
		json.UpdatedAt,
		convertFishingSpots(json.FishingSpots),
		convertTide(json.Tides),
		convertImages(json.Images),
	)
	fmt.Println(area)

	return area, nil
}

func (ga *GormAdapter) FindORM(ctx context.Context, table string, result interface{}) error {
    return ga.DB.Table(table).
		Where("deleted_at IS NULL").
		Preload("Images", "deleted_at IS NULL").
		Order("id asc").
		Find(result).Error
}

func (ga *GormAdapter) FindByPrefectureIdORM(ctx context.Context, table string, query interface{}, result interface{}) error {
    return ga.DB.Table(table).Where(query).
		Where("deleted_at IS NULL").
		Preload("Images", "deleted_at IS NULL").
		Order("id asc").
		Find(result).Error
}


func (ga *GormAdapter) FindOneAreaORM(ctx context.Context, table string, id int, result interface{}) error {
	return ga.DB.Table(table).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Preload("Images", "deleted_at IS NULL").
		Preload("FishingSpots", "deleted_at IS NULL").
		Preload("FishingSpots.Tags", "deleted_at IS NULL").
		Preload("FishingSpots.Images", "deleted_at IS NULL").
		First(result).Error
}

