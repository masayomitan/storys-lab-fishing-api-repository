
package database

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"fmt"
	"os"
	"time"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

type config struct {
	host     string
	database string
	port     string
	user     string
	password string

	ctxTimeout time.Duration
}

func NewConfigDB() *config {
	return &config{
		host:       os.Getenv("DB_HOST"),
		database:   os.Getenv("DB_DATABASE"),
		password:   os.Getenv("DB_PASS"),
		user:       os.Getenv("DB_USER"),
		ctxTimeout: 60 * time.Second,
	}
}

func DBConnect(cfg *config) (*gorm.DB, error) {
	err := godotenv.Load((".env"))
    if err != nil {
		return nil, fmt.Errorf("envファイルの読み込みに失敗しました: %w", err)
    }

	var PROTOCOL string
	var PASS string
	
	env := os.Getenv("ENV")
    if env == "dev" {
        PROTOCOL = os.Getenv("DB_PROTOCOL")
		PASS = os.Getenv("DB_PASS")
    }
    USER := os.Getenv("DB_USER")
    DBNAME := os.Getenv("DB_DATABASE")

	fmt.Println("Database connection info:", USER, PASS, DBNAME, PROTOCOL)

	CONNECT := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
        cfg.user, cfg.password, cfg.host, cfg.port, cfg.database)
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("データベースの接続に失敗しました: %w", err)
	}
	
	return db, nil
}

