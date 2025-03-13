package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (

	TagAdminRepository interface {
		FindByAdmin(context.Context) ([]domain.Tag, error)
	}
)
