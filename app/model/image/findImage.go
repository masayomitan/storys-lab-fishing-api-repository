package model


import (
	"context"
	// "fmt"
	// "gorm.io/gorm"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
)

func (a ImageSQL) FindAllByAdmin(ctx context.Context, query map[string]interface{}) ([]domain.Image, error) {
	var json = make([]domain.Image, 0)

	if err := a.db.FindAll(ctx, a.tableName, query, &json); err != nil {
		return []domain.Image{}, errors.Wrap(err, "error listing images")
	}

	var images = make([]domain.Image, 0)

	for _, json := range json {
		var fish = domain.NewImage(
			json.ID,
			json.Name,
			json.ImageUrl,
			json.S3Url,
			json.ImageType,
			json.CreatedAt,
			json.UpdatedAt,
		)

		images = append(images, fish)
	}

	return images, nil
}

func (ga *GormAdapter) FindAll(ctx context.Context, table string, query interface{}, result interface{}) error {
    return ga.DB.Table(table).Where(query).
		Order("id asc").
		Find(result).Error
}
