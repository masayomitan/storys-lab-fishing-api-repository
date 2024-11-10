package model

import (
	"gorm.io/gorm"
)

type GormAdapter struct {
	DB *gorm.DB
}

type ArticleSQL struct {
	db		  	*GormAdapter
	tableName 	string
}

func NewArticleSQL(db *gorm.DB) ArticleSQL {
	return ArticleSQL{
		db:			&GormAdapter{DB: db},
		tableName: 	"articles",
	}
}

func NewArticleOneSQL(db *gorm.DB) ArticleSQL {
	return ArticleSQL{
		db:         &GormAdapter{DB: db},
		tableName: 	"articles",
	}
}
