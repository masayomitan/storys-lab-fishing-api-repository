package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	// "gorm.io/gorm"
	"storys-lab-fishing-api/app/domain"
)


func (a AreaSQL) FindOne(ctx context.Context, id string) (domain.Area, error) {
	var areaJSON = domain.Area{}
	fmt.Println(id)

	if err := a.db.FindOneArea(ctx, a.collectionName, id, &areaJSON); err != nil {
		return domain.Area{}, errors.Wrap(err, "error listing fishes")
	}

	var area = domain.NewArea(
		areaJSON.ID,
		areaJSON.Name,
		areaJSON.Description,
		areaJSON.PrefectureId,
		convertFishingSpots(areaJSON.FishingSpots),
	)
	fmt.Println(area)

	return area, nil
}

func (ga *GormAdapter) FindOneArea(ctx context.Context, table string, id string, result interface{}) error {
	return ga.DB.Table(table).Where("id = ?", id).
		Preload("FishingSpots").
		// Preload("Areas.AreaImages").
		First(result).Error
}

func convertFishingSpots(fishingSpot []domain.FishingSpot) []domain.FishingSpot {
	var result []domain.FishingSpot
	for _, f := range fishingSpot {
		result = append(result, domain.FishingSpot{
			ID: f.ID,
			Name: f.Name,
			Description: f.Description,
			AreaId: f.AreaId,
		})
	}
	return result
}
