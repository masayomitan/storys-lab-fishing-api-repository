package model

import (
	"context"
	// "fmt"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
)

func (a AreaSQL) FindAll(ctx context.Context) ([]domain.Area, error) {
	var areaJSON = make([]domain.Area, 0)

	if err := a.db.FindAll(ctx, a.collectionName, domain.Area{}, &areaJSON); err != nil {
		return []domain.Area{}, errors.Wrap(err, "error listing areas")
	}

	var areas = make([]domain.Area, 0)

	for _, areaJSON := range areaJSON {
		var area = domain.NewArea(
			areaJSON.ID,
			areaJSON.Name,
			areaJSON.Description,
			areaJSON.PrefectureId,
			convertFishingSpots(areaJSON.FishingSpots),
			convertTide(areaJSON.Tides),
		)

		areas = append(areas, area)
	}

	return areas, nil
}

func (ga *GormAdapter) FindAll(ctx context.Context, table string, query interface{}, result interface{}) error {
    return ga.DB.Table(table).Where(query).
		// Preload("AreaImages").
		Find(result).Error
}
