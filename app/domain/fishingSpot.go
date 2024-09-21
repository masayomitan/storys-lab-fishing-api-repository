package domain

func (FishingSpot) TableName() string {
	return "fishing_spots"
}

type FishingSpot struct {
	ID          int `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AreaId      int `json:"area_id"`

	Fishes []Fish `gorm:"many2many:fish_fishing_spot_tags;foreignKey:ID;joinForeignKey:FishingSpotId;References:ID;joinReferences:FishId"`
	Area   Area   `gorm:"foreignKey:AreaId;references:ID"`
}


func NewFishingSpot(
	ID int,
	name string,
	description string,
	areaId int,

	fishes []Fish,
	area Area,
) FishingSpot {
	return FishingSpot{
		ID: ID,
		Name: name,
		Description: description,
		AreaId: areaId,

		Area: area,
		Fishes: fishes,
	}
}
