package domain

func (Area) TableName() string {
    return "areas"
}

type Area struct {
    ID string `gorm:"primaryKey" json:"id"`
    Name string  `json:"name"`
	Description string `json:"description"`
	PrefectureId string `json:"foreignKey:Prefecture_id"`
}


func NewArea(
	ID string,
	name string,
	description string,
	prefectureId string,
) Area {
	return Area{
		ID:                      ID,
		Name:                    name,
		Description:             description,
		PrefectureId:          	 prefectureId,
	}
}
