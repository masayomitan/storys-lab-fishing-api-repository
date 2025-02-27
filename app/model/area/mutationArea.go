package model

import (
	"context"
	"fmt"
	"time"
	"gorm.io/gorm"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/utils"
)

func (a AreaSQL) CreateByAdmin(ctx context.Context, requestParam domain.Area) (domain.Area, error) {
	err := a.db.Transaction(ctx, func(*gorm.DB) error {
        // Set timestamps for creation
        utils.SetCreateTimestamps(&requestParam)
        fmt.Println(requestParam)

        // Insert the area record and foreign table
		if err := a.db.Store(ctx, a.tableName, &requestParam); err != nil {
			return errors.Wrap(err, "error creating area")
		}

        return nil
    })
    if err != nil {
        return domain.Area{}, err
    }
    return requestParam, nil
}

func (a AreaSQL) DeleteByAdmin(ctx context.Context, id int) error {
	var existingDeletedAt *time.Time

	// 事前に `deleted_at` の値をチェック
	err := a.db.DB.Table(a.tableName).
		Select("deleted_at").
		Where("id = ?", id).
		First(&existingDeletedAt).Error

	// IDが存在しない場合
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("no record found with id %d", id)
	}

	// すでに `deleted_at` が設定されている場合
	if existingDeletedAt != nil {
		return fmt.Errorf("record with id %d is already deleted", id)
	}

	// 🔹 `deleted_at` を更新 (削除)
	result := a.db.DB.Table(a.tableName).
		Where("id = ?", id).
		Update("deleted_at", gorm.Expr("CURRENT_TIMESTAMP"))

	// SQLエラーが発生した場合
	if result.Error != nil {
		return errors.Wrap(result.Error, "error deleting area")
	}

	return nil
}



func (ga *GormAdapter) Store(ctx context.Context, table string, entity interface{}) error {
    return ga.DB.Table(table).Create(entity).Error
}

func (ga *GormAdapter) Update(ctx context.Context, table string, entity interface{}, id int) error {
    return ga.DB.Table(table).
		Where("id = ?", id).
		Updates(entity).Error
}

func (ga *GormAdapter) Delete(ctx context.Context, table string, id int) error {
	return ga.DB.Table(table).
		Where("id = ?", id).
		Update("deleted_at", gorm.Expr("CURRENT_TIMESTAMP")).Error
}

func (ga *GormAdapter) Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
    return ga.DB.WithContext(ctx).Transaction(fn)
}
