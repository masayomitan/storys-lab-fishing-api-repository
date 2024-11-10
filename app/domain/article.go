package domain

import "time"

func (Article) TableName() string {
    return "articles"
}

type Article struct {
	ID                int       `gorm:"primaryKey" json:"id"`
	Title             string    `json:"title"`
	SubTitle          string    `json:"sub_title"`
	InstructorId      string    `json:"instructor_id"`
	AdminId           string    `json:"admin_id"`
	Description       string    `json:"description"`
	IsDisplay         bool      `json:"is_display"`
	PublishedDateTime time.Time `json:"published_date_time"`
	CategoryId        int       `json:"category_id"`
	ViewCount         int       `json:"view_count"`
}

func NewArticle(
	ID              int,
	title           string,
	subTitle        string,
	instructorId    string,
	adminId         string,
	description     string,
	isDisplay       bool,
	publishedDateTime time.Time,
	categoryId      int,
	viewCount       int,
) Article {
	return Article{
		ID:               ID,
		Title:            title,
		SubTitle:         subTitle,
		InstructorId:     instructorId,
		AdminId:          adminId,
		Description:      description,
		IsDisplay:        isDisplay,
		PublishedDateTime: publishedDateTime,
		CategoryId:       categoryId,
		ViewCount:        viewCount,
	}
}
