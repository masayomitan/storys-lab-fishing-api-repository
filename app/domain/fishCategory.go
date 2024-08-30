package domain

func (FishCategory) TableName() string {
    return "fish_categories"
}

type FishCategory struct {
    ID   string `gorm:"primaryKey" json:"id"`
    Name string  `json:"name"`
	Description string `json:"description"`
}

