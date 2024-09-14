package domain

import (
	"time"
)

func (Tide) TableName() string {
	return "tides"
}


type Tide struct {
    ID string `gorm:"primaryKey" json:"id"`
    AreaId string `json:"area_id"`
	PrefectureId string `json:"prefecture_id"`
	Date time.Time `json:"date"`
}
