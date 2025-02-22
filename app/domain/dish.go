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
	// TODO 変更
	// Images     []Image 		`gorm:"many2many:dishes_to_images;joinForeignKey:DishID;JoinReferences:ImageID" validate:"-"`
}

type DishImage struct {
    ID string `gorm:"primaryKey" json:"id"`
    DishId string  `json:"dish_id"`
	ImageUrl string `json:"image_url"`
	Sort string `json:"sort"`
	IsMain string `json:"is_main"`
}
