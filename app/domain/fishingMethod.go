package domain

func (FishingMethod) TableName() string {
	return "fishing_methods"
}

type FishingMethod struct {
    ID string `gorm:"primaryKey" json:"id"`
    Name string `json:"name"`
	Description string `json:"description"`
	DifficultyLevel string `json:"difficulty_level"`
	IsTraditional bool `json:"is_traditional" gorm:"not null,DEFAULT false"`
}
