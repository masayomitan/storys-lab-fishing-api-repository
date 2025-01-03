package domain
import "time"

func (FishCategory) TableName() string {
    return "fish_categories"
}

func (Fish) TableName() string {
    return "fishes"
}

func (FishImage) TableName() string {
    return "fish_images"
}

type FishCategory struct {
	ID          int    		`gorm:"primaryKey" json:"id"`
	Name        string 		`json:"name" validate:"required,min=1,max=255"`
	// FamilyName            string          `json:"family_name" validate:"omitempty,max=255"`
	Description string 		`json:"description" validate:"max=2000"`       
	CreatedAt   time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}

type Fish struct {
	ID                    int          	  `gorm:"primaryKey" json:"id"`
	Name                  string          `json:"name" validate:"required,min=1,max=255"`
	FamilyName            string          `json:"family_name"`
	ScientificName        string          `json:"scientific_name" validate:"omitempty,max=255"`
	FishCategoryID        int             `json:"fish_category_id" validate:"required"`
	Description           string          `json:"description" validate:"omitempty,max=2000"`
	Length                float64         `gorm:"type:decimal(10,3)" json:"length" validate:"gte=0,lte=1000"`
	Weight                float64         `gorm:"type:decimal(10,3)" json:"weight" validate:"gte=0,lte=100"`
	Habitat               string          `json:"habitat" validate:"omitempty,max=500"`
	DepthRange            int             `json:"depth_range" validate:"gte=0,lte=1000"`
	WaterTemperatureRange int             `json:"water_temperature_range"`
	ConservationStatus    string          `json:"conservation_status" validate:"omitempty,max=100"`
	CreatedAt             time.Time       `json:"created_at"`
	UpdatedAt             time.Time       `json:"updated_at"`

	FishCategory          FishCategory    `gorm:"foreignKey:FishCategoryID" validate:"-"`
	FishingMethods        []FishingMethod `gorm:"many2many:fishing_methods_fishes;foreignKey:ID;joinForeignKey:FishId;References:ID;joinReferences:FishingMethodId" validate:"-"`
	Dishes                []Dish          `gorm:"many2many:fishes_dishes;foreignKey:ID;joinForeignKey:FishId;References:ID;joinReferences:DishId" validate:"-"`
	FishImages            []FishImage     `gorm:"foreignKey:FishID" validate:"-"`
}

type FishImage struct {
    ID 			int 	`gorm:"primaryKey" json:"id"`
	ImageUrl 	string 	`json:"image_url"`
	S3Url 		string 	`json:"s3_url"`
	Sort 		string 	`json:"sort"`
	IsMain 		string 	`json:"is_main"`
}

func NewFish(
	ID int,
	name string,
	familyName string,
	scientificName string,
	fishCategoryID int,
	description string,
	length float64,
	weight float64,
	habitat string,
	depthRange int,
	waterTemperatureRange int,
	conservationStatus string,

	fishCategory FishCategory,
	fishingMethods []FishingMethod,
	dishes []Dish,
	fishImages []FishImage,
) Fish {
	return Fish{
		ID: ID,
		Name: name,
		ScientificName: scientificName,
		FishCategoryID: fishCategoryID,
		Description: description,
		Length: length,
		Weight: weight,
		Habitat: habitat,
		DepthRange: depthRange,
		WaterTemperatureRange: waterTemperatureRange,
		ConservationStatus: conservationStatus,

		FishCategory: fishCategory,
		FishingMethods: fishingMethods,
		Dishes: dishes,
		FishImages: fishImages,
	}
}

func NewFishCategory(
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
