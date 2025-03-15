package model

import (
	"context"
	"gorm.io/gorm"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/utils"
)

func (a FishingSpotSQL) CreateByAdmin(ctx context.Context, requestParam domain.FishingSpot) (domain.FishingSpot, error) {
	err := a.db.Transaction(ctx, func(*gorm.DB) error {
        utils.SetCreateTimestamps(&requestParam)

        // Insert the area record and foreign table
		if err := a.db.Store(a.tableName, &requestParam); err != nil {
			return errors.Wrap(err, "error creating area")
		}

        return nil
    })
    if err != nil {
        return domain.FishingSpot{}, err
    }
    return requestParam, nil
}

func (a FishingSpotSQL) UpdateByAdmin(ctx context.Context, requestParam domain.FishingSpot, id int) (domain.FishingSpot, error) {
	err := a.db.Transaction(ctx, func(tx *gorm.DB) error {
		utils.SetCreateTimestamps(&requestParam)

		// fishing_spots テーブルの更新（Tags, Images フィールドは除外）
		if err := tx.Table(a.tableName).
			Where("id = ?", id).
			Omit("Tags", "Images").
			Updates(requestParam).Error; err != nil {
			return errors.Wrap(err, "error updating fishing spot")
		}

		// fishing_spots_to_tags の古いレコードを削除
		if err := tx.Table("fishing_spots_to_tags").
			Where("fishing_spot_id = ?", id).
			Delete(nil).Error; err != nil {
			return errors.Wrap(err, "error deleting old tag associations")
		}

		// fishing_spots_to_images の古いレコードを削除
		if err := tx.Table("fishing_spots_to_images").
			Where("fishing_spot_id = ?", id).
			Delete(nil).Error; err != nil {
			return errors.Wrap(err, "error deleting old image associations")
		}

		// requestParam.Tags に基づいて新しいレコードを登録
		for _, tag := range requestParam.Tags {
			joinRecord := map[string]interface{}{
				"fishing_spot_id": id,
				"tag_id":          tag.ID,
			}
			if err := tx.Table("fishing_spots_to_tags").
				Create(joinRecord).Error; err != nil {
				return errors.Wrap(err, "error inserting new tag association")
			}
		}

		// requestParam.Images に基づいて新しいレコードを登録
		for _, image := range requestParam.Images {
			joinRecord := map[string]interface{}{
				"fishing_spot_id": id,
				"image_id":        image.ID,
			}
			if err := tx.Table("fishing_spots_to_images").
				Create(joinRecord).Error; err != nil {
				return errors.Wrap(err, "error inserting new image association")
			}
		}

		return nil
	})

	if err != nil {
		return domain.FishingSpot{}, err
	}

	return requestParam, nil
}

func (a FishingSpotSQL) DeleteByAdmin(ctx context.Context, id int) error {

	if err := a.db.Delete(a.tableName, id); err != nil {
		return errors.Wrap(err, "error deleting area")
	}
	return nil
}

func (ga *GormAdapter) Store(table string, entity interface{}) error {
    return ga.DB.Table(table).Create(entity).Error
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
