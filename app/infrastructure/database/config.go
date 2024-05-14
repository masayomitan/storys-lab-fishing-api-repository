package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var db *gorm.DB

type dbConfig struct {
	host     string
	database string
	port     string
	user     string
	pass     string

	ctxTimeout time.Duration
}

type SecretData struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DBClusterIdentifier string `json:"dbClusterIdentifier"`
}

func NewConfigDB() *dbConfig {
	return &dbConfig{
		user:       os.Getenv("DB_USER"),
		pass:       os.Getenv("DB_PASS"),
		host:       os.Getenv("DB_HOST"),
		port:       os.Getenv("DB_PORT"),
		database:   os.Getenv("DB_DATABASE"),
		ctxTimeout: 60 * time.Second,
	}
}

func DBConnect() (*gorm.DB, error) {

	secretName := "prod/fishing-db"
	region := "ap-northeast-1"

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		// https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
		log.Fatal(err.Error())
	}

	// Decrypts secret using the associated KMS key.
	var secretString string = *result.SecretString
	var secretData SecretData
	err = json.Unmarshal([]byte(secretString), &secretData)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Username:", secretData.Username)
	fmt.Println("Password:", secretData.Password)
	fmt.Println("Engine:", secretData.Engine)
	fmt.Println("Host:", secretData.Host)
	fmt.Println("Port:", secretData.Port)
	fmt.Println("DBClusterIdentifier:", secretData.DBClusterIdentifier)

	err = godotenv.Load((".env"))
	if err != nil {
		return nil, fmt.Errorf("envファイルの読み込みに失敗しました: %w", err)
	}
	cfg := NewConfigDB()
	fmt.Println(cfg)

	// connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	cfg.user,
	// 	cfg.pass,
	// 	cfg.host,
	// 	cfg.port,
	// 	cfg.database,
	// )
	connect := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		secretData.Username,
		secretData.Password,
		secretData.Host,
		secretData.Port,
		cfg.database,
	)

	fmt.Println("データベースの接続情報:", connect)

	db, err = gorm.Open(mysql.Open(connect))
	if err != nil {
		return nil, fmt.Errorf("データベースの接続に失敗しました: %w", err)
	}
	fmt.Println("データベースの接続に成功しました!")

	return db, nil
}
