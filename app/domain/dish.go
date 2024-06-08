package domain

func (Dish) TableName() string {
    return "dishes"
}

type Dish struct {
    ID   string `gorm:"primaryKey"`
    Name string  `json:"name"`
	Description string `json:"description"`
	Ingredients string `json:"ingredients"`
	Kind string `json:"kind"`
	Level string `json:"level"`
}

