package model

import (
	"storys-lab-fishing-api/app/domain"
)

func convertFishes(fishes []domain.Fish) []domain.Fish {
	var result []domain.Fish
	for _, f := range fishes {
		result = append(result, domain.Fish{
			ID:        f.ID,
			Name:      f.Name,
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
			AreaID:           t.AreaID,
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
