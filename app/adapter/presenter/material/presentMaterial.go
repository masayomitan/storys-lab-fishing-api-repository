package presenter

import (
	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/material"
)

type materialPresenter struct{}

func NewMaterialPresenter() usecase.MaterialPresenter {
	return materialPresenter{}
}

func (a materialPresenter) Present(material domain.Material) domain.Material {
	return domain.Material{
		ID:             			material.ID,
		Name:           			material.Name,
		Description:				material.Description,
		CreatedAt:					material.CreatedAt,
		UpdatedAt:					material.UpdatedAt,
	}
}

func (a materialPresenter) Presents(material []domain.Material) []domain.Material {
	var output = make([]domain.Material, 0)
	for _, material := range material {
		output = append(output, domain.Material{
			ID:             			material.ID,
			Name:           			material.Name,
			Description:				material.Description,
			CreatedAt:					material.CreatedAt,
			UpdatedAt:					material.UpdatedAt,
		})
	}
	return output
}
