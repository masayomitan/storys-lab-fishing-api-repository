package presenter

import (
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/dish"
)

type findDishPresenter struct{}

func NewDishPresenter() usecase.DishPresenter {
	return findDishPresenter{}
}

func (a findDishPresenter) Present(dish domain.Dish) domain.Dish {
	return domain.Dish{
		ID: 			dish.ID,
		Name: 			dish.Name,
		Description: 	dish.Description,
		Ingredients: 	dish.Ingredients,
		Kind: 			dish.Kind,
		Level: 			dish.Level,
		CreatedAt:		dish.CreatedAt,
		UpdatedAt:		dish.UpdatedAt,
		Images:			dish.Images,
	}
}

func (a findDishPresenter) Presents(dishes []domain.Dish) []domain.Dish {
	var output = make([]domain.Dish, 0)
	for _, dish := range dishes {
		output = append(output, domain.Dish{
			ID:				dish.ID,
			Name:			dish.Name,
			Description: 	dish.Description,
			Ingredients: 	dish.Ingredients,
			Kind: 			dish.Kind,
			Level: 			dish.Level,
			CreatedAt:		dish.CreatedAt,
			UpdatedAt:		dish.UpdatedAt,
			Images:			dish.Images,
		})
	}
	return output
}
