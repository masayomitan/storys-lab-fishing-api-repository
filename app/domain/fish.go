package domain

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
    ID string `gorm:"primaryKey" json:"id"`
    Name string  `json:"name"`
	Description string `json:"description"`
}

type Fish struct {
	ID string `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	FamilyName string `json:"family_name"`
	ScientificName string `json:"scientific_name"`
	FishCategoryId int `json:"fish_category"`
	Description string `json:"description"`
	Length float64 `json:"length"`
	Weight float64 `json:"weight"`
	Habitat string `json:"habitat"`
	DepthRange string `json:"depth_range"`
	WaterTemperatureRange string `json:"water_temperature_range"`
	ConservationStatus string `json:"conservation_status"`

	FishCategory FishCategory `gorm:"foreignKey:FishCategoryId"`
	FishingMethods []FishingMethod `gorm:"many2many:fishing_methods_fishes;foreignKey:ID;joinForeignKey:FishId;References:ID;joinReferences:FishingMethodId"`
	Dishes []Dish `gorm:"many2many:fishes_dishes;foreignKey:ID;joinForeignKey:FishId;References:ID;joinReferences:DishId"`
	FishImages []FishImage `gorm:"foreignKey:FishId"`
}

type FishImage struct {
    ID string `gorm:"primaryKey" json:"id"`
    FishId string  `json:"fish_id"`
	ImageUrl string `json:"image_url"`
	Sort string `json:"sort"`
	IsMain string `json:"is_main"`
}

func NewFish(
	ID string,
	name string,
	familyName string,
	scientificName string,
	fishCategoryId int,
	description string,
	length float64,
	weight float64,
	habitat string,
	depthRange string,
	waterTemperatureRange string,
	conservationStatus string,

	fishCategory FishCategory,
	fishingMethods []FishingMethod,
	dishes []Dish,
	fishImages []FishImage,
) Fish {
	return Fish{
		ID: ID,
		Name: name,
		FamilyName: familyName,
		ScientificName: scientificName,
		FishCategoryId: fishCategoryId,
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
