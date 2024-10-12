package domain

func (Tool) TableName() string {
    return "tools"
}

type Tool struct {
    ID int `gorm:"primaryKey" json:"id"`
    Name string  `json:"name"`
	Description string `json:"description"`
	ToolCategoryId int `json:"tool_category_id"`
	MaterialID int `json:"material_id"`
	Size string `json:"size"`
	Weight  float64 `json:"weight"`
	Durability string `json:"durability"`
	ToolUsage string `json:"tool_usage"`
	Price int `json:"price"`
}
