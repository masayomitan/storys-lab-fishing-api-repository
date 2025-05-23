package domain
import "time"

func (ToolCategory) TableName() string {
    return "tool_categories"
}

func (Tool) TableName() string {
    return "tools"
}

type ToolCategory struct {
    ID              int         `gorm:"primaryKey" json:"id"`
    Name            string      `json:"name"`
	Description     string      `json:"description"`
    CreatedAt   	time.Time	`gorm:"created_at" json:"created_at"`
    UpdatedAt		time.Time  	`gorm:"updated_at" json:"updated_at"`
	DeletedAt		*time.Time  `gorm:"default:NULL"`

	Tools           []Tool      `gorm:"foreignKey:ToolCategoryID"`
}

type Tool struct {
    ID             int          `gorm:"primaryKey" json:"id"`
    Name           string       `json:"name"`
    Description    string       `json:"description"`
    ToolCategoryID int          `json:"tool_category_id"`
    MaterialID     int          `json:"material_id"`
    Size           float64      `json:"size"`
    Weight         float64      `json:"weight"`
    Price          int          `json:"price"`
    Maker          string       `json:"maker"`
    Recommend      int          `json:"recommend"`
    EasyFishing    int          `json:"easy_fishing"`
    CreatedAt      time.Time	`gorm:"created_at" json:"created_at"`
    UpdatedAt	   time.Time  	`gorm:"updated_at" json:"updated_at"`
	DeletedAt	   *time.Time   `gorm:"default:NULL"`

    Images         []Image      `gorm:"many2many:tools_to_images;" validate:"-"`
}

func NewToolCategory(
	ID          int,
	name        string,
	description string,
	tools       []Tool,

) ToolCategory {
	return ToolCategory{
		ID:             ID,
		Name:           name,
		Description:    description,
		Tools:          tools,
	}
}

func NewTool(
    ID              int,
    name            string,
    description     string,
    toolCategoryID  int,
    materialID      int,
    size            float64,
    weight          float64,
    price           int,
    maker           string,
    recommend       int,
    easyFishing     int,
    images          []Image,
) Tool {
    return Tool{
        ID:             ID,
        Name:           name,
        Description:    description,
        ToolCategoryID: toolCategoryID,
        MaterialID:     materialID,
        Size:           size,
        Weight:         weight,
        Price:          price,
        Maker:          maker,
        Recommend:      recommend,
        EasyFishing:    easyFishing,
        Images:         images,
    }
}
