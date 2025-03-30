package model

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	DB *gorm.DB
}

type MaterialSQL struct {
	db             *GormAdapter
	tableName string
}

func NewMaterialSQL(db *gorm.DB) MaterialSQL {
	return MaterialSQL{
		db:             &GormAdapter{DB: db},
		tableName: "materials",
	}
}

func NewFishingSpotOneSQL(db *gorm.DB) MaterialSQL {
	return MaterialSQL{
		db:             &GormAdapter{DB: db},
		tableName: "materials",
	}
}

