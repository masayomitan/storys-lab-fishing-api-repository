package domain

func (FishImage) TableName() string {
    return "fish_images"
}

type FishImage struct {
    ID string `gorm:"primaryKey" json:"id"`
    FishId string  `json:"fish_id"`
	ImageUrl string `json:"image_url"`
	Sort string `json:"sort"`
	IsMain string `json:"is_main"`
}

