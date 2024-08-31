package model

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	DB *gorm.DB
}

type PrefSQL struct {
	db             *GormAdapter
	collectionName string
}

func NewPrefSQL(db *gorm.DB) PrefSQL {
	return PrefSQL{
		db:             &GormAdapter{DB: db},
		collectionName: "prefectures",
	}
}

func NewPrefOneSQL(db *gorm.DB) PrefSQL {
	return PrefSQL{
		db:             &GormAdapter{DB: db},
		collectionName: "prefectures",
	}
}
