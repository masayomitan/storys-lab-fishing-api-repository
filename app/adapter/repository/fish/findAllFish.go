package repository

import (
	"context"
	// "time"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
)

func (a FishSQL) FindAll(ctx context.Context) ([]domain.Fish, error) {
	var fishJSON = make([]FishJSON, 0)

	if err := a.db.FindAll(ctx, a.collectionName, FishJSON{}, &fishJSON); err != nil {
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
		)

		fishes = append(fishes, fish)
	}

	return fishes, nil
}
