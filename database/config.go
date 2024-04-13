
package database

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func GormConnect() *gorm.DB {
	err := godotenv.Load((".env"))
    if err != nil {
			panic ("envファイルの読み込みに失敗しました。")
    }

	var PROTOCOL string
	var PASS string
	
	env := os.Getenv("ENV")
    if env != "local" {
        PROTOCOL = os.Getenv("DB_PROTOCOL_LOCAL")
		PASS = ""
    } else {
        PROTOCOL = os.Getenv("DB_PROTOCOL")
		PASS = os.Getenv("DB_PASS")
    }

    USER := os.Getenv("DB_USER")
    DBNAME := os.Getenv("DB_DATABASE")

	fmt.Println("Database connection info:", USER, PASS, DBNAME, PROTOCOL)

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})

	if err != nil {
		panic (err)
	}
	
	return db
}
