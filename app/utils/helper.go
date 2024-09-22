package utils

import (
	"strconv"
)

func StrToInt(idStr string) (int) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0
	}
	return id
}
