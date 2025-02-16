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
