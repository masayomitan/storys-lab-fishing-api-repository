package usecase

import (
	"context"
	"storys-lab-fishing-api/app/domain"
)


func (t tagAdminInteractor) FindExecuteByAdmin(ctx context.Context) ([]domain.Tag, error) {
	ctx, cancel := context.WithTimeout(ctx, t.ctxTimeout)
	defer cancel()

	tags, err := t.repo.FindByAdmin(ctx)
	if err != nil {
		return t.presenter.Presents([]domain.Tag{}), err
	}

	return t.presenter.Presents(tags), nil
}
