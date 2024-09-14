package domain

func (FishingSpot) TableName() string {
	return "fishing_spots"
}

type FishingSpot struct {
    ID string `gorm:"primaryKey" json:"id"`
    Name string `json:"name"`
	Description string `json:"description"`

	AreaId string `json:"foreignKey:AreaId"`
	Fishes []Fish `gorm:"many2many:fish_fishing_spot_tags;foreignKey:ID;joinForeignKey:FishingSpotId;References:ID;joinReferences:FishId"`
	Area Area `gorm:"foreignKey:AreaId"`
	Tide Tide `gorm:"foreignKey:AreaId"`
}

func NewFishingSpot(
	ID string,
	name string,
	description string,
	areaId string,

	fishes []Fish,
	area Area,
	tide Tide,
) FishingSpot {
	return FishingSpot{
		ID: ID,
		Name: name,
		Description: description,
		AreaId: areaId,

		Area: area,
		Fishes: fishes,
		Tide: tide,
	}
}
