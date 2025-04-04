package usecase

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

func (t dishInteractor) FindOneExecute(ctx context.Context, id int) (domain.Dish, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	dish, err := t.repo.FindOne(ctx, id)
	if err != nil {
		return t.presenter.Present(domain.Dish{}), err
	}

	return t.presenter.Present(dish), nil
}

func (t dishInteractor) FindExecute(ctx context.Context) ([]domain.Dish, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	dish, err := t.repo.Find(ctx)
	if err != nil {
		return t.presenter.Presents([]domain.Dish{}), err
	}

	return t.presenter.Presents(dish), nil
}

func (t dishAdminInteractor) FindOneExecuteByAdmin(ctx context.Context, id int) (domain.Dish, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	dish, err := t.repo.FindOneByAdmin(ctx, id)
	if err != nil {
		return t.presenter.Present(domain.Dish{}), err
	}

	return t.presenter.Present(dish), nil
}

func (t dishAdminInteractor) FindExecuteByAdmin(ctx context.Context) ([]domain.Dish, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	dishes, err := t.repo.FindByAdmin(ctx)
	if err != nil {
		return t.presenter.Presents([]domain.Dish{}), err
	}

	return t.presenter.Presents(dishes), nil
}

func (t dishAdminInteractor) CreateExecuteByAdmin(ctx context.Context, requestParam domain.Dish) (domain.Dish, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	dish, err := t.repo.CreateByAdmin(ctx, requestParam)
	if err != nil {
		return t.presenter.Present(domain.Dish{}), err
	}

	return t.presenter.Present(dish), nil
}

func (t dishAdminInteractor) UpdateExecuteByAdmin(ctx context.Context, requestParam domain.Dish, id int) (domain.Dish, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	dish, err := t.repo.UpdateByAdmin(ctx, requestParam, id)
	if err != nil {
		return t.presenter.Present(domain.Dish{}), err
	}

	return t.presenter.Present(dish), nil
}

func (t dishAdminInteractor) DeleteExecuteByAdmin(ctx context.Context, id int) error {
    // Repositoryを呼び出して削除処理を実行
    return t.repo.DeleteByAdmin(ctx, id)
}
