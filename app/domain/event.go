package domain

import "time"

func (Event) TableName() string {
    return "events"
}

type Event struct {
    ID                  int        `gorm:"primaryKey" json:"id"`
    Title               string     `json:"title"`
    SubTitle            string     `json:"sub_title"`
    Description         string     `json:"description"`
    StartDatetime       time.Time  `json:"start_datetime"`
    EndDatetime         time.Time  `json:"end_datetime"`
    AdminID             int        `json:"admin_id"`
    InstructorID        int        `json:"instructor_id"`
    FishingSpotID       int        `json:"fishing_spot_id"`
    AdditionalInfo      string     `json:"additional_information"`
    MaxPersons          int        `json:"max_persons"`
    MinPersons          int        `json:"min_persons"`
    IsDisplay           bool       `json:"is_display"`

    Instructor          Instructor      `gorm:"foreignKey:InstructorId;references:ID"`
    Admin               Admin           `gorm:"foreignKey:AdminId;references:ID"`
}

type EventImage struct {
    ID        int    `gorm:"primaryKey" json:"id"`
    EventID   string `json:"event_id"`
    ImageUrl  string `json:"image_url"`
    Sort      string `json:"sort"`
    IsMain    string `json:"is_main"`
}

// NewEvent creates a new Event instance
func NewEvent(
    ID int,
    title string,
    subTitle string,
    description string,
    startDatetime time.Time,
    endDatetime time.Time,
    adminID int,
    instructorID int,
    fishingSpotID int,
    additionalInfo string,
    maxPersons int,
    minPersons int,
    isDisplay bool,
    instructor Instructor,
    admin Admin,
) Event {
    return Event{
        ID:              ID,
        Title:           title,
        SubTitle:        subTitle,
        Description:     description,
        StartDatetime:   startDatetime,
        EndDatetime:     endDatetime,
        AdminID:         adminID,
        InstructorID:    instructorID,
        FishingSpotID:   fishingSpotID,
        AdditionalInfo:  additionalInfo,
        MaxPersons:      maxPersons,
        MinPersons:      minPersons,
        IsDisplay:       isDisplay,
        Instructor:      instructor,
        Admin:           admin,
    }
}
