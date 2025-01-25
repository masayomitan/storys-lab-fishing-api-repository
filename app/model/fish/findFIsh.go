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

	if err := a.db.FindOneORM(ctx, a.tableName, id, &json); err != nil {
		return domain.Fish{}, errors.Wrap(err, "error listing fishes")
	}

	var fish = domain.NewFish(
		json.ID,
		json.Name,
		json.ScientificName,
		json.EnglishName,
		json.FishCategoryID,
		json.Description,
		json.Length,
		json.Weight,
		json.Habitat,
		json.DepthRangeMax,
		json.DepthRangeMin,
		json.WaterTemperatureRangeMax,
		json.WaterTemperatureRangeMin,
		json.ConservationStatus,

		convertFishCategory(json.FishCategory),
		convertFishingMethods(json.FishingMethods),
		convertDishes(json.Dishes),
		convertImages(json.Images),
	)
	fmt.Println("fishStruct")
	fmt.Println(fish)

	return fish, nil
}

func (a FishSQL) FindAll(ctx context.Context) ([]domain.Fish, error) {
	var fishJSON = make([]domain.Fish, 0)

	if err := a.db.FindAllORM(ctx, a.tableName, domain.Fish{}, &fishJSON); err != nil {
		return []domain.Fish{}, errors.Wrap(err, "error listing fishes")
	}
	var fishes = make([]domain.Fish, 0)

	for _, json := range fishJSON {
		var fish = domain.NewFish(
			json.ID,
			json.Name,
			json.EnglishName,
			json.ScientificName,
			json.FishCategoryID,
			json.Description,
			json.Length,
			json.Weight,
			json.Habitat,
			json.DepthRangeMax,
			json.DepthRangeMin,
			json.WaterTemperatureRangeMax,
			json.WaterTemperatureRangeMin,
			json.ConservationStatus,
			convertFishCategory(json.FishCategory),
			convertFishingMethods(json.FishingMethods),
			convertDishes(json.Dishes),
			convertImages(json.Images),
		)
		fishes = append(fishes, fish)
	}
	return fishes, nil
}

func (a FishSQL) FindOneByAdmin(ctx context.Context, id int) (domain.Fish, error) {
	var json = domain.Fish{}
	fmt.Println(id)

	if err := a.db.FindOneORM(ctx, a.tableName, id, &json); err != nil {
		return domain.Fish{}, errors.Wrap(err, "error listing fishes")
	}

	var fish = domain.NewFish(
		json.ID,
		json.Name,
		json.EnglishName,
		json.ScientificName,
		json.FishCategoryID,
		json.Description,
		json.Length,
		json.Weight,
		json.Habitat,
		json.DepthRangeMax,
		json.DepthRangeMin,
		json.WaterTemperatureRangeMax,
		json.WaterTemperatureRangeMin,
		json.ConservationStatus,

		convertFishCategory(json.FishCategory),
		convertFishingMethods(json.FishingMethods),
		convertDishes(json.Dishes),
		convertImages(json.Images),
	)
	fmt.Println("fishStruct")
	fmt.Println(fish)

	return fish, nil
}

func (a FishSQL) FindAllByAdmin(ctx context.Context) ([]domain.Fish, error) {
	var fishJSON = make([]domain.Fish, 0)

	if err := a.db.FindAllORM(ctx, a.tableName, domain.Fish{}, &fishJSON); err != nil {
		return []domain.Fish{}, errors.Wrap(err, "error listing fishes")
	}

	var fishes = make([]domain.Fish, 0)

	for _, json := range fishJSON {
		var fish = domain.NewFish(
			json.ID,
			json.Name,
			json.EnglishName,
			json.ScientificName,
			json.FishCategoryID,
			json.Description,
			json.Length,
			json.Weight,
			json.Habitat,
			json.DepthRangeMax,
			json.DepthRangeMin,
			json.WaterTemperatureRangeMax,
			json.WaterTemperatureRangeMin,
			json.ConservationStatus,
			convertFishCategory(json.FishCategory),
			convertFishingMethods(json.FishingMethods),
			convertDishes(json.Dishes),
			convertImages(json.Images),
		)

		fishes = append(fishes, fish)
	}

	return fishes, nil
}

func (ga *GormAdapter) FindOneORM(ctx context.Context, table string, id int, result interface{}) error {
	return ga.DB.Table(table).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
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

func (ga *GormAdapter) FindAllORM(ctx context.Context, table string, query interface{}, result interface{}) error {
    return ga.DB.Table(table).Where(query).
		// Preload("FishImages").
		Where("deleted_at IS NULL").
		Order("id asc").
		Find(result).Error
}
