package main

import (
	"fmt"
	"os"
	"time"
	"github.com/joho/godotenv"

	// "github.com/gin-gonic/gin"
	"storys-lab-fishing-api/app/infrastructure"
	"storys-lab-fishing-api/app/infrastructure/log"
	"storys-lab-fishing-api/app/infrastructure/router"
	// "storys-lab-fishing-api/app/infrastructure/database"
)

func main() {
	fmt.Println("build start....")
	err := godotenv.Load((".env"))
	if err != nil {
		fmt.Println("envファイルの読み込みに失敗しました")
	}
	var app = infrastructure.NewConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		Logger(log.InstanceLogrusLogger).
		// Validator(validation.InstanceGoPlayground).
		DbSQL()

    fmt.Println(app)
	app.WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceGin).
		Start()
}
