package domain

func (Area) TableName() string {
    return "areas"
}

type Area struct {
	ID           int 	`gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	PrefectureId string `json:"prefecture_id"`
	ImageUrl 	 string `json:"image_url"`

	FishingSpots []FishingSpot `gorm:"foreignKey:AreaId"`
	Tides []Tide `gorm:"foreignKey:AreaId"`
}


func NewArea(
	ID 				int,
	name			string,
	description 	string,
	prefectureId 	string,
	imageUrl 		string,

	fishingSpots []FishingSpot,
	tides []Tide,
) Area {
	return Area{
		ID: 			ID,
		Name:			name,
		Description: 	description,
		PrefectureId: 	prefectureId,
		ImageUrl: 		imageUrl,

		FishingSpots: 	fishingSpots,
		Tides: 			tides,
	}
}
