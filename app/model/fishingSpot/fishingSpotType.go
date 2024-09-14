package model

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	DB *gorm.DB
}

type FishingSpotSQL struct {
	db             *GormAdapter
	tableName string
}

func NewFishingSpotSQL(db *gorm.DB) FishingSpotSQL {
	return FishingSpotSQL{
		db:             &GormAdapter{DB: db},
		tableName: "fishing_spots",
	}
}

func NewFishingSpotOneSQL(db *gorm.DB) FishingSpotSQL {
	return FishingSpotSQL{
		db:             &GormAdapter{DB: db},
		tableName: "fishing_spots",
	}
}
