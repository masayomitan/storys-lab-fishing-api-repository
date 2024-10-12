package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	ToolCategoryRepository interface {
		// Create(context.Context, domain.Fish) (domain.Fish, error)
		// FindOne(context.Context, int) (domain.ToolCategory, error)
		FindAll(context.Context) ([]domain.ToolCategory, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}
)
