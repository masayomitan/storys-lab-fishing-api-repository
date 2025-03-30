package model

import (
	"context"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
)


func (a MaterialSQL) FindByAdmin(ctx context.Context) ([]domain.Material, error) {
	var json = make([]domain.Material, 0)
	var materials = make([]domain.Material, 0)

	if err := a.db.FindORM(ctx, a.tableName, &json); err != nil {
		return []domain.Material{}, errors.Wrap(err, "error listing tags")
	}

	for _, json := range json {
		var material = domain.NewMaterial(
			json.ID,
			json.Name,
			json.Description,
			json.CreatedAt,
			json.UpdatedAt,
		)
		materials = append(materials, material)
	}

	return materials, nil
}

func (ga *GormAdapter) FindORM(ctx context.Context, table string, result interface{}) error {
    return ga.DB.Table(table).
		Where("deleted_at IS NULL").
		Order("id asc").
		Find(result).Error
}
