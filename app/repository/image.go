package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	ImageAdminRepository interface {
		// FindOneByAdmin(context.Context, int) (domain.Image, error)
		FindAllByAdmin(context.Context, map[string]interface{}) ([]domain.Image, error)
		UploadByAdmin(context.Context, domain.Image) (domain.Image, error)
		// WithTransaction(context.Context, func(context.Context) error) error
	}
)
