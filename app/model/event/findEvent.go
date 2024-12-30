package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
)

func (a EventSQL) FindAll(ctx context.Context) ([]domain.Event, error) {
	var json = make([]domain.Event, 0)

	if err := a.db.FindAll(ctx, a.tableName, domain.Event{}, &json); err != nil {
		return []domain.Event{}, errors.Wrap(err, "error listing Events")
	}
	var events = make([]domain.Event, 0)

	for _, json := range json {
		var event = domain.NewEvent(
			json.ID,
			json.Title,
			json.SubTitle,
			json.Description,
			json.StartDatetime,
			json.EndDatetime,
			json.AdminID,
			json.InstructorID,
			json.FishingSpotID,
			json.AdditionalInfo,
			json.MaxPersons,
			json.MinPersons,
			json.IsDisplay,
			json.Instructor,
			json.Admin,
		)
		events = append(events, event)
	}
	return events, nil
}

func (a EventSQL) FindOne(ctx context.Context, id int) (domain.Event, error) {
	var json = domain.Event{}

	if err := a.db.FindOne(ctx, a.tableName, id, &json); err != nil {
		return domain.Event{}, errors.Wrap(err, "error listing Event")
	}

	var event = domain.NewEvent(
		json.ID,
		json.Title,
		json.SubTitle,
		json.Description,
		json.StartDatetime,
		json.EndDatetime,
		json.AdminID,
		json.InstructorID,
		json.FishingSpotID,
		json.AdditionalInfo,
		json.MaxPersons,
		json.MinPersons,
		json.IsDisplay,
		json.Instructor,
		json.Admin,
	)

	fmt.Println("")
	return event, nil
}

func (ga *GormAdapter) FindAll(ctx context.Context, table string, query interface{}, result interface{}) error {
    return ga.DB.Table(table).Where(query).
		// Preload("EventImages").
		Order("id asc").
		Find(result).Error
}

func (ga *GormAdapter) FindOne(ctx context.Context, table string, event_id int, result interface{}) error {
	return ga.DB.Table(table).Where("id = ?", event_id).
		// Preload("EventImages").
		First(result).Error
}
