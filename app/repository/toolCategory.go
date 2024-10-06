package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	ToolRepository interface {
		// Create(context.Context, domain.Fish) (domain.Fish, error)
		FindAll(context.Context) ([]domain.Tool, error)
		FindOne(context.Context, string) (domain.Tool, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}
)
