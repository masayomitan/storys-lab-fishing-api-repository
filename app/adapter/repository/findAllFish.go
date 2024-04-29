package repository

import (
	"context"
	// "time"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
)

type FishJSON struct {
	ID string `json:"id"`
	Name string `json:"fish_name"`
	FamilyName string `json:"family_name"`
	ScientificName string `json:"scientific_name"`
	FishCategory int `json:"fish_category"`
	Description string `json:"description"`
}

type FishSQL struct {
	collectionName string
	db             SQL
}

func NewFishSQL(db SQL) FishSQL {
	return FishSQL{
		db:             db,
		collectionName: "fishes",
	}
}

func (a FishSQL) Create(ctx context.Context, fish domain.Fish) (domain.Fish, error) {
	var fishJSON = FishJSON{
		ID:   fish.ID().String(),
		Name: fish.Name(),
		FamilyName: fish.FamilyName(),
		ScientificName: fish.ScientificName(),
		FishCategory: fish.FishCategory(),
		Description: fish.Description(),
	}

	if err := a.db.Store(ctx, a.collectionName, fishJSON); err != nil {
		return domain.Fish{}, errors.Wrap(err, "error creating account")
	}

	return fish, nil
}

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
			fishJSON.FishCategory,
			fishJSON.Description,
		)

		fishes = append(fishes, fish)
	}

	return fishes, nil
}
