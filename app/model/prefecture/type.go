package model

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	DB *gorm.DB
}

type PrefectureSQL struct {
	db             *GormAdapter
	tableName string
}

func NewPrefectureSQL(db *gorm.DB) PrefectureSQL {
	return PrefectureSQL{
		db:             &GormAdapter{DB: db},
		tableName: "prefectures",
	}
}

func NewPrefectureOneSQL(db *gorm.DB) PrefectureSQL {
	return PrefectureSQL{
		db:             &GormAdapter{DB: db},
		tableName: "prefectures",
	}
}
