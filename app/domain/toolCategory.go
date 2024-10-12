package domain

func (ToolCategory) TableName() string {
    return "tool_categories"
}

type ToolCategory struct {
    ID int `gorm:"primaryKey" json:"id"`
    Name string  `json:"name"`
	Description string `json:"description"`
	Tools []Tool `gorm:"foreignKey:ToolCategoryId"`

}

func NewToolCategory(
	ID int,
	name string,
	description string,
	tools []Tool,

) ToolCategory {
	return ToolCategory{
		ID: ID,
		Name: name,
		Description: description,
		Tools: tools,
	}
}
