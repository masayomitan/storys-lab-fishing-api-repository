package domain

func (Prefecture) TableName() string {
    return "prefectures"
}

type Prefecture struct {
	ID int `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	NameKana string `json:"name_kana"`
	ImageUrl string `json:"image_url"`

	Areas []Area `gorm:"foreignKey:PrefectureID"`
}

func NewPrefecture(
	ID int,
	name string,
	nameKana string,
	imageUrl string,

	areas []Area,
) Prefecture {
	return Prefecture{
		ID: ID,
		Name: name,
		NameKana: nameKana,
		ImageUrl: imageUrl,
		Areas: areas,
	}
}
