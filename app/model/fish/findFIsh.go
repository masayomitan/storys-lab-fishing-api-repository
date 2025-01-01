package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"storys-lab-fishing-api/app/domain"
	"github.com/pkg/errors"
)


func (a FishSQL) FindOne(ctx context.Context, id int) (domain.Fish, error) {
	var json = domain.Fish{}
	fmt.Println(id)

	if err := a.db.FindOne(ctx, a.tableName, id, &json); err != nil {
		return domain.Fish{}, errors.Wrap(err, "error listing fishes")
	}

	var fish = domain.NewFish(
		json.ID,
		json.Name,
		json.FamilyName,
		json.ScientificName,
		json.FishCategoryId,
		json.Description,
		json.Length,
		json.Weight,
		json.Habitat,
		json.DepthRange,
		json.WaterTemperatureRange,
		json.ConservationStatus,

		convertFishCategory(json.FishCategory),
		convertFishingMethods(json.FishingMethods),
		convertDishes(json.Dishes),
		json.FishImages,
	)
	fmt.Println("fishStruct")
	fmt.Println(fish)

	return fish, nil
}

func (ga *GormAdapter) FindOne(ctx context.Context, table string, id int, result interface{}) error {
	return ga.DB.Table(table).Where("id = ?", id).
		Preload("FishImages").
		Preload("FishCategory").
		Preload("FishingMethods", func(*gorm.DB) *gorm.DB {
			return ga.DB.Select("fishing_methods.*, fishing_methods_fishes.is_traditional").
				Joins("inner join fishing_methods_fishes on fishing_methods_fishes.fishing_method_id = fishing_methods.id").
				Where("fishing_methods_fishes.fish_id = ?", id)
		}).
		Preload("Dishes.DishImages").
		First(result).Error
}

func (a FishSQL) FindAll(ctx context.Context) ([]domain.Fish, error) {
	var fishJSON = make([]domain.Fish, 0)

	if err := a.db.FindAll(ctx, a.tableName, domain.Fish{}, &fishJSON); err != nil {
		return []domain.Fish{}, errors.Wrap(err, "error listing fishes")
	}
	var fishes = make([]domain.Fish, 0)

	for _, fishJSON := range fishJSON {
		var fish = domain.NewFish(
			fishJSON.ID,
			fishJSON.Name,
			fishJSON.FamilyName,
			fishJSON.ScientificName,
			fishJSON.FishCategoryId,
			fishJSON.Description,
			fishJSON.Length,
			fishJSON.Weight,
			fishJSON.Habitat,
			fishJSON.DepthRange,
			fishJSON.WaterTemperatureRange,
			fishJSON.ConservationStatus,
			convertFishCategory(fishJSON.FishCategory),
			convertFishingMethods(fishJSON.FishingMethods),
			convertDishes(fishJSON.Dishes),
			fishJSON.FishImages,
		)
		fishes = append(fishes, fish)
	}
	return fishes, nil
}

func (a FishSQL) FindOneByAdmin(ctx context.Context, id int) (domain.Fish, error) {
	var json = domain.Fish{}
	fmt.Println(id)

	if err := a.db.FindOneByAdmin(ctx, a.tableName, id, &json); err != nil {
		return domain.Fish{}, errors.Wrap(err, "error listing fishes")
	}

	var fish = domain.NewFish(
		json.ID,
		json.Name,
		json.FamilyName,
		json.ScientificName,
		json.FishCategoryId,
		json.Description,
		json.Length,
		json.Weight,
		json.Habitat,
		json.DepthRange,
		json.WaterTemperatureRange,
		json.ConservationStatus,

		convertFishCategory(json.FishCategory),
		convertFishingMethods(json.FishingMethods),
		convertDishes(json.Dishes),
		json.FishImages,
	)
	fmt.Println("fishStruct")
	fmt.Println(fish)

	return fish, nil
}

func (ga *GormAdapter) FindOneByAdmin(ctx context.Context, table string, id int, result interface{}) error {
	return ga.DB.Table(table).Where("id = ?", id).
		Preload("FishImages").
		Preload("FishCategory").
		Preload("FishingMethods", func(*gorm.DB) *gorm.DB {
			return ga.DB.Select("fishing_methods.*, fishing_methods_fishes.is_traditional").
				Joins("inner join fishing_methods_fishes on fishing_methods_fishes.fishing_method_id = fishing_methods.id").
				Where("fishing_methods_fishes.fish_id = ?", id)
		}).
		Preload("Dishes.DishImages").
		First(result).Error
}


func (a FishSQL) FindAllByAdmin(ctx context.Context) ([]domain.Fish, error) {
	var fishJSON = make([]domain.Fish, 0)

	if err := a.db.FindAll(ctx, a.tableName, domain.Fish{}, &fishJSON); err != nil {
		return []domain.Fish{}, errors.Wrap(err, "error listing fishes")
	}

	var fishes = make([]domain.Fish, 0)

	for _, fishJSON := range fishJSON {
		var fish = domain.NewFish(
			fishJSON.ID,
			fishJSON.Name,
			fishJSON.FamilyName,
			fishJSON.ScientificName,
			fishJSON.FishCategoryId,
			fishJSON.Description,
			fishJSON.Length,
			fishJSON.Weight,
			fishJSON.Habitat,
			fishJSON.DepthRange,
			fishJSON.WaterTemperatureRange,
			fishJSON.ConservationStatus,
			convertFishCategory(fishJSON.FishCategory),
			convertFishingMethods(fishJSON.FishingMethods),
			convertDishes(fishJSON.Dishes),
			fishJSON.FishImages,
		)

		fishes = append(fishes, fish)
	}

	return fishes, nil
}

func (ga *GormAdapter) FindAll(ctx context.Context, table string, query interface{}, result interface{}) error {
    return ga.DB.Table(table).Where(query).
		Preload("FishImages").
		Order("id asc").
		Find(result).Error
}
