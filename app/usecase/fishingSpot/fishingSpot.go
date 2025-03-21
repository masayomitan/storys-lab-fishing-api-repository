package usecase

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

func (a FishingSpotInteractor) FindOneExecute(ctx context.Context, id int) (domain.FishingSpot, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	fishingSpot, err := a.repo.FindOne(ctx, id)
	if err != nil {
		return a.presenter.Present(domain.FishingSpot{}), err
	}

	return a.presenter.Present(fishingSpot), nil
}

func (a FishingSpotInteractor) FindByAreaIdExecute(ctx context.Context, id int) ([]domain.FishingSpot, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	fishingSpot, err := a.repo.FindByAreaId(ctx, id)
	if err != nil {
		return a.presenter.Presents([]domain.FishingSpot{}), err
	}

	return a.presenter.Presents(fishingSpot), nil
}


func (t FishingSpotInteractor) FindAllExecute(ctx context.Context) ([]domain.FishingSpot, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fishingSpots, err := t.repo.Find(ctx)
	if err != nil {
		return t.presenter.Presents([]domain.FishingSpot{}), err
	}

	return t.presenter.Presents(fishingSpots), nil
}

func (t FishingSpotAdminInteractor) FindExecuteByAdmin(ctx context.Context) ([]domain.FishingSpot, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fishingSpots, err := t.repo.FindByAdmin(ctx)
	if err != nil {
		return t.presenter.Presents([]domain.FishingSpot{}), err
	}

	return t.presenter.Presents(fishingSpots), nil
}

func (t FishingSpotAdminInteractor) FindOneExecuteByAdmin(ctx context.Context, id int) (domain.FishingSpot, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fishingSpot, err := t.repo.FindOneByAdmin(ctx, id)
	if err != nil {
		return t.presenter.Present(domain.FishingSpot{}), err
	}

	return t.presenter.Present(fishingSpot), nil
}

func (t FishingSpotAdminInteractor) CreateExecuteByAdmin(ctx context.Context, requestParam domain.FishingSpot) (domain.FishingSpot, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fishingSpot, err := t.repo.CreateByAdmin(ctx, requestParam)
	if err != nil {
		return t.presenter.Present(domain.FishingSpot{}), err
	}

	return t.presenter.Present(fishingSpot), nil
}

func (t FishingSpotAdminInteractor) UpdateExecuteByAdmin(ctx context.Context, requestParam domain.FishingSpot, id int) (domain.FishingSpot, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fishingSpot, err := t.repo.UpdateByAdmin(ctx, requestParam, id)
	if err != nil {
		return t.presenter.Present(domain.FishingSpot{}), err
	}

	return t.presenter.Present(fishingSpot), nil
}

func (t FishingSpotAdminInteractor) DeleteExecuteByAdmin(ctx context.Context, id int) error {
    // Repositoryを呼び出して削除処理を実行
    return t.repo.DeleteByAdmin(ctx, id)
}

