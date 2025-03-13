package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
)


func (a FishingSpotSQL) Find(ctx context.Context) ([]domain.FishingSpot, error) {
	var json = make([]domain.FishingSpot, 0)
	var fishingSpots = make([]domain.FishingSpot, 0)
	if err := a.db.FindByPrefectureIdORM(ctx, a.tableName, domain.FishingSpot{}, &json); err != nil {
		return []domain.FishingSpot{}, errors.Wrap(err, "error listing areas")
	}

	for _, json := range json {
		var area = domain.NewFishingSpot(
			json.ID,
			json.Name,
			json.Description,
			json.AreaID,
			json.RecommendedFishingMethods,
			json.CreatedAt,
			json.UpdatedAt,
			convertArea(json.Area),
			convertFishes(json.Fishes),
			convertTags(json.Tags),
			convertImages(json.Images),
		)
		fishingSpots = append(fishingSpots, area)
	}

	return fishingSpots, nil
}

func (a FishingSpotSQL) FindOne(ctx context.Context, id int) (domain.FishingSpot, error) {
	var json = domain.FishingSpot{}
	fmt.Println("fishingSpot FindOne")
	if err := a.db.FindOneORM(ctx, a.tableName, id, &json); err != nil {
		return domain.FishingSpot{}, errors.Wrap(err, "error listing fishes")
	}

	var fishingSpot = domain.NewFishingSpot(
		json.ID,
		json.Name,
		json.Description,
		json.AreaID,
		json.RecommendedFishingMethods,
		json.CreatedAt,
		json.UpdatedAt,
		convertArea(json.Area),
		convertFishes(json.Fishes),
		convertTags(json.Tags),
		convertImages(json.Images),
	)

	return fishingSpot, nil
}

func (a FishingSpotSQL) FindByAreaId(ctx context.Context, area_id int) ([]domain.FishingSpot, error) {
	var json = []domain.FishingSpot{}
	fmt.Println("fishingSpot FindByAreaId")
	if err := a.db.FindByAreaIdORM(ctx, a.tableName, area_id, &json); err != nil {
		return []domain.FishingSpot{}, errors.Wrap(err, "error listing fishes")
	}

	var fishingSpots = make([]domain.FishingSpot, 0)
	for _, json := range json {
		var fishingSpot = domain.NewFishingSpot(
			json.ID,
			json.Name,
			json.Description,
			json.AreaID,
			json.RecommendedFishingMethods,
			json.CreatedAt,
			json.UpdatedAt,
			convertArea(json.Area),
			convertFishes(json.Fishes),
			convertTags(json.Tags),
			convertImages(json.Images),
		)
		fishingSpots = append(fishingSpots, fishingSpot)
	}

	return fishingSpots, nil
}

func (a FishingSpotSQL) FindByAdmin(ctx context.Context) ([]domain.FishingSpot, error) {
	var json = make([]domain.FishingSpot, 0)
	var fishingSpots = make([]domain.FishingSpot, 0)

	if err := a.db.FindORM(ctx, a.tableName, &json); err != nil {
		return []domain.FishingSpot{}, errors.Wrap(err, "error listing fishingSpots")
	}

	for _, json := range json {
		var fishingSpot = domain.NewFishingSpot(
			json.ID,
			json.Name,
			json.Description,
			json.AreaID,
			json.RecommendedFishingMethods,
			json.CreatedAt,
			json.UpdatedAt,
			convertArea(json.Area),
			convertFishes(json.Fishes),
			convertTags(json.Tags),
			convertImages(json.Images),
		)
		fishingSpots = append(fishingSpots, fishingSpot)
	}

	return fishingSpots, nil
}

func (a FishingSpotSQL) FindOneByAdmin(ctx context.Context, id int) (domain.FishingSpot, error) {
	var json = domain.FishingSpot{}
	fmt.Println(id)

	if err := a.db.FindOneORM(ctx, a.tableName, id, &json); err != nil {
		return domain.FishingSpot{}, errors.Wrap(err, "error listing fishingSpot")
	}

	var fishingSpot = domain.NewFishingSpot(
		json.ID,
		json.Name,
		json.Description,
		json.AreaID,
		json.RecommendedFishingMethods,
		json.CreatedAt,
		json.UpdatedAt,
		convertArea(json.Area),
		convertFishes(json.Fishes),
		convertTags(json.Tags),
		convertImages(json.Images),
	)
	fmt.Println(fishingSpot)

	return fishingSpot, nil
}

func (ga *GormAdapter) FindOneORM(ctx context.Context, table string, id int, result interface{}) error {
	return ga.DB.Table(table).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Preload("Area", "deleted_at IS NULL").
		Preload("Area.Images", "deleted_at IS NULL").
		Preload("Tags", "deleted_at IS NULL").
		Preload("Images", "deleted_at IS NULL").
		First(result).Error
}

func (ga *GormAdapter) FindByAreaIdORM(ctx context.Context, table string, area_id int, result interface{}) error {
	return ga.DB.Table(table).
		Where("area_id = ?", area_id).
		Where("deleted_at IS NULL").
		Preload("Images", "deleted_at IS NULL").
		Find(result).Error
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
