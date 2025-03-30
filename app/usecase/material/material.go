package usecase

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)


func (t materialAdminInteractor) FindExecuteByAdmin(ctx context.Context) ([]domain.Material, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	materials, err := t.repo.FindByAdmin(ctx)
	if err != nil {
		return t.presenter.Presents([]domain.Material{}), err
	}

	return t.presenter.Presents(materials), nil
}
