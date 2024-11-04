package model

import (
	"context"
	// "time"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
)

func (a FishSQL) Create(ctx context.Context, fish domain.Fish) (domain.Fish, error) {
	var fishJSON = domain.Fish{
		ID:   fish.ID,
		Name: fish.Name,
		FamilyName: fish.FamilyName,
		ScientificName: fish.ScientificName,
		FishCategoryId: fish.FishCategoryId,
		Description: fish.Description,
	}

	if err := a.db.Store(ctx, a.tableName, fishJSON); err != nil {
		return domain.Fish{}, errors.Wrap(err, "error creating account")
	}

	return fish, nil
}

func (ga *GormAdapter) Store(ctx context.Context, table string, entity interface{}) error {
    return ga.DB.Table(table).Create(entity).Error
}
