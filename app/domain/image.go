package domain

import "time"

func (Image) TableName() string {
    return "images"
}

type Image struct {
    ID          int             `gorm:"primaryKey" json:"id"`
    Name        string          `json:"name" validate:"required,min=1,max=255"`
    ImageUrl    string          `json:"image_url"`
    S3Url       string          `json:"s3_url"`
    ImageType   int             `json:"image_type"`
	CreatedAt	time.Time       `gorm:"created_at"`
	UpdatedAt	time.Time       `gorm:"updated_at"`
	DeletedAt	*time.Time      `gorm:"default:NULL"`
}

func NewImage(
    ID          int,
    name        string,
    imageUrl    string,
    s3Url       string,
    imageType   int,
    createdAt   time.Time,
    updatedAt   time.Time,
) Image {
    return Image{
        ID:        ID,
        Name:      name,
        ImageUrl:  imageUrl,
        S3Url:     s3Url,
        ImageType: imageType,
        CreatedAt: createdAt,
        UpdatedAt: updatedAt,
    }
}
