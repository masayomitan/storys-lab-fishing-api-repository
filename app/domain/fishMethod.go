package domain

func (FishingMethod) TableName() string {
	return "fishing_methods"
}

type FishingMethod struct {
    ID   string `gorm:"primaryKey"`
    Name string  `json:"name"`
	Description string `json:"description"`
	Difficulty_level string `json:"DifficultyLevel"`
}
