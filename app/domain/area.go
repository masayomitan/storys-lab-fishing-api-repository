package domain

func (Area) TableName() string {
    return "areas"
}

type Area struct {
	ID          string `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PrefectureId string `json:"prefecture_id"`

	FishingSpots []FishingSpot `gorm:"foreignKey:AreaId"`
	Tides []Tide `gorm:"foreignKey:AreaId"`
}


func NewArea(
	ID string,
	name string,
	description string,
	prefectureId string,

	fishingSpots []FishingSpot,
	tides []Tide,
) Area {
	return Area{
		ID: ID,
		Name: name,
		Description: description,
		PrefectureId: prefectureId,

		FishingSpots: fishingSpots,
		Tides: tides,
	}
}
