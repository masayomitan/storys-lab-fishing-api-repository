package action

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	_ "github.com/go-sql-driver/mysql"
)

type dbConfig struct {
	host       string
	database   string
	port       string
	user       string
	pass       string
	ctxTimeout time.Duration
}

type SecretData struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DBClusterIdentifier string `json:"dbClusterIdentifier"`
	Database            string `json:"database"`
}

func NewConfigDB() *dbConfig {
	return &dbConfig{
		host:       os.Getenv("DB_HOST"),
		port:       os.Getenv("DB_PORT"),
		user:       os.Getenv("DB_USER"),
		pass:       os.Getenv("DB_PASS"),
		database:   os.Getenv("DB_DATABASE"),
		ctxTimeout: 60 * time.Second,
	}
}

func Migration(c *gin.Context) {
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

	cfg := NewConfigDB()
	var connect string
	var migrationsPath string
	// データベースURLを作成
	if (os.Getenv("ENV") == "local") {
		connect = fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
			cfg.user,
			cfg.pass,
			cfg.host,
			cfg.port,
			cfg.database,
		)
		migrationsPath = "file:///app/app//infrastructure/database/migrations"
	} else {
		connect = fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
			secretData.Username,
			secretData.Password,
			secretData.Host,
			secretData.Port,
			secretData.Database,
		)
		migrationsPath = "file:///app/infrastructure/database/migrations"
	}

	log.Printf("Migrations path: %s", migrationsPath)
	fmt.Println("データベースの接続情報:", connect)

	m, err := migrate.New(
		migrationsPath,
		connect,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		http.Error(c.Writer, fmt.Sprintf("マイグレーションに失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	successMessage := "Succeed migration!!"
	log.Println(successMessage)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": successMessage,
	})
}
