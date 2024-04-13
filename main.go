package main

import (

    "github.com/gin-gonic/gin"
    "storys-lab-fishing-api/database"
    "fmt"
)

func main() {
    fmt.Println("start")
    r := gin.Default()
    db := database.GormConnect()
    fmt.Println(db)

    // ルートパス
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "hello world",
        })
    })

    // ヘルスチェック用のパス
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "ok",
        })
    })
    r.Run(":80")
}
