package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	PrefRepository interface {
		// Create(context.Context, domain.Fish) (domain.Fish, error)
		// FindAll(context.Context) ([]domain.Fish, error)
		FindOne(context.Context, string) (domain.Pref, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}
)
