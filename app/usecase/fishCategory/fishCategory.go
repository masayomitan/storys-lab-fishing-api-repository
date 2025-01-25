package usecase

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)

func (t fishCategoryAdminInteractor) FindOneExecuteByAdmin(ctx context.Context, id int) (domain.FishCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fish, err := t.repo.FindOneByAdmin(ctx, id)
	if err != nil {
		return t.presenter.PresentOne(domain.FishCategory{}), err
	}

	return t.presenter.PresentOne(fish), nil
}

func (t fishCategoryAdminInteractor) FindAllExecuteByAdmin(ctx context.Context) ([]domain.FishCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fishes, err := t.repo.FindAllByAdmin(ctx)
	if err != nil {
		return t.presenter.PresentAll([]domain.FishCategory{}), err
	}

	return t.presenter.PresentAll(fishes), nil
}

func (t fishCategoryAdminInteractor) CreateExecuteByAdmin(ctx context.Context, requestParam domain.FishCategory) (domain.FishCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fishCategory, err := t.repo.CreateByAdmin(ctx, requestParam)
	if err != nil {
		return t.presenter.PresentOne(domain.FishCategory{}), err
	}

	return t.presenter.PresentOne(fishCategory), nil
}

func (t fishCategoryAdminInteractor) UpdateExecuteByAdmin(ctx context.Context, requestParam domain.FishCategory, id int) (domain.FishCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	fishCategory, err := t.repo.CreateByAdmin(ctx, requestParam)
	if err != nil {
		return t.presenter.PresentOne(domain.FishCategory{}), err
	}

	return t.presenter.PresentOne(fishCategory), nil
}

func (t fishCategoryAdminInteractor) DeleteExecuteByAdmin(ctx context.Context, id int) error {
    // Repositoryを呼び出して削除処理を実行
    return t.repo.DeleteByAdmin(ctx, id)
}
