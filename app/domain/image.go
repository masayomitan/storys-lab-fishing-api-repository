package domain
import "time"

func (Image) TableName() string {
    return "images"
}

type Image struct {
    ID 			int 		`gorm:"primaryKey" json:"id"`
	ImageUrl 	string 		`json:"image_url"`
	S3Url 		string 		`json:"s3_url"`
	Type 		int 		`json:"type"`
	Sort 		int 		`json:"sort"`
	IsMain 		int 		`json:"is_main"`
	CreatedAt   time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}

func NewImage(
	ID int,
	name string,
	// familyName: string,
	description string,
	createdAt time.Time,
	updatedAt time.Time,
) FishCategory {
	return FishCategory{
		ID: ID,
		Name: name,
		// FamilyName: familyName,
		Description: description,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
