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
	InstructorId      string    `json:"instructor_id"`
	AdminId           string    `json:"admin_id"`
	Description       string    `json:"description"`
	IsDisplay         bool      `json:"is_display"`
	PublishedDatetime time.Time `json:"published_datetime"`
	ArticleCategoryId int       `json:"article_category_id"`
	ViewCount         int       `json:"view_count"`

	ArticleCategory ArticleCategory `gorm:"foreignKey:ArticleCategoryId;references:ID"`
	ArticleImages []ArticleImage `gorm:"foreignKey:ArticleId"`
}

type ArticleImage struct {
    ID string `gorm:"primaryKey" json:"id"`
    ArticleId string  `json:"article_id"`
	ImageUrl string `json:"image_url"`
	Sort string `json:"sort"`
	IsMain string `json:"is_main"`
}

func NewArticle(
	ID              	int,
	title           	string,
	subTitle        	string,
	instructorId   		string,
	adminId        	 	string,
	description     	string,
	isDisplay       	bool,
	publishedDatetime 	time.Time,
	articleCategoryId   int,
	viewCount       	int,
	articleCategory 	ArticleCategory,
	articleImages       []ArticleImage,
) Article {
	return Article{
		ID:					ID,
		Title:				title,
		SubTitle:			subTitle,
		InstructorId:		instructorId,
		AdminId:            adminId,
		Description:        description,
		IsDisplay:          isDisplay,
		PublishedDatetime:  publishedDatetime,
		ArticleCategoryId: 	articleCategoryId,
		ViewCount:        	viewCount,
		ArticleCategory:  	articleCategory,
		ArticleImages: 		articleImages,
	}
}
