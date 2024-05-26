package domain

import (
	"context"
	"errors"
	// "time"
)

var (
	ErrFishNotFound = errors.New("Fish not found")

	ErrFishOriginNotFound = errors.New("Fish origin not found")

	ErrFishDestinationNotFound = errors.New("Fish destination not found")

	ErrInsufficientBalance = errors.New("origin Fish does not have sufficient balance")
)

type FishStruct struct {
	ID                      string  `json:"id"`
	Name                    string  `json:"fish_name"`
	FamilyName              string  `json:"family_name"`
	ScientificName          string  `json:"scientific_name"`
	FishCategoryId          int     `json:"fish_category"`
	Description             string  `json:"description"`
	Length                  float64 `json:"length"`
	Weight                  float64 `json:"weight"`
	Habitat                 string  `json:"habitat"`
	Depth_range             string  `json:"depth_range"`
	Water_temperature_range string  `json:"water_temperature_range"`
	Conservation_status     string  `json:"conservation_status"`
}

type FishID string

func (f FishID) String() string {
	return string(f)
}

type (
	FishRepository interface {
		Create(context.Context, Fish) (Fish, error)
		FindAll(context.Context) ([]Fish, error)
		FindOne(context.Context, string) (Fish, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}

	Fish struct {
		id                      FishID
		name                    string
		familyName              string
		scientificName          string
		fishCategoryId          int
		description             string
		length                  float64
		weight                  float64
		habitat                 string
		depth_range             string
		water_temperature_range string
		conservation_status     string
	}
)

func NewFish(
	ID FishID,
	name string,
	familyName string,
	scientificName string,
	fishCategoryId int,
	description string,
	length float64,
	weight float64,
	habitat string,
	depth_range string,
	water_temperature_range string,
	conservation_status string,
) Fish {
	return Fish{
		id:                      ID,
		name:                    name,
		familyName:              familyName,
		scientificName:          scientificName,
		fishCategoryId:          fishCategoryId,
		description:             description,
		length:                  length,
		weight:                  weight,
		habitat:                 habitat,
		depth_range:             depth_range,
		water_temperature_range: water_temperature_range,
		conservation_status:     conservation_status,
	}
}

func (f Fish) ID() FishID {
	return f.id
}

func (f Fish) Name() string {
	return f.name
}

func (f Fish) FamilyName() string {
	return f.familyName
}

func (f Fish) ScientificName() string {
	return f.scientificName
}

func (f Fish) FishCategoryId() int {
	return f.fishCategoryId
}

func (f Fish) Description() string {
	return f.description
}

func (f Fish) Length() float64 {
	return f.length
}
