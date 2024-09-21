package domain

import (
	"database/sql"
	"time"
)

func (Tide) TableName() string {
	return "tides"
}


type Tide struct {
    ID int `gorm:"primaryKey" json:"id"`
	PrefectureId int `json:"prefecture_id"`
	AreaId int `json:"area_id"`
	Date             time.Time `json:"date"`
	FormatDate       string    `json:"format_date"`
	SunriseTime      sql.NullString `json:"sunrise_time"`
	SunsetTime       sql.NullString `json:"sunset_time"`
	MoonriseTime     sql.NullString `json:"moonrise_time"`
	MoonsetTime      sql.NullString `json:"moonset_time"`
	Weather          string    `json:"weather"`
	Temperature      float64   `json:"temperature"`
	WaterTemperature float64   `json:"water_temperature"`
	WindSpeed        float64   `json:"wind_speed"`
	Precipitation    float64   `json:"precipitation"`
	WaveHeight       float64   `json:"wave_height"`
	TideFlow         string    `json:"tide_flow"`
	LowTideTime1     sql.NullString `gorm:"column:low_tide_time_1" json:"low_tide_time_1"`
	HighTideTime1    sql.NullString `gorm:"column:high_tide_time_1" json:"high_tide_time_1"`
	LowTideTime2     sql.NullString `gorm:"column:low_tide_time_2" json:"low_tide_time_2"`
	HighTideTime2    sql.NullString `gorm:"column:high_tide_time_2" json:"high_tide_time_2"`
	LowTideTime3     sql.NullString `gorm:"column:low_tide_time_3" json:"low_tide_time_3"`
	HighTideTime3    sql.NullString `gorm:"column:high_tide_time_3" json:"high_tide_time_3"`
	TideHeight1      int       `gorm:"column:tide_height_1" json:"tide_height_1"`
	TideHeight2      int       `gorm:"column:tide_height_2" json:"tide_height_2"`
	TideHeight3      int       `gorm:"column:tide_height_3" json:"tide_height_3"`
	MoonAge          float64   `gorm:"column:moon_age" json:"moon_age"`
}

// func NewTide(
// 	ID int,
// 	prefectureId int,
// 	areaId int,
// 	date time.Time,
// 	sunriseTime sql.NullString,
// 	sunsetTime sql.NullString,
// 	moonriseTime sql.NullString,
// 	moonsetTime sql.NullString,
// 	weather string,
// 	temperature float64,
// 	waterTemperature float64,
// 	windSpeed float64,
// 	precipitation float64,
// 	waveHeight float64,
// 	tideFlow string,
// 	lowTideTime1 sql.NullString,
// 	highTideTime1 sql.NullString,
// 	lowTideTime2 sql.NullString,
// 	highTideTime2 sql.NullString,
// 	lowTideTime3 sql.NullString,
// 	highTideTime3 sql.NullString,
// 	tideHeight1 int,
// 	tideHeight2 int,
// 	tideHeight3 int,
// 	moonAge float64,
// ) Tide {
// 	return Tide{
// 		ID:               ID,
// 		PrefectureId:     prefectureId,
// 		AreaId: 		  areaId,
// 		Date:             FormatDate(date),
// 		SunriseTime:      sunriseTime,
// 		SunsetTime:       sunsetTime,
// 		MoonriseTime:     moonriseTime,
// 		MoonsetTime:      moonsetTime,
// 		Weather:          weather,
// 		Temperature:      temperature,
// 		WaterTemperature: waterTemperature,
// 		WindSpeed:        windSpeed,
// 		Precipitation:    precipitation,
// 		WaveHeight:       waveHeight,
// 		TideFlow:         tideFlow,
// 		LowTideTime1:     lowTideTime1,
// 		HighTideTime1:    highTideTime1,
// 		LowTideTime2:     lowTideTime2,
// 		HighTideTime2:    highTideTime2,
// 		LowTideTime3:     lowTideTime3,
// 		HighTideTime3:    highTideTime3,
// 		TideHeight1:      tideHeight1,
// 		TideHeight2:      tideHeight2,
// 		TideHeight3:      tideHeight3,
// 		MoonAge:          moonAge,
// 	}
// }
