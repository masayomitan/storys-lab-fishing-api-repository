package repository

import (
	"context"
	// "time"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
)

func (a FishSQL) FindAll(ctx context.Context) ([]domain.Fish, error) {
	var fishJSON = make([]domain.FishStruct, 0)

	if err := a.db.FindAll(ctx, a.collectionName, domain.FishStruct{}, &fishJSON); err != nil {
		return []domain.Fish{}, errors.Wrap(err, "error listing fishes")
	}

	var fishes = make([]domain.Fish, 0)

	for _, fishJSON := range fishJSON {
		var fish = domain.NewFish(
			domain.FishID(fishJSON.ID),
			fishJSON.Name,
			fishJSON.FamilyName,
			fishJSON.ScientificName,
			fishJSON.FishCategoryId,
			fishJSON.Description,
			fishJSON.Length,
			fishJSON.Weight,
			fishJSON.Habitat,
			fishJSON.Depth_range,
			fishJSON.Water_temperature_range,
			fishJSON.Conservation_status,
		)

		fishes = append(fishes, fish)
	}

	return fishes, nil
}
