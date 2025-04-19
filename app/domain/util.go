package domain

import (
	"time"
	"encoding/json"
    "database/sql/driver"
)


func FormatDate(d time.Time) string {
    return d.Format("2006-01-02")
}

func Encode[T any](v T) (driver.Value, error) {
	return json.Marshal(v)
}

func Decode[T any](data interface{}, dest *T) error {
	bytes, _ := data.([]byte)
	return json.Unmarshal(bytes, dest)
}
