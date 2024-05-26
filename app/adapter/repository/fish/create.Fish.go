package repository

import (
	"context"
	// "time"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
)

func (a FishSQL) Create(ctx context.Context, fish domain.Fish) (domain.Fish, error) {
	var fishJSON = FishJSON{
		ID:   fish.ID().String(),
		Name: fish.Name(),
		FamilyName: fish.FamilyName(),
		ScientificName: fish.ScientificName(),
		FishCategoryId: fish.FishCategoryId(),
		Description: fish.Description(),
	}

	if err := a.db.Store(ctx, a.collectionName, fishJSON); err != nil {
		return domain.Fish{}, errors.Wrap(err, "error creating account")
	}

	return fish, nil
}
