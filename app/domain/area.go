package domain
import "time"

func (Area) TableName() string {
    return "areas"
}

type Area struct {
	ID           	int 			`gorm:"primaryKey" json:"id"`
	Name         	string 			`json:"name"`
	Description  	string 			`json:"description"`
	PrefectureId 	string 			`json:"prefecture_id"`
	CreatedAt 		time.Time  		`gorm:"created_at" json:"created_at"`
    UpdatedAt 		time.Time  		`gorm:"updated_at" json:"updated_at"`
	DeletedAt  		*time.Time  	`gorm:"default:NULL"`

	FishingSpots 	[]FishingSpot 	`gorm:"foreignKey:AreaID"`
	Tides 			[]Tide 			`gorm:"foreignKey:AreaID"`
	Images 			[]Image 		`gorm:"many2many:areas_to_images;joinForeignKey:AreaID;JoinReferences:ImageID" validate:"-"`
}


func NewArea(
	ID 				int,
	name			string,
	description 	string,
	prefectureId 	string,
	createdAt       time.Time,
	updatedAt       time.Time,

	fishingSpots 	[]FishingSpot,
	tides 			[]Tide,
	images			[]Image,
) Area {
	return Area{
		ID: 			ID,
		Name:			name,
		Description: 	description,
		PrefectureId: 	prefectureId,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,

		FishingSpots: 	fishingSpots,
		Tides: 			tides,
		Images: 		images,
	}
}
