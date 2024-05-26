package repository

import (
	"storys-lab-fishing-api/app/adapter/repository"
)

type FishJSON struct {
	ID string `json:"id"`
	Name string `json:"fish_name"`
	FamilyName string `json:"family_name"`
	ScientificName string `json:"scientific_name"`
	FishCategoryId int `json:"fish_category"`
	Description string `json:"description"`
}

type FishSQL struct {
	collectionName string
	db             repository.DBMethods
}

func NewFishSQL(db repository.DBMethods) FishSQL {
	return FishSQL{
		db:             db,
		collectionName: "fishes",
	}
}

func NewFishOneSQL(db repository.DBMethods) FishSQL {
	return FishSQL{
		db:             db,
		collectionName: "fishes",
	}
}
