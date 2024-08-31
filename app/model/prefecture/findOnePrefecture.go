package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	// "gorm.io/gorm"
	"storys-lab-fishing-api/app/domain"
)


func (a PrefSQL) FindOne(ctx context.Context, id string) (domain.Pref, error) {
	var prefJSON = domain.Pref{}
	fmt.Println(id)

	if err := a.db.FindOnePref(ctx, a.collectionName, id, &prefJSON); err != nil {
		return domain.Pref{}, errors.Wrap(err, "error listing fishes")
	}

	var pref = domain.NewPref(
		prefJSON.ID,
		prefJSON.Name,
		prefJSON.NameKana,

		convertAreas(prefJSON.Areas),
	)
	fmt.Println(pref)

	return pref, nil
}

func (ga *GormAdapter) FindOnePref(ctx context.Context, table string, id string, result interface{}) error {
	return ga.DB.Table(table).Where("id = ?", id).
		Preload("Areas").
		// Preload("Areas.AreaImages").
		First(result).Error
}

func convertAreas(areas []domain.Area) []domain.Area {
	var result []domain.Area
	for _, a := range areas {
		result = append(result, domain.Area{
			ID:   a.ID,
			Name: a.Name,
			Description: a.Description,
			PrefectureId: a.PrefectureId,
		})
	}
	return result
}
