package domain
import "time"

func (Tag) TableName() string {
	return "tags"
}

type Tag struct {
	ID 			int 		`gorm:"primaryKey" json:"id"`
	Name 		string 		`json:"name"`
	CreatedAt	time.Time	`gorm:"created_at" json:"created_at"`
    UpdatedAt	time.Time 	`gorm:"updated_at" json:"updated_at"`
	DeletedAt	*time.Time 	`gorm:"default:NULL"`
}

func NewTag(
	ID 							int,
	name 						string,
	createdAt       			time.Time,
	updatedAt       			time.Time,
) Tag {
	return Tag{
		ID: 							ID,
		Name: 							name,
		CreatedAt:       				createdAt,
		UpdatedAt:      				updatedAt,
	}
}
