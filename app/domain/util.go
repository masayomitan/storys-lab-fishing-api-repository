package domain

import (
	gouuid "github.com/satori/go.uuid"
	"time"
)

func NewUUID() string {
	return gouuid.NewV4().String()
}

func IsValidUUID(uuid string) bool {
	_, err := gouuid.FromString(uuid)
	return err == nil
}

func FormatDate(d time.Time) string {
    return d.Format("2006-01-02")
}
