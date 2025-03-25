package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	ToolRepository interface {
		// Create(context.Context, domain.Tool) (domain.Tool, error)
		FindOne(context.Context, int) (domain.Tool, error)
		Find(context.Context) ([]domain.Tool, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}

	ToolAdminRepository interface {
		FindByAdmin(context.Context) ([]domain.Tool, error)
		FindOneByAdmin(context.Context, int) (domain.Tool, error)
		CreateByAdmin(context.Context, domain.Tool) (domain.Tool, error)
		UpdateByAdmin(context.Context, domain.Tool, int) (domain.Tool, error)
		DeleteByAdmin(context.Context, int) error
		// WithTransaction(context.Context, func(context.Context) error) error
	}
)
