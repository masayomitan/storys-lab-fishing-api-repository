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

type FishID string

func (f FishID) String() string {
	return string(f)
}

type (
	FishRepository interface {
		Create(context.Context, Fish) (Fish, error)
		FindAll(context.Context) ([]Fish, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}

	Fish struct {
		id   FishID
		name string
		familyName string
		scientificName string
		fishCategoryId int
		description string
	}
)

func NewFish(
	ID FishID,
	name string,
	familyName string,
	scientificName string,
	fishCategoryId int,
	description string,
) Fish {
	return Fish{
		id:   ID,
		name: name,
		familyName: familyName,
		scientificName: scientificName,
		fishCategoryId: fishCategoryId,
		description: description,
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
