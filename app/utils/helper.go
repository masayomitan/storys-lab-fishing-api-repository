package utils

import (
	"strconv"
	"time"
	"reflect"
)

func StrToInt(idStr string) (int) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0
	}
	return id
}

func SetTimestamps(entity interface{}) {
	val := reflect.ValueOf(entity).Elem()

	// created_at フィールドの設定
	if createdAt := val.FieldByName("CreatedAt"); createdAt.IsValid() && createdAt.CanSet() {
		if createdAt.Kind() == reflect.Struct && createdAt.Type().String() == "time.Time" {
			createdAt.Set(reflect.ValueOf(time.Now()))
		}
	}

	// updated_at フィールドの設定
	if updatedAt := val.FieldByName("UpdatedAt"); updatedAt.IsValid() && updatedAt.CanSet() {
		if updatedAt.Kind() == reflect.Struct && updatedAt.Type().String() == "time.Time" {
			updatedAt.Set(reflect.ValueOf(time.Now()))
		}
	}
}
