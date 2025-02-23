package repository

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

type (
	PrefectureRepository interface {
		Find(context.Context) ([]domain.Prefecture, error)
		FindOne(context.Context, int) (domain.Prefecture, error)
	}

	PrefectureAdminRepository interface {
		FindByAdmin(context.Context) ([]domain.Prefecture, error)
		FindOneByAdmin(context.Context, int) (domain.Prefecture, error)
	}
)
