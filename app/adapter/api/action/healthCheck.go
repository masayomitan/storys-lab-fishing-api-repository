package action

import (
	"fmt"
	"net/http"

    "github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
    // レスポンスヘッダーのステータスコードを設定
    c.Writer.WriteHeader(http.StatusOK)
    
    // レスポンスヘッダーのコンテントタイプを設定
    c.Writer.Header().Set("Content-Type", "application/json")
    
    // レスポンスボディに文言を書き込み
    responseMessage := `{"status": "healthy", "message": "Service is up and running"}`
    _, err := c.Writer.Write([]byte(responseMessage))
    if err != nil {
        // エラーハンドリング
        fmt.Println("Error writing response:", err)
    }
}

