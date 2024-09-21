package domain

import (
	"errors"
)

var (
	ErrPrefNotFound = errors.New("Pref not found")

	ErrPrefOriginNotFound = errors.New("Pref origin not found")

	ErrPrefDestinationNotFound = errors.New("Pref destination not found")

)

func (Pref) TableName() string {
    return "prefectures"
}

type Pref struct {
	ID string `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	NameKana string `json:"name_kana"`

	Areas []Area `gorm:"foreignKey:PrefectureId"`
}

type PrefID string

func (f PrefID) String() string {
	return string(f)
}

func NewPref(
	ID string,
	name string,
	nameKana string,

	areas []Area,
) Pref {
	return Pref{
		ID: ID,
		Name: name,
		NameKana: nameKana,

		Areas: areas,
	}
}
