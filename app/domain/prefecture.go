package domain

func (Pref) TableName() string {
    return "prefectures"
}

type Pref struct {
	ID string `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	NameKana string `json:"name_kana"`
	ImageUrl string `json:"image_url"`

	Areas []Area `gorm:"foreignKey:PrefectureId"`
}

type PrefID string

func (f PrefID) String() string {
	return string(f)
}

func NewPref(
	ID string,
	name string,
	nameKana string,
	imageUrl string,

	areas []Area,
) Pref {
	return Pref{
		ID: ID,
		Name: name,
		NameKana: nameKana,
		ImageUrl: imageUrl,
		Areas: areas,
	}
}
