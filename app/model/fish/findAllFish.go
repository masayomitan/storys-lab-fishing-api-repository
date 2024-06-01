package model

import (
	"context"
	// "time"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
)

func (a FishSQL) FindAll(ctx context.Context) ([]domain.Fish, error) {
	var fishJSON = make([]domain.Fish, 0)

	if err := a.db.FindAll(ctx, a.collectionName, domain.Fish{}, &fishJSON); err != nil {
		return []domain.Fish{}, errors.Wrap(err, "error listing fishes")
	}

	var fishes = make([]domain.Fish, 0)

	for _, fishJSON := range fishJSON {
		var fish = domain.NewFish(
			fishJSON.ID,
			fishJSON.Name,
			fishJSON.FamilyName,
			fishJSON.ScientificName,
			fishJSON.FishCategoryId,
			fishJSON.Description,
			fishJSON.Length,
			fishJSON.Weight,
			fishJSON.Habitat,
			fishJSON.DepthRange,
			fishJSON.WaterTemperatureRange,
			fishJSON.ConservationStatus,
			convertFishCategory(fishJSON.FishCategory),
			convertFishingMethods(fishJSON.FishingMethods),
			convertDishes(fishJSON.Dishes),
		)

		fishes = append(fishes, fish)
	}

	return fishes, nil
}

func (ga *GormAdapter) FindAll(ctx context.Context, table string, query interface{}, result interface{}) error {
    return ga.DB.Table(table).Where(query).Find(result).Error
}
