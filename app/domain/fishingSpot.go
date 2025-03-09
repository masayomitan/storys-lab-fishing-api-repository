package domain
import "time"

func (FishingSpot) TableName() string {
	return "fishing_spots"
}

type FishingSpot struct {
	ID          					int 		`gorm:"primaryKey" json:"id"`
	Name        					string 		`json:"name"`
	Description 					string 		`json:"description"`
	AreaID      					int 		`json:"area_id"`
	RecommendedFishingMethods 		string 		`json:"recommended_fishing_methods"`
	CreatedAt   					time.Time	`gorm:"created_at" json:"created_at"`
    UpdatedAt 						time.Time 	`gorm:"updated_at" json:"updated_at"`
	DeletedAt  						*time.Time 	`gorm:"default:NULL"`

	Fishes 							[]Fish 		`gorm:"many2many:fish_fishing_spot_time_tags;foreignKey:ID;joinForeignKey:FishingSpotId;References:ID;joinReferences:FishId"`
	Area   							Area 		`gorm:"foreignKey:AreaID;references:ID"`
	Tags 							[]Tag 		`gorm:"many2many:fishing_spots_to_tags;foreignKey:ID;joinForeignKey:FishingSpotId;References:ID;joinReferences:TagId"`
	Images 							[]Image 	`gorm:"many2many:fishing_spots_to_images;" validate:"-"`
}


func NewFishingSpot(
	ID 							int,
	name 						string,
	description 				string,
	areaID 						int,
	recommendedFishingMethods 	string,
	createdAt       			time.Time,
	updatedAt       			time.Time,

	area 						Area,
	fishes 						[]Fish,
	tags 						[]Tag,
	images						[]Image,
) FishingSpot {
	return FishingSpot{
		ID: 							ID,
		Name: 							name,
		Description: 					description,
		AreaID: 						areaID,
		RecommendedFishingMethods: 		recommendedFishingMethods,
		CreatedAt:       				createdAt,
		UpdatedAt:      				updatedAt,
		Area: 							area,
		Fishes: 						fishes,
		Tags: 							tags,
		Images: 						images,
	}
}
