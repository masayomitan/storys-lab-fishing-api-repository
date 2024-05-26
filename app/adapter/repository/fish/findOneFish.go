package repository

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
)

func (a FishSQL) FindOne(ctx context.Context, id string) (domain.Fish, error) {
	var fishJSON = domain.FishStruct{}
	fmt.Println(id)
	if err := a.db.FindOne(ctx, a.collectionName, id, &fishJSON); err != nil {
		return domain.Fish{}, errors.Wrap(err, "error listing fishes")
	}

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

	return fish, nil
}
