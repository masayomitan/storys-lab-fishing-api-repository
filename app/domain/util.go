package domain

import (
	"time"
)


func FormatDate(d time.Time) string {
    return d.Format("2006-01-02")
}
