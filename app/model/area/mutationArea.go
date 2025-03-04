package model

import (
	"context"
	"fmt"
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
		if err := a.db.Store(a.tableName, &requestParam); err != nil {
			return errors.Wrap(err, "error creating area")
		}

        return nil
    })
    if err != nil {
        return domain.Area{}, err
    }
    return requestParam, nil
}

func (a AreaSQL) UpdateByAdmin(ctx context.Context, requestParam domain.Area, id int) (domain.Area, error) {
    err := a.db.Transaction(ctx, func(*gorm.DB) error {
        // Set timestamps for creation
        utils.SetCreateTimestamps(&requestParam)
        fmt.Println(requestParam)

		if err := a.db.Update(a.tableName, &requestParam, id); err != nil {
			return errors.Wrap(err, "error creating fish")
		}

        return nil
    })

    if err != nil {
        return domain.Area{}, err
    }

    return requestParam, nil
}

func (a AreaSQL) DeleteByAdmin(ctx context.Context, id int) error {

	if err := a.db.Delete(a.tableName, id); err != nil {
		return errors.Wrap(err, "error deleting area")
	}
	return nil
}


func (ga *GormAdapter) Store(table string, entity interface{}) error {
    return ga.DB.Table(table).Create(entity).Error
}

func (ga *GormAdapter) Update(table string, entity interface{}, id int) error {
    return ga.DB.Table(table).
		Where("id = ?", id).
		Updates(entity).Error
}

func (ga *GormAdapter) Delete(table string, id int) error {
	return ga.DB.Table(table).
		Where("id = ?", id).
		Update("updated_at", gorm.Expr("CURRENT_TIMESTAMP")).
		Update("deleted_at", gorm.Expr("CURRENT_TIMESTAMP")).Error
}

func (ga *GormAdapter) Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
    return ga.DB.WithContext(ctx).Transaction(fn)
}
