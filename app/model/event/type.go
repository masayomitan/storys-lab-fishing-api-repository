package model

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	DB *gorm.DB
}

type EventSQL struct {
	db             *GormAdapter
	tableName string
}

func NewEventSQL(db *gorm.DB) EventSQL {
	return EventSQL{
		db:             &GormAdapter{DB: db},
		tableName: "events",
	}
}

func NewEventOneSQL(db *gorm.DB) EventSQL {
	return EventSQL{
		db:             &GormAdapter{DB: db},
		tableName: "events",
	}
}
