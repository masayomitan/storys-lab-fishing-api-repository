package domain

func (Dish) TableName() string {
    return "dishes"
}

type Dish struct {
    ID string `gorm:"primaryKey" json:"id"`
    Name string  `json:"name"`
	Description string `json:"description"`
	Ingredients string `json:"ingredients"`
	Kind string `json:"kind"`
	Level string `json:"level"`
	DishImages []DishImage `gorm:"foreignKey:DishId"`
}

