package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	// "gorm.io/gorm"
	"storys-lab-fishing-api/app/domain"
)


func (a FishingSpotSQL) FindOne(ctx context.Context, id string) (domain.FishingSpot, error) {
	var json = domain.FishingSpot{}

	if err := a.db.FindOneFish(ctx, a.tableName, id, &json); err != nil {
		return domain.FishingSpot{}, errors.Wrap(err, "error listing fishes")
	}

	var fishingSpot = domain.NewFishingSpot(
		json.ID,
		json.Name,
		json.Description,
		json.AreaId,

		convertFishes(json.Fishes),
		convertArea(json.Area),
		convertTide(json.Tide),
	)

	fmt.Println("fishStruct")
	fmt.Println(fishingSpot)

	return fishingSpot, nil
}

func (ga *GormAdapter) FindOneFish(ctx context.Context, table string, id string, result interface{}) error {
	return ga.DB.Table(table).Where("id = ?", id).
		Preload("Fishes").
		Preload("Areas.Tides").
		First(result).Error
}


func convertFishes(fishes []domain.Fish) []domain.Fish {
	var result []domain.Fish
	for _, f := range fishes {
		result = append(result, domain.Fish{
			ID:   f.ID,
			Name: f.Name,
			FamilyName: f.FamilyName,
		})
	}
	return result
}

func convertArea(a domain.Area) domain.Area {
	return domain.Area {
		ID: a.ID,
		PrefectureId: a.PrefectureId,
	}
}

func convertTide(t domain.Tide) domain.Tide {
	return domain.Tide {
		ID: t.ID,
		AreaId: t.AreaId,
		PrefectureId: t.PrefectureId,
	}
}

