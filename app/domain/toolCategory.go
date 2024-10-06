package domain

func (ToolCategory) TableName() string {
    return "tool_categories"
}

type ToolCategory struct {
    ID int `gorm:"primaryKey" json:"id"`
    Name string  `json:"name"`
	Description string `json:"description"`
}

