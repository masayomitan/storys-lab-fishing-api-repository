package usecase

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

func (a areaInteractor) FindOneExecute(ctx context.Context, id int) (domain.Area, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	area, err := a.repo.FindOne(ctx, id)
	if err != nil {
		return a.presenter.PresentOne(domain.Area{}), err
	}

	return a.presenter.PresentOne(area), nil
}


func (t areaInteractor) FindAllExecute(ctx context.Context) ([]domain.Area, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	areas, err := t.repo.FindAll(ctx)
	if err != nil {
		return t.presenter.PresentAll([]domain.Area{}), err
	}

	return t.presenter.PresentAll(areas), nil
}

func (t areaAdminInteractor) FindExecuteByAdmin(ctx context.Context) ([]domain.Area, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	areas, err := t.repo.FindByAdmin(ctx)
	if err != nil {
		return t.presenter.PresentAll([]domain.Area{}), err
	}

	return t.presenter.PresentAll(areas), nil
}

func (t areaAdminInteractor) FindOneExecuteByAdmin(ctx context.Context, id int) (domain.Area, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	Area, err := t.repo.FindOneByAdmin(ctx, id)
	if err != nil {
		return t.presenter.PresentOne(domain.Area{}), err
	}

	return t.presenter.PresentOne(Area), nil
}

func (t areaAdminInteractor) CreateExecuteByAdmin(ctx context.Context, requestParam domain.Area) (domain.Area, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	area, err := t.repo.CreateByAdmin(ctx, requestParam)
	if err != nil {
		return t.presenter.PresentOne(domain.Area{}), err
	}

	return t.presenter.PresentOne(area), nil
}

func (t areaAdminInteractor) UpdateExecuteByAdmin(ctx context.Context, requestParam domain.Area, id int) (domain.Area, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	area, err := t.repo.UpdateByAdmin(ctx, requestParam, id)
	if err != nil {
		return t.presenter.PresentOne(domain.Area{}), err
	}

	return t.presenter.PresentOne(area), nil
}

func (t areaAdminInteractor) DeleteExecuteByAdmin(ctx context.Context, id int) error {
    // Repositoryを呼び出して削除処理を実行
    return t.repo.DeleteByAdmin(ctx, id)
}

