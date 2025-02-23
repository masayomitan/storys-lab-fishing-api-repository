package presenter

import (
	// "time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/prefecture"
)

type PrefecturePresenter struct{}

func NewPrefecturePresenter() usecase.PrefecturePresenter {
	return PrefecturePresenter{}
}

func (a PrefecturePresenter) Present(prefecture domain.Prefecture) domain.Prefecture {
	return domain.Prefecture{
		ID: prefecture.ID,
		Name: prefecture.Name,
		NameKana: prefecture.NameKana,
		ImageUrl: prefecture.ImageUrl,
		Areas: prefecture.Areas,
	}
}

func (a PrefecturePresenter) Presents(prefectures []domain.Prefecture) []domain.Prefecture {
	var output = make([]domain.Prefecture, 0)
	for _, prefectures := range prefectures {
		output = append(output, domain.Prefecture{
			ID: prefectures.ID,
			Name: prefectures.Name,
			NameKana: prefectures.NameKana,
			ImageUrl: prefectures.ImageUrl,
			Areas: prefectures.Areas,
		})
	}
	return output
}

