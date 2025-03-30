package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (

	MaterialAdminRepository interface {
		FindByAdmin(context.Context) ([]domain.Material, error)
	}
)
