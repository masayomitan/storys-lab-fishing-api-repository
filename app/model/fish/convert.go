package model

import (
	"storys-lab-fishing-api/app/domain"
)


func convertFishCategory(c domain.FishCategory) domain.FishCategory {
	return domain.FishCategory {
		ID: c.ID,
		Name: c.Name,
		Description: c.Description,
	}
}

func convertFishingMethods(methods []domain.FishingMethod) []domain.FishingMethod {
	var result []domain.FishingMethod
	for _, m := range methods {
		result = append(result, domain.FishingMethod{
			ID: m.ID,
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

func convertImages(methods []domain.Image) []domain.Image {
	var result []domain.Image
	for _, i := range methods {
		result = append(result, domain.Image{
			ID: 		i.ID,
			Name: 		i.Name,
			ImageUrl: 	i.ImageUrl,
			S3Url: 		i.S3Url,
		})
	}
	return result
}
