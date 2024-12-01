package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	EventRepository interface {
		// Create(context.Context, domain.Area) (domain.Area, error)
		FindAll(context.Context) ([]domain.Event, error)
		FindOne(context.Context, int) (domain.Event, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}
)
