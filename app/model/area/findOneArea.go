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
		areaJSON.ImageUrl,
		convertFishingSpots(areaJSON.FishingSpots),
		convertTide(areaJSON.Tides),
	)
	fmt.Println(area)

	return area, nil
}

func (ga *GormAdapter) FindOneArea(ctx context.Context, table string, id string, result interface{}) error {
	return ga.DB.Table(table).Where("id = ?", id).
		Preload("FishingSpots").
		Preload("FishingSpots.Tags").
		// Preload("Areas.AreaImages").
		First(result).Error
}

func convertFishingSpots(fishingSpot []domain.FishingSpot) []domain.FishingSpot {
	var result []domain.FishingSpot
	for _, f := range fishingSpot {
		result = append(result, domain.FishingSpot{
			ID: 			f.ID,
			Name: 			f.Name,
			Description: 	f.Description,
			AreaId: 		f.AreaId,
			Tags: 			convertTags(f.Tags),
		})
	}
	return result
}

func convertTide(tide []domain.Tide) []domain.Tide {
	var result []domain.Tide
	for _, t := range tide {
		result = append(result, domain.Tide{
			ID: t.ID,
			AreaId:			 t.AreaId,
			PrefectureId:    t.PrefectureId,
		})
	}
	return result
}

func convertTags(tags []domain.Tag) []domain.Tag {
	var result []domain.Tag
	for _, t := range tags {
		result = append(result, domain.Tag{
			ID:        t.ID,
			Name:      t.Name,
		})
	}
	return result
}
