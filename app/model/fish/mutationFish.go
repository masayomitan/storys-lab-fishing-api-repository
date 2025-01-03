package model

import (
	"context"
	// "time"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/utils"
)

func (a FishSQL) CreateByAdmin(ctx context.Context, requestParam domain.Fish) (domain.Fish, error) {
	utils.SetTimestamps(&requestParam)

	if err := a.db.Store(ctx, a.tableName, &requestParam); err != nil {
		return domain.Fish{}, errors.Wrap(err, "error creating account")
	}

	return requestParam, nil
}

func (ga *GormAdapter) Store(ctx context.Context, table string, entity interface{}) error {
    return ga.DB.Table(table).Create(entity).Error
}
