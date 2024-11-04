package model

import (
	"context"
	// "time"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
)

func (a AreaSQL) Create(ctx context.Context, area domain.Area) (domain.Area, error) {
	var areaJSON = domain.Area{
		ID:   area.ID,
		Name: area.Name,
		Description: area.Description,
		PrefectureId: area.PrefectureId,
	}

	if err := a.db.Store(ctx, a.tableName, areaJSON); err != nil {
		return domain.Area{}, errors.Wrap(err, "error creating account")
	}

	return area, nil
}

func (ga *GormAdapter) Store(ctx context.Context, table string, entity interface{}) error {
    return ga.DB.Table(table).Create(entity).Error
}
