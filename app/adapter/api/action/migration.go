package action

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	_ "github.com/go-sql-driver/mysql"
)

type SecretData struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DBClusterIdentifier string `json:"dbClusterIdentifier"`
}

func Migration(w http.ResponseWriter, _ *http.Request) {
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

	// データベースURLを作成
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		secretData.Username,
		secretData.Password,
		secretData.Host,
		secretData.Port,
		"story_fishing_db",
	)

	migrationsPath := "file:///app/infrastructure/database/migrations"
	log.Printf("Migrations path: %s", migrationsPath)

	m, err := migrate.New(
		migrationsPath,
		dbURL,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Migrate all the way up ...
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		http.Error(w, fmt.Sprintf("Failed to apply migrations: %v", err), http.StatusInternalServerError)
		return
	}

	log.Println("Migrations applied successfully")
	w.WriteHeader(http.StatusOK)
}
