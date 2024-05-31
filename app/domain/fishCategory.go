package domain

type Tabler interface {
  TableName() string
}

func (FishCategory) TableName() string {
    return "fish_categories"
}

type FishCategory struct {
    ID   string `gorm:"primaryKey"`
    Name string  `json:"name"`
	  Description string `json:"description"`
}

