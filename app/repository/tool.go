package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	ToolRepository interface {
		// Create(context.Context, domain.Fish) (domain.Fish, error)
		FindOne(context.Context, int) (domain.Tool, error)
		FindAll(context.Context) ([]domain.Tool, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}
)
