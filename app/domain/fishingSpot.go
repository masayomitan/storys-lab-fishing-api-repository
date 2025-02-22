package domain

func (FishingSpot) TableName() string {
	return "fishing_spots"
}

type FishingSpot struct {
	ID          	int `gorm:"primaryKey" json:"id"`
	Name        	string `json:"name"`
	ImageUrl 	 	string `json:"image_url"`
	Description 	string `json:"description"`
	AreaID      	int `json:"area_id"`

	Fishes 			[]Fish `gorm:"many2many:fish_fishing_spot_time_tags;foreignKey:ID;joinForeignKey:FishingSpotId;References:ID;joinReferences:FishId"`
	Area   			Area   `gorm:"foreignKey:AreaID;references:ID"`
	Tags 			[]Tag `gorm:"many2many:fishing_spots_to_tags;foreignKey:ID;joinForeignKey:FishingSpotId;References:ID;joinReferences:TagId"`
}


func NewFishingSpot(
	ID 			int,
	name 		string,
	imageUrl 	string,
	description string,
	areaID 		int,

	area 		Area,
	fishes 		[]Fish,
	tags 		[]Tag,
) FishingSpot {
	return FishingSpot{
		ID: 			ID,
		Name: 			name,
		ImageUrl: 		imageUrl,
		Description: 	description,
		AreaID: 		areaID,

		Area: 			area,
		Fishes: 		fishes,
		Tags: 			tags,
	}
}
