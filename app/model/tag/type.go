package model

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	DB *gorm.DB
}

type TagSQL struct {
	db             *GormAdapter
	tableName string
}

func NewTagSQL(db *gorm.DB) TagSQL {
	return TagSQL{
		db:             &GormAdapter{DB: db},
		tableName: "tags",
	}
}

func NewFishingSpotOneSQL(db *gorm.DB) TagSQL {
	return TagSQL{
		db:             &GormAdapter{DB: db},
		tableName: "tags",
	}
}

