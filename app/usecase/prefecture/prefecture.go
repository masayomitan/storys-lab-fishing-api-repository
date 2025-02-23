package usecase

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

func (t prefectureInteractor) FindExecute(ctx context.Context) ([]domain.Prefecture, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	prefectures, err := t.repo.Find(ctx)
	if err != nil {
		return t.presenter.Presents([]domain.Prefecture{}), err
	}

	return t.presenter.Presents(prefectures), nil
}

func (t prefectureInteractor) FindOneExecute(ctx context.Context, id int) (domain.Prefecture, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	prefecture, err := t.repo.FindOne(ctx, id)
	if err != nil {
		return t.presenter.Present(domain.Prefecture{}), err
	}

	return t.presenter.Present(prefecture), nil
}

func (t prefectureAdminInteractor) FindExecuteByAdmin(ctx context.Context) ([]domain.Prefecture, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	prefectures, err := t.repo.FindByAdmin(ctx)
	if err != nil {
		return t.presenter.Presents([]domain.Prefecture{}), err
	}

	return t.presenter.Presents(prefectures), nil
}

func (t prefectureAdminInteractor) FindOneExecuteByAdmin(ctx context.Context, id int) (domain.Prefecture, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	prefecture, err := t.repo.FindOneByAdmin(ctx, id)
	if err != nil {
		return t.presenter.Present(domain.Prefecture{}), err
	}

	return t.presenter.Present(prefecture), nil
}
