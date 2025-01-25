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

func SetCreateTimestamps(entity interface{}) {
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

func SetUpdateTimestamps(entity interface{}) {
	val := reflect.ValueOf(entity).Elem()

	// updated_at フィールドの設定
	if updatedAt := val.FieldByName("UpdatedAt"); updatedAt.IsValid() && updatedAt.CanSet() {
		if updatedAt.Kind() == reflect.Struct && updatedAt.Type().String() == "time.Time" {
			updatedAt.Set(reflect.ValueOf(time.Now()))
		}
	}
}

func SetTimestampToDelete(entity interface{}) {
	val := reflect.ValueOf(entity).Elem()

	// deleted_at フィールドの設定
	if DeletedAt := val.FieldByName("DeletedAt"); DeletedAt.IsValid() && DeletedAt.CanSet() {
		if DeletedAt.Kind() == reflect.Struct && DeletedAt.Type().String() == "time.Time" {
			DeletedAt.Set(reflect.ValueOf(time.Now()))
		}
	}
}
