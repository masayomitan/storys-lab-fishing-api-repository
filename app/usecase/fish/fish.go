package usecase

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

func (t fishInteractor) FindOneExecute(ctx context.Context, id int) (domain.Fish, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fish, err := t.repo.FindOne(ctx, id)
	if err != nil {
		return t.presenter.PresentOne(domain.Fish{}), err
	}

	return t.presenter.PresentOne(fish), nil
}


func (t fishInteractor) FindAllExecute(ctx context.Context) ([]domain.Fish, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fishes, err := t.repo.FindAll(ctx)
	if err != nil {
		return t.presenter.PresentAll([]domain.Fish{}), err
	}

	return t.presenter.PresentAll(fishes), nil
}

func (t fishAdminInteractor) FindOneExecuteByAdmin(ctx context.Context, id int) (domain.Fish, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fish, err := t.repo.FindOneByAdmin(ctx, id)
	if err != nil {
		return t.presenter.PresentOne(domain.Fish{}), err
	}

	return t.presenter.PresentOne(fish), nil
}

func (t fishAdminInteractor) FindAllExecuteByAdmin(ctx context.Context) ([]domain.Fish, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fishes, err := t.repo.FindAllByAdmin(ctx)
	if err != nil {
		return t.presenter.PresentAll([]domain.Fish{}), err
	}

	return t.presenter.PresentAll(fishes), nil
}

func (t fishAdminInteractor) CreateExecuteByAdmin(ctx context.Context, requestParam domain.Fish) (domain.Fish, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fish, err := t.repo.CreateByAdmin(ctx, requestParam)
	if err != nil {
		return t.presenter.PresentOne(domain.Fish{}), err
	}

	return t.presenter.PresentOne(fish), nil
}

func (t fishAdminInteractor) UpdateExecuteByAdmin(ctx context.Context, requestParam domain.Fish, id int) (domain.Fish, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fish, err := t.repo.UpdateByAdmin(ctx, requestParam, id)
	if err != nil {
		return t.presenter.PresentOne(domain.Fish{}), err
	}

	return t.presenter.PresentOne(fish), nil
}

func (t fishAdminInteractor) DeleteExecuteByAdmin(ctx context.Context, id int) error {

    return t.repo.DeleteByAdmin(ctx, id)
}

func (t fishAdminInteractor) UpdateFishDishesExecute(ctx context.Context, fishID int, dishIDs []int) (domain.FishDishRelationResult, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fish, _ := t.repo.UpdateFishDishesByAdmin(ctx, fishID, dishIDs)
	return fish, nil
}

