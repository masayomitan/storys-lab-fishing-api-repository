package presenter

import (
	// "time"

	"storys-lab-fishing-api/app/domain"
	"storys-lab-fishing-api/app/usecase/event"
)

type findAllEventPresenter struct{}
type findOneEventPresenter struct{}

func NewFindAllEventPresenter() usecase.FindAllEventPresenter {
	return findAllEventPresenter{}
}

func NewFindOneEventPresenter() usecase.FindOneEventPresenter {
	return findOneEventPresenter{}
}

func (a findOneEventPresenter) Output(event domain.Event) domain.Event {
	return domain.Event{
		ID:                event.ID,
		Title:             event.Title,
		SubTitle:          event.SubTitle,
		Description:       event.Description,
		StartDatetime:     event.StartDatetime,
		EndDatetime:       event.EndDatetime,
		AdminID:           event.AdminID,
		InstructorID:      event.InstructorID,
		FishingSpotID:     event.FishingSpotID,
		AdditionalInfo:    event.AdditionalInfo,
		MaxPersons:        event.MaxPersons,
		MinPersons:        event.MinPersons,
		IsDisplay:         event.IsDisplay,

	}
}


func (a findAllEventPresenter) Output(events []domain.Event) []domain.Event {
	var output = make([]domain.Event, 0)
	for _, event := range events {
		output = append(output, domain.Event{
			ID:                event.ID,
			Title:             event.Title,
			SubTitle:          event.SubTitle,
			Description:       event.Description,
			StartDatetime:     event.StartDatetime,
			EndDatetime:       event.EndDatetime,
			AdminID:           event.AdminID,
			InstructorID:      event.InstructorID,
			FishingSpotID:     event.FishingSpotID,
			AdditionalInfo:    event.AdditionalInfo,
			MaxPersons:        event.MaxPersons,
			MinPersons:        event.MinPersons,
			IsDisplay:         event.IsDisplay,
		})
	}
	return output
}
