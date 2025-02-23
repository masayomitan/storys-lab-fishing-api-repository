package usecase

import (
	"context"
	"time"
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/repository"
)

type PrefectureUseCase interface {
	FindExecute(ctx context.Context) ([]domain.Prefecture, error)
	FindOneExecute(ctx context.Context, id int) (domain.Prefecture, error)
}

type PrefectureAdminUseCase interface {
	FindExecuteByAdmin(ctx context.Context) ([]domain.Prefecture, error)
	FindOneExecuteByAdmin(ctx context.Context, id int) (domain.Prefecture, error)
}

type PrefecturePresenter interface {
	Present(domain.Prefecture) domain.Prefecture
	Presents([]domain.Prefecture) []domain.Prefecture
}

type prefectureInteractor struct {
	repo       repository.PrefectureRepository
	presenter  PrefecturePresenter
	ctxTimeout time.Duration
}

type prefectureAdminInteractor struct {
	repo       repository.PrefectureAdminRepository
	presenter  PrefecturePresenter
	ctxTimeout time.Duration
}

func NewPrefectureInteractor(
	repo repository.PrefectureRepository,
	presenter PrefecturePresenter,
	t time.Duration,
) PrefectureUseCase {
	return &prefectureInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func NewPrefectureAdminInteractor(
	repo repository.PrefectureAdminRepository,
	presenter PrefecturePresenter,
	timeout time.Duration,
) PrefectureAdminUseCase {
	return &prefectureAdminInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}
