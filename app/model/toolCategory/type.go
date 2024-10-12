package model

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	DB *gorm.DB
}

type ToolCategorySQL struct {
	db             *GormAdapter
	collectionName string
}

func NewToolCategorySQL(db *gorm.DB) ToolCategorySQL {
	return ToolCategorySQL{
		db:             &GormAdapter{DB: db},
		collectionName: "tool_categories",
	}
}
