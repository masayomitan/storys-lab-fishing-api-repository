package domain
import "time"

func (Dish) TableName() string {
    return "dishes"
}

type Dish struct {
    ID 			int 		`gorm:"primaryKey" json:"id"`
    Name 		string 		`json:"name"`
	Description string 		`json:"description"`
	Ingredients string 		`json:"ingredients"`
	Kind 		string 		`json:"kind"`
	Level 		int 		`json:"level"`
	CreatedAt	time.Time	`gorm:"created_at" json:"created_at"`
    UpdatedAt	time.Time  	`gorm:"updated_at" json:"updated_at"`
	DeletedAt	*time.Time  `gorm:"default:NULL"`

	Images		[]Image		`gorm:"many2many:dishes_to_images;" validate:"-"`
	
}



func NewDish(
    ID              int,
    name            string,
    description     string,
    ingredients  	string,
    kind      		string,
    level           int,
	createdAt		time.Time,
	updatedAt		time.Time,

    images          []Image,
) Dish {
    return Dish{
        ID:             ID,
        Name:           name,
        Description:    description,
        Ingredients: 	ingredients,
        Kind:     		kind,
        Level:          level,
        Images:         images,
		CreatedAt:		createdAt,
		UpdatedAt:      updatedAt,
    }
}

