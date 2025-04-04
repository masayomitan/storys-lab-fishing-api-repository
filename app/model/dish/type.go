package model

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	DB *gorm.DB
}

type DishSQL struct {
	db             *GormAdapter
	tableName 		string
}

func NewDishSQL(db *gorm.DB) DishSQL {
	return DishSQL{
		db:             &GormAdapter{DB: db},
		tableName: 		"dishes",
	}
}
