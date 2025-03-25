package usecase

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

func (t toolCategoryInteractor) FindOneExecute(ctx context.Context, id int) (domain.ToolCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	toolCategory, err := t.repo.FindOne(ctx, id)
	if err != nil {
		return t.presenter.Present(domain.ToolCategory{}), err
	}

	return t.presenter.Present(toolCategory), nil
}

func (t toolCategoryInteractor) FindExecute(ctx context.Context) ([]domain.ToolCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	tool, err := t.repo.Find(ctx)
	if err != nil {
		return t.presenter.Presents([]domain.ToolCategory{}), err
	}

	return t.presenter.Presents(tool), nil
}

func (t toolCategoryAdminInteractor) FindOneExecuteByAdmin(ctx context.Context, id int) (domain.ToolCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	tool, err := t.repo.FindOneByAdmin(ctx, id)
	if err != nil {
		return t.presenter.Present(domain.ToolCategory{}), err
	}

	return t.presenter.Present(tool), nil
}

func (t toolCategoryAdminInteractor) FindAllExecuteByAdmin(ctx context.Context) ([]domain.ToolCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	tools, err := t.repo.FindByAdmin(ctx)
	if err != nil {
		return t.presenter.Presents([]domain.ToolCategory{}), err
	}

	return t.presenter.Presents(tools), nil
}

func (t toolCategoryAdminInteractor) CreateExecuteByAdmin(ctx context.Context, requestParam domain.ToolCategory) (domain.ToolCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	toolCategory, err := t.repo.CreateByAdmin(ctx, requestParam)
	if err != nil {
		return t.presenter.Present(domain.ToolCategory{}), err
	}

	return t.presenter.Present(toolCategory), nil
}

func (t toolCategoryAdminInteractor) UpdateExecuteByAdmin(ctx context.Context, requestParam domain.ToolCategory, id int) (domain.ToolCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	toolCategory, err := t.repo.UpdateByAdmin(ctx, requestParam, id)
	if err != nil {
		return t.presenter.Present(domain.ToolCategory{}), err
	}

	return t.presenter.Present(toolCategory), nil
}

func (t toolCategoryAdminInteractor) DeleteExecuteByAdmin(ctx context.Context, id int) error {
    // Repositoryを呼び出して削除処理を実行
    return t.repo.DeleteByAdmin(ctx, id)
}
