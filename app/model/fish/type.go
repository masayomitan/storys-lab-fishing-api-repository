package model

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	DB *gorm.DB
}

type FishSQL struct {
	db             *GormAdapter
	collectionName string
}

func NewFishSQL(db *gorm.DB) FishSQL {
	return FishSQL{
		db:             &GormAdapter{DB: db},
		collectionName: "fishes",
	}
}

func NewFishOneSQL(db *gorm.DB) FishSQL {
	return FishSQL{
		db:             &GormAdapter{DB: db},
		collectionName: "fishes",
	}
}
