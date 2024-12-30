package domain

import "time"

func (Article) TableName() string {
    return "articles"
}

func (ArticleImage) TableName() string {
    return "article_images"
}

type Article struct {
	ID                int       `gorm:"primaryKey" json:"id"`
	Title             string    `json:"title"`
	SubTitle          string    `json:"sub_title"`
	InstructorID      int   	`json:"instructor_id"`
	AdminID           int   	`json:"admin_id"`
	Description       string    `json:"description"`
	IsDisplay         bool      `json:"is_display"`
	PublishedDatetime time.Time `json:"published_datetime"`
	ArticleCategoryID int       `json:"article_category_id"`
	ViewCount         int       `json:"view_count"`

	ArticleCategory ArticleCategory `gorm:"foreignKey:ArticleCategoryID;references:ID"`
	ArticleImages   []ArticleImage  `gorm:"foreignKey:ArticleID"`
	Instructor      Instructor      `gorm:"foreignKey:InstructorID;references:ID"`
	Admin           Admin           `gorm:"foreignKey:AdminID;references:ID"`
}

type ArticleImage struct {
	ID        int 	 `gorm:"primaryKey" json:"id"`
	ArticleID string `json:"article_id"`
	ImageUrl  string `json:"image_url"`
	Sort      string `json:"sort"`
	IsMain    string `json:"is_main"`
}

func NewArticle(
	ID int,
	title string,
	subTitle string,
	instructorID int,
	adminID int,
	description string,
	isDisplay bool,
	publishedDatetime time.Time,
	articleCategoryID int,
	viewCount int,
	articleCategory ArticleCategory,
	articleImages []ArticleImage,
	instructor Instructor,
	admin Admin,
) Article {
	return Article{
		ID:                ID,
		Title:             title,
		SubTitle:          subTitle,
		InstructorID:      instructorID,
		AdminID:           adminID,
		Description:       description,
		IsDisplay:         isDisplay,
		PublishedDatetime: publishedDatetime,
		ArticleCategoryID: articleCategoryID,
		ViewCount:         viewCount,
		ArticleCategory:   articleCategory,
		ArticleImages:     articleImages,
		Instructor:        instructor,
		Admin:             admin,
	}
}
