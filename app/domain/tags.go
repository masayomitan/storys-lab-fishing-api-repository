package domain

func (Tag) TableName() string {
	return "tags"
}

type Tag struct {
	ID int `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
