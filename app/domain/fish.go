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
		// Create(context.Context, Transfer) (Transfer, error)
		FindAll(context.Context) ([]Fish, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}

	Fish struct {
		id   FishID
		name string
	}
)

func NewFish(
	ID FishID,
	name string,
) Fish {
	return Fish{
		id:   ID,
		name: name,
	}
}

func (f Fish) ID() FishID {
	return f.id
}

func (f Fish) Name() string {
	return f.name
}
