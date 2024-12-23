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
	InstructorId      int   	`json:"instructor_id"`
	AdminId           int   	`json:"admin_id"`
	Description       string    `json:"description"`
	IsDisplay         bool      `json:"is_display"`
	PublishedDatetime time.Time `json:"published_datetime"`
	ArticleCategoryId int       `json:"article_category_id"`
	ViewCount         int       `json:"view_count"`

	ArticleCategory ArticleCategory `gorm:"foreignKey:ArticleCategoryId;references:ID"`
	ArticleImages   []ArticleImage  `gorm:"foreignKey:ArticleId"`
	Instructor      Instructor      `gorm:"foreignKey:InstructorId;references:ID"`
	Admin           Admin           `gorm:"foreignKey:AdminId;references:ID"`
}

type ArticleImage struct {
	ID        int 	 `gorm:"primaryKey" json:"id"`
	ArticleId string `json:"article_id"`
	ImageUrl  string `json:"image_url"`
	Sort      string `json:"sort"`
	IsMain    string `json:"is_main"`
}

func NewArticle(
	ID int,
	title string,
	subTitle string,
	instructorId int,
	adminId int,
	description string,
	isDisplay bool,
	publishedDatetime time.Time,
	articleCategoryId int,
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
		InstructorId:      instructorId,
		AdminId:           adminId,
		Description:       description,
		IsDisplay:         isDisplay,
		PublishedDatetime: publishedDatetime,
		ArticleCategoryId: articleCategoryId,
		ViewCount:         viewCount,
		ArticleCategory:   articleCategory,
		ArticleImages:     articleImages,
		Instructor:        instructor,
		Admin:             admin,
	}
}
