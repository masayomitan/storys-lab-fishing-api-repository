package usecase

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

func (t toolInteractor) FindOneExecute(ctx context.Context, id int) (domain.Tool, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	tool, err := t.repo.FindOne(ctx, id)
	if err != nil {
		return t.presenter.Present(domain.Tool{}), err
	}

	return t.presenter.Present(tool), nil
}

func (t toolInteractor) FindExecute(ctx context.Context) ([]domain.Tool, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	tool, err := t.repo.Find(ctx)
	if err != nil {
		return t.presenter.Presents([]domain.Tool{}), err
	}

	return t.presenter.Presents(tool), nil
}

func (t toolAdminInteractor) FindOneExecuteByAdmin(ctx context.Context, id int) (domain.Tool, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	tool, err := t.repo.FindOneByAdmin(ctx, id)
	if err != nil {
		return t.presenter.Present(domain.Tool{}), err
	}

	return t.presenter.Present(tool), nil
}

func (t toolAdminInteractor) FindAllExecuteByAdmin(ctx context.Context) ([]domain.Tool, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	tooles, err := t.repo.FindByAdmin(ctx)
	if err != nil {
		return t.presenter.Presents([]domain.Tool{}), err
	}

	return t.presenter.Presents(tooles), nil
}

func (t toolAdminInteractor) CreateExecuteByAdmin(ctx context.Context, requestParam domain.Tool) (domain.Tool, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	tool, err := t.repo.CreateByAdmin(ctx, requestParam)
	if err != nil {
		return t.presenter.Present(domain.Tool{}), err
	}

	return t.presenter.Present(tool), nil
}

func (t toolAdminInteractor) UpdateExecuteByAdmin(ctx context.Context, requestParam domain.Tool, id int) (domain.Tool, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	tool, err := t.repo.UpdateByAdmin(ctx, requestParam, id)
	if err != nil {
		return t.presenter.Present(domain.Tool{}), err
	}

	return t.presenter.Present(tool), nil
}

func (t toolAdminInteractor) DeleteExecuteByAdmin(ctx context.Context, id int) error {
    // Repositoryを呼び出して削除処理を実行
    return t.repo.DeleteByAdmin(ctx, id)
}
