package domain
import "time"

func (Material) TableName() string {
	return "materials"
}

type Material struct {
	ID 			int 		`gorm:"primaryKey" json:"id"`
	Name 		string 		`json:"name"`
	Description string		`json:"description" validate:"max=2000"` 
	CreatedAt	time.Time	`gorm:"created_at" json:"created_at"`
    UpdatedAt	time.Time 	`gorm:"updated_at" json:"updated_at"`
	DeletedAt	*time.Time 	`gorm:"default:NULL"`
}

func NewMaterial(
	ID			int,
	name		string,
	Description	string,
	createdAt	time.Time,
	updatedAt	time.Time,
) Material {
	return Material{
		ID:			ID,
		Name:		name,
		CreatedAt:	createdAt,
		UpdatedAt:	updatedAt,
	}
}
