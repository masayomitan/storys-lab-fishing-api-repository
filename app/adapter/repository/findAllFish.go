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
    ID   string `json:"id"`
    FishName string `json:"fish_name"`
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
		FishName: fish.FishName(),
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
			fishJSON.FishName,
		)

		fishes = append(fishes, fish)
	}

	return fishes, nil
}
