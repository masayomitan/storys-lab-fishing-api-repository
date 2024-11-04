package model

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	DB *gorm.DB
}

type ToolSQL struct {
	db             *GormAdapter
	tableName string
}

func NewToolSQL(db *gorm.DB) ToolSQL {
	return ToolSQL{
		db:             &GormAdapter{DB: db},
		tableName: "tools",
	}
}
