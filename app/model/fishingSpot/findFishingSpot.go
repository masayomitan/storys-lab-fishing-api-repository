package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
)

func (a FishingSpotSQL) FindOne(ctx context.Context, id string) (domain.FishingSpot, error) {
	var json = domain.FishingSpot{}

	if err := a.db.FindOneFishingSpot(ctx, a.tableName, id, &json); err != nil {
		return domain.FishingSpot{}, errors.Wrap(err, "error listing fishes")
	}

	var fishingSpot = domain.NewFishingSpot(
		json.ID,
		json.Name,
		json.Description,
		json.AreaId,
		convertArea(json.Area),
		convertFishes(json.Fishes),
		convertTags(json.Tags),
	)

	fmt.Println("fishingSpot")
	fmt.Println(fishingSpot)

	return fishingSpot, nil
}

func (a FishingSpotSQL) FindAllByAreaId(ctx context.Context, id int) ([]domain.FishingSpot, error) {
	var json = []domain.FishingSpot{}

	if err := a.db.FindAllByAreaId(ctx, a.tableName, id, &json); err != nil {
		return []domain.FishingSpot{}, errors.Wrap(err, "error listing fishes")
	}

	var fishingSpots = make([]domain.FishingSpot, 0)
	for _, json := range json {
		var fishingSpot = domain.NewFishingSpot(
			json.ID,
			json.Name,
			json.Description,
			json.AreaId,
			convertArea(json.Area),
			convertFishes(json.Fishes),
			convertTags(json.Tags),
		)
		fishingSpots = append(fishingSpots, fishingSpot)
	}
	
	fmt.Println("fishingSpot")
	fmt.Println(fishingSpots)

	return fishingSpots, nil
}

func (ga *GormAdapter) FindOneFishingSpot(ctx context.Context, table string, id string, result interface{}) error {
	return ga.DB.Table(table).Where("id = ?", id).
		Preload("Area").
		Preload("Area.Tides").
		Preload("Fishes").
		Preload("Tags").
		First(result).Error
}

func (ga *GormAdapter) FindAllByAreaId(ctx context.Context, table string, area_id int, result interface{}) error {
	return ga.DB.Table(table).Where("area_id = ?", area_id).
		Find(result).Error
}

func convertFishes(fishes []domain.Fish) []domain.Fish {
	var result []domain.Fish
	for _, f := range fishes {
		result = append(result, domain.Fish{
			ID:        f.ID,
			Name:      f.Name,
			FamilyName: f.FamilyName,
		})
	}
	return result
}

func convertArea(a domain.Area) domain.Area {
	return domain.Area {
		ID:           a.ID,
		Name:         a.Name,
		Description:  a.Description,
		PrefectureId: a.PrefectureId,
		Tides:        convertTides(a.Tides),
	}
}

func convertTags(tags []domain.Tag) []domain.Tag {
	var result []domain.Tag
	for _, t := range tags {
		result = append(result, domain.Tag{
			ID:        t.ID,
			Name:      t.Name,
		})
	}
	return result
}

func convertTides(tides []domain.Tide) []domain.Tide {
	var result []domain.Tide
	for _, t := range tides {
		formatDate := domain.FormatDate(t.Date)
		result = append(result, domain.Tide{
			ID:               t.ID,
			AreaId:           t.AreaId,
			PrefectureId:     t.PrefectureId,
			Date:             t.Date,
			FormatDate:       formatDate,
			SunriseTime:      t.SunriseTime,
			SunsetTime:       t.SunsetTime,
			MoonriseTime:     t.MoonriseTime,
			MoonsetTime:      t.MoonsetTime,
			Weather:          t.Weather,
			Temperature:      t.Temperature,
			WaterTemperature: t.WaterTemperature,
			WindSpeed:        t.WindSpeed,
			Precipitation:    t.Precipitation,
			WaveHeight:       t.WaveHeight,
			TideFlow:         t.TideFlow,
			LowTideTime1:     t.LowTideTime1,
			HighTideTime1:    t.HighTideTime1,
			LowTideTime2:     t.LowTideTime2,
			HighTideTime2:    t.HighTideTime2,
			LowTideTime3:     t.LowTideTime3,
			HighTideTime3:    t.HighTideTime3,
			TideHeight1:      t.TideHeight1,
			TideHeight2:      t.TideHeight2,
			TideHeight3:      t.TideHeight3,
			MoonAge:          t.MoonAge,
		})
	}
	return result
}

