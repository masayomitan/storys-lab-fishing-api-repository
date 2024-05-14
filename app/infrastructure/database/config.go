package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

var db *gorm.DB

type config struct {
	host     string
	database string
	port     string
	user     string
	pass     string

	ctxTimeout time.Duration
}

func NewConfigDB() *config {
	return &config{
		user:       os.Getenv("DB_USER"),
		pass:       os.Getenv("DB_PASS"),
		host:       os.Getenv("DB_HOST"),
		port:       os.Getenv("DB_PORT"),
		database:   os.Getenv("DB_DATABASE"),
		ctxTimeout: 60 * time.Second,
	}
}

func DBConnect() (*gorm.DB, error) {
	err := godotenv.Load((".env"))
	if err != nil {
		return nil, fmt.Errorf("envファイルの読み込みに失敗しました: %w", err)
	}
	cfg := NewConfigDB()

	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.user,
		cfg.pass,
		cfg.host,
		cfg.port,
		cfg.database,
	)

	fmt.Println("データベースの接続情報:", connect)

	db, err = gorm.Open(mysql.Open(connect))
	if err != nil {
		return nil, fmt.Errorf("データベースの接続に失敗しました!!!: %w", err)
	}
	fmt.Println("データベースの接続に成功しました!!!")

	return db, nil
}
