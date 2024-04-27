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
		fishName string
	}
)

func NewFish(
	ID FishID,
	fishName string,
) Fish {
	return Fish{
		id:   ID,
		fishName: fishName,
	}
}

func (f Fish) ID() FishID {
	return f.id
}

func (f Fish) FishName() string {
	return f.fishName
}
