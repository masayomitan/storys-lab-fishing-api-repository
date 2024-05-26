package repository

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
)

func (a FishSQL) FindOne(ctx context.Context, id string) (domain.Fish, error) {
	var fishJSON = FishJSON{}
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
	)

	return fish, nil
}
