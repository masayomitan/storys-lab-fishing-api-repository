package main

import (
    "fmt"
	"os"
	"time"

    // "github.com/gin-gonic/gin"
    "storys-lab-fishing-api/app/infrastructure"
    // "storys-lab-fishing-api/app/infrastructure/log"
    "storys-lab-fishing-api/app/infrastructure/router"
    // "storys-lab-fishing-api/app/infrastructure/database"

)

func main() {
    fmt.Println("build start...")

    var app = infrastructure.NewConfig().
        Name(os.Getenv("APP_NAME")).
        ContextTimeout(10 * time.Second).
        // Logger(log.InstanceLogrusLogger).
        // Validator(validation.InstanceGoPlayground).
        DbSQL()

    app.WebServerPort(os.Getenv("APP_PORT")).
        WebServer(router.InstanceGin).
        Start()
}
