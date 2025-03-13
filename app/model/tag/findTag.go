package model

import (
	"context"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
)


func (a TagSQL) FindByAdmin(ctx context.Context) ([]domain.Tag, error) {
	var json = make([]domain.Tag, 0)
	var tags = make([]domain.Tag, 0)

	if err := a.db.FindORM(ctx, a.tableName, &json); err != nil {
		return []domain.Tag{}, errors.Wrap(err, "error listing tags")
	}

	for _, json := range json {
		var tag = domain.NewTag(
			json.ID,
			json.Name,
			json.CreatedAt,
			json.UpdatedAt,
		)
		tags = append(tags, tag)
	}

	return tags, nil
}

func (ga *GormAdapter) FindORM(ctx context.Context, table string, result interface{}) error {
    return ga.DB.Table(table).
		Where("deleted_at IS NULL").
		Order("id asc").
		Find(result).Error
}
