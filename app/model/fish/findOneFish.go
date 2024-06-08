package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"storys-lab-fishing-api/app/domain"
)


func (a FishSQL) FindOne(ctx context.Context, id string) (domain.Fish, error) {
	var fishJSON = domain.Fish{}
	fmt.Println(id)

	if err := a.db.FindOneFish(ctx, a.collectionName, id, &fishJSON); err != nil {
		return domain.Fish{}, errors.Wrap(err, "error listing fishes")
	}

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
	fmt.Println("fishStruct")
	fmt.Println(fish)

	return fish, nil
}

func (ga *GormAdapter) FindOneFish(ctx context.Context, table string, id string, result interface{}) error {
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

func convertFishCategory(category domain.FishCategory) domain.FishCategory {
	return domain.FishCategory{
		ID:   category.ID,
		Name: category.Name,
		Description: category.Description,
	}
}

func convertFishingMethods(methods []domain.FishingMethod) []domain.FishingMethod {
	var result []domain.FishingMethod
	for _, m := range methods {
		result = append(result, domain.FishingMethod{
			ID:   m.ID,
			Name: m.Name,
			Description: m.Description,
			DifficultyLevel: m.DifficultyLevel,
			IsTraditional: m.IsTraditional,
		})
	}
	return result
}

func convertDishes(dishes []domain.Dish) []domain.Dish {
	var result []domain.Dish
	for _, d := range dishes {
		var dishImages []domain.DishImage
		for _, img := range d.DishImages {
			dishImages = append(dishImages, domain.DishImage{
				ID:     	img.ID,
				DishId: 	img.DishId,
				ImageUrl:   img.ImageUrl,
			})
		}

		result = append(result, domain.Dish{
			ID:          d.ID,
			Name:        d.Name,
			Description: d.Description,
			Ingredients: d.Ingredients,
			Kind:        d.Kind,
			Level:       d.Level,
			DishImages:  dishImages,
		})
	}
	return result
}

