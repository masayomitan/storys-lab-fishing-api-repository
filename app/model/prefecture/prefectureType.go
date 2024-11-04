package model

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	DB *gorm.DB
}

type PrefSQL struct {
	db             *GormAdapter
	tableName string
}

func NewPrefSQL(db *gorm.DB) PrefSQL {
	return PrefSQL{
		db:             &GormAdapter{DB: db},
		tableName: "prefectures",
	}
}

func NewPrefOneSQL(db *gorm.DB) PrefSQL {
	return PrefSQL{
		db:             &GormAdapter{DB: db},
		tableName: "prefectures",
	}
}
