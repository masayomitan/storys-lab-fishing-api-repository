package repository

import (
	"storys-lab-fishing-api/app/adapter/repository"
)

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
