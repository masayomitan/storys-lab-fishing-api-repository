package domain

func (ArticleCategory) TableName() string {
    return "article_categories"
}

type ArticleCategory struct {
    ID string `gorm:"primaryKey" json:"id"`
    Name string  `json:"name"`
}

