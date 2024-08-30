package model

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	DB *gorm.DB
}

type AreaSQL struct {
	db             *GormAdapter
	collectionName string
}

func NewAreaSQL(db *gorm.DB) AreaSQL {
	return AreaSQL{
		db:             &GormAdapter{DB: db},
		collectionName: "areas",
	}
}

func NewAreaOneSQL(db *gorm.DB) AreaSQL {
	return AreaSQL{
		db:             &GormAdapter{DB: db},
		collectionName: "areas",
	}
}
