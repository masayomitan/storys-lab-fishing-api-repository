package model

import (
	"context"
	// "time"
	"fmt"
	"gorm.io/gorm"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/utils"
)

func (a FishSQL) CreateByAdmin(ctx context.Context, requestParam domain.Fish) (domain.Fish, error) {
    err := a.db.Transaction(ctx, func(*gorm.DB) error {
        // Set timestamps for creation
        utils.SetCreateTimestamps(&requestParam)
        fmt.Println(requestParam)

        // Insert the Fish record and foreign table
		if err := a.db.Store(ctx, a.tableName, &requestParam); err != nil {
			return errors.Wrap(err, "error creating fish")
		}

        return nil
    })

    if err != nil {
        return domain.Fish{}, err
    }

    return requestParam, nil
}

func (a FishSQL) UpdateByAdmin(ctx context.Context, requestParam domain.Fish, id int) (domain.Fish, error) {
    err := a.db.Transaction(ctx, func(*gorm.DB) error {
        // Set timestamps for creation
        utils.SetCreateTimestamps(&requestParam)
        fmt.Println(requestParam)

        // Insert the Fish record and foreign table
		if err := a.db.Update(ctx, a.tableName, &requestParam, id); err != nil {
			return errors.Wrap(err, "error creating fish")
		}

        return nil
    })

    if err != nil {
        return domain.Fish{}, err
    }

    return requestParam, nil
}

func (a FishSQL) DeleteByAdmin(ctx context.Context, id int) error {

	if err := a.db.Delete(ctx, a.tableName, id); err != nil {
		return errors.Wrap(err, "error deleting fish")
	}
	return nil
}

func (a FishSQL) UpdateFishDishesByAdmin(ctx context.Context, fishID int, dishIDs []int) (domain.FishDishRelationResult, error) {
	var fish domain.Fish

	err := a.db.Transaction(ctx, func(tx *gorm.DB) error {

		if err := tx.First(&fish, fishID).Error; err != nil {
			return errors.Wrap(err, "fish not found")
		}

		if err := tx.Model(&fish).Association("Dishes").Clear(); err != nil {
			return errors.Wrap(err, "failed to clear existing dishes")
		}

		if len(dishIDs) > 0 {
			var dishes []domain.Dish
			if err := tx.Where("id IN ?", dishIDs).Find(&dishes).Error; err != nil {
				return errors.Wrap(err, "failed to load new dishes")
			}
			if err := tx.Model(&fish).Association("Dishes").Append(dishes); err != nil {
				return errors.Wrap(err, "failed to append new dishes")
			}
		}

		return nil
	})

	if err != nil {
		return domain.FishDishRelationResult{}, err
	}

	return domain.FishDishRelationResult{
		FishID:  fishID,
		DishIDs: dishIDs,
	}, nil
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
        Update("updated_at", gorm.Expr("CURRENT_TIMESTAMP")).
		Update("deleted_at", gorm.Expr("CURRENT_TIMESTAMP")).Error
}


func (ga *GormAdapter) Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
    return ga.DB.WithContext(ctx).Transaction(fn)
}
