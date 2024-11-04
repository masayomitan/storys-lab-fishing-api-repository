package model

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	DB *gorm.DB
}

type AreaSQL struct {
	db             *GormAdapter
	tableName string
}

func NewAreaSQL(db *gorm.DB) AreaSQL {
	return AreaSQL{
		db:             &GormAdapter{DB: db},
		tableName: "areas",
	}
}

func NewAreaOneSQL(db *gorm.DB) AreaSQL {
	return AreaSQL{
		db:             &GormAdapter{DB: db},
		tableName: "areas",
	}
}
