package domain

func (DishImage) TableName() string {
    return "dish_images"
}

type DishImage struct {
    ID   string `gorm:"primaryKey" json:"id"`
    DishId string  `json:"dish_id"`
	ImageUrl string `json:"image_url"`
	Sort string `json:"sort"`
	IsMain string `json:"is_main"`
}

