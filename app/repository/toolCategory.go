package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	ToolCategoryRepository interface {
		// Create(context.Context, domain.ToolCategory) (domain.ToolCategory, error)
		FindOne(context.Context, int) (domain.ToolCategory, error)
		Find(context.Context) ([]domain.ToolCategory, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}

	ToolCategoryAdminRepository interface {
		FindByAdmin(context.Context) ([]domain.ToolCategory, error)
		FindOneByAdmin(context.Context, int) (domain.ToolCategory, error)
		CreateByAdmin(context.Context, domain.ToolCategory) (domain.ToolCategory, error)
		UpdateByAdmin(context.Context, domain.ToolCategory, int) (domain.ToolCategory, error)
		DeleteByAdmin(context.Context, int) error
		// WithTransaction(context.Context, func(context.Context) error) error
	}
)
