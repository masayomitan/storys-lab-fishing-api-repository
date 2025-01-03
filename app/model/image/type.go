package model

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	DB *gorm.DB
}

type ImageSQL struct {
	db			*GormAdapter
	tableName 	string
}

func NewImageSQL(db *gorm.DB) ImageSQL {
	return ImageSQL{
		db:			&GormAdapter{DB: db},
		tableName: 	"images",
	}
}
