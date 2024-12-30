package model

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	DB *gorm.DB
}

type FishCategorySQL struct {
	db             *GormAdapter
	tableName string
}

func NewFishCategorySQL(db *gorm.DB) FishCategorySQL {
	return FishCategorySQL{
		db:             &GormAdapter{DB: db},
		tableName: "fish_categories",
	}
}

func NewFishCategoryOneSQL(db *gorm.DB) FishCategorySQL {
	return FishCategorySQL{
		db:             &GormAdapter{DB: db},
		tableName: "fish_categories",
	}
}

