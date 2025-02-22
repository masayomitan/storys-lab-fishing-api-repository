package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
)

func (a FishingSpotSQL) FindOne(ctx context.Context, id string) (domain.FishingSpot, error) {
	var json = domain.FishingSpot{}

	if err := a.db.FindOneFishingSpot(ctx, a.tableName, id, &json); err != nil {
		return domain.FishingSpot{}, errors.Wrap(err, "error listing fishes")
	}

	var fishingSpot = domain.NewFishingSpot(
		json.ID,
		json.Name,
		json.ImageUrl,
		json.Description,
		json.AreaID,
		convertArea(json.Area),
		convertFishes(json.Fishes),
		convertTags(json.Tags),
	)

	fmt.Println("fishingSpot")
	fmt.Println(fishingSpot)

	return fishingSpot, nil
}

func (a FishingSpotSQL) FindAllByAreaId(ctx context.Context, id int) ([]domain.FishingSpot, error) {
	var json = []domain.FishingSpot{}

	if err := a.db.FindAllByAreaId(ctx, a.tableName, id, &json); err != nil {
		return []domain.FishingSpot{}, errors.Wrap(err, "error listing fishes")
	}

	var fishingSpots = make([]domain.FishingSpot, 0)
	for _, json := range json {
		var fishingSpot = domain.NewFishingSpot(
			json.ID,
			json.Name,
			json.ImageUrl,
			json.Description,
			json.AreaID,
			convertArea(json.Area),
			convertFishes(json.Fishes),
			convertTags(json.Tags),
		)
		fishingSpots = append(fishingSpots, fishingSpot)
	}
	
	fmt.Println("fishingSpot")
	fmt.Println(fishingSpots)

	return fishingSpots, nil
}

func (ga *GormAdapter) FindOneFishingSpot(ctx context.Context, table string, id string, result interface{}) error {
	return ga.DB.Table(table).Where("id = ?", id).
		Preload("Area").
		Preload("Area.Tides").
		Preload("Fishes").
		Preload("Tags").
		First(result).Error
}

func (ga *GormAdapter) FindAllByAreaId(ctx context.Context, table string, area_id int, result interface{}) error {
	return ga.DB.Table(table).Where("area_id = ?", area_id).
		Find(result).Error
}
