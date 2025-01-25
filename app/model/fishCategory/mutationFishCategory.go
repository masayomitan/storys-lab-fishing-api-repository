package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/utils"
)

func (a FishCategorySQL) CreateByAdmin(ctx context.Context, requestParam domain.FishCategory) (domain.FishCategory, error) {
	utils.SetCreateTimestamps(&requestParam)

	if err := a.db.Store(ctx, a.tableName, &requestParam); err != nil {
		return domain.FishCategory{}, errors.Wrap(err, "error creating fish-category")
	}

	fmt.Println("")
	return requestParam, nil
}

func (a FishCategorySQL) UpdateByAdmin(ctx context.Context, requestParam domain.FishCategory, id int) (domain.FishCategory, error) {
	utils.SetUpdateTimestamps(&requestParam)

	if err := a.db.Update(ctx, a.tableName, &requestParam, id); err != nil {
		return domain.FishCategory{}, errors.Wrap(err, "error creating fish-category")
	}

	fmt.Println("")
	return requestParam, nil
}

func (a FishCategorySQL) DeleteByAdmin(ctx context.Context, id int) error {

	if err := a.db.Delete(ctx, a.tableName, id); err != nil {
		return errors.Wrap(err, "error deleting fish-category")
	}

	return nil
}

func (ga *GormAdapter) Store(ctx context.Context, table string, entity interface{}) error {
    return ga.DB.Table(table).
		Create(entity).Error
}

func (ga *GormAdapter) Update(ctx context.Context, table string, entity interface{}, id int) error {
    return ga.DB.Table(table).
		Where("id = ?", id).
		Updates(entity).Error
}

func (ga *GormAdapter) Delete(ctx context.Context, table string, id int) error {
	// モデルを直接指定
	return ga.DB.Table(table).
		Where("id = ?", id).
		Update("deleted_at", gorm.Expr("CURRENT_TIMESTAMP")).Error
}
