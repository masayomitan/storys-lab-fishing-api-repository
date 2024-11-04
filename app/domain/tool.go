package domain

func (Tool) TableName() string {
    return "tools"
}

type Tool struct {
    ID             int     `gorm:"primaryKey" json:"id"`
    Name           string  `json:"name"`
    Description    string  `json:"description"`
    ToolCategoryId int     `json:"tool_category_id"`
    MaterialID     int     `json:"material_id"`
    Size           string  `json:"size"`
    Weight         float64 `json:"weight"`
    Durability     string  `json:"durability"`
    ToolUsage      string  `json:"tool_usage"`
    Price          int     `json:"price"`
    Maker          string  `json:"maker"`
    Recommend      int     `json:"recommend"`
    EasyFishing    int     `json:"easy_fishing"`
}

func NewTool(
    ID int,
    name string,
    description string,
    toolCategoryId int,
    materialID int,
    size string,
    weight float64,
    durability string,
    toolUsage string,
    price int,
    maker string,
    recommend int,
    easyFishing int,
) Tool {
    return Tool{
        ID:             ID,
        Name:           name,
        Description:    description,
        ToolCategoryId: toolCategoryId,
        MaterialID:     materialID,
        Size:           size,
        Weight:         weight,
        Durability:     durability,
        ToolUsage:      toolUsage,
        Price:          price,
        Maker:          maker,
        Recommend:      recommend,
        EasyFishing:    easyFishing,
    }
}
