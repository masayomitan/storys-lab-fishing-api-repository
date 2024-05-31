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

func (FishStruct) TableName() string {
    return "fishes"
}

type FishStruct struct {
	ID                      string  `gorm:"primaryKey"`
	Name                    string  `json:"name"`
	FamilyName              string  `json:"family_name"`
	ScientificName          string  `json:"scientific_name"`
	FishCategoryId          int     `json:"fish_category"`
	Description             string  `json:"description"`
	Length                  float64 `json:"length"`
	Weight                  float64 `json:"weight"`
	Habitat                 string  `json:"habitat"`
	DepthRange             string  `json:"depth_range"`
	WaterTemperatureRange string  `json:"water_temperature_range"`
	ConservationStatus     string  `json:"conservation_status"`

	FishCategory            FishCategory `gorm:"foreignKey:FishCategoryId"`
	FishingMethods []FishingMethod `gorm:"many2many:fishing_methods_fishes;foreignKey:ID;joinForeignKey:FishID;References:ID;joinReferences:FishingMethodID"`
	Dishes []Dish `gorm:"many2many:fishes_dishes;foreignKey:ID;joinForeignKey:FishID;References:ID;joinReferences:DishID"`
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
		id                      string
		name                    string
		familyName              string
		scientificName          string
		fishCategoryId          int
		description             string
		length                  float64
		weight                  float64
		habitat                 string
		depthRange             string
		waterTemperatureRange string
		conservationStatus     string

		fishCategory            FishCategory
		fishingMethods          []FishingMethod
		dishes                  []Dish
	}
)

func NewFish(
	ID string,
	name string,
	familyName string,
	scientificName string,
	fishCategoryId int,
	description string,
	length float64,
	weight float64,
	habitat string,
	depthRange string,
	waterTemperatureRange string,
	conservationStatus string,

	fishCategory            FishCategory,
	fishingMethods          []FishingMethod,
	dishes                  []Dish,
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
		depthRange:              depthRange,
		waterTemperatureRange: 	 waterTemperatureRange,
		conservationStatus:      conservationStatus,

		fishCategory:            fishCategory,
		fishingMethods:          fishingMethods,
		dishes:                  dishes,
	}
}

func (f Fish) ID() string {
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

func (f Fish) Weight() float64 {
	return f.weight
}

func (f Fish) Habitat() string {
	return f.habitat
}

func (f Fish) DepthRange() string {
	return f.depthRange
}

func (f Fish) WaterTemperatureRange() string {
	return f.waterTemperatureRange
}

func (f Fish) ConservationStatus() string {
	return f.conservationStatus
}

func (f Fish) FishCategory() FishCategory {
	return f.fishCategory
}

func (f Fish) FishingMethods() []FishingMethod {
	return f.fishingMethods
}

func (f Fish) Dishes() []Dish {
	return f.dishes
}
