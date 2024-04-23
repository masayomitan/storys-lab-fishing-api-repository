package action

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
    // レスポンスヘッダーのステータスコードを設定
    w.WriteHeader(http.StatusOK)
    
    // レスポンスヘッダーのコンテントタイプを設定
    w.Header().Set("Content-Type", "application/json")
    
    // レスポンスボディに文言を書き込み
    responseMessage := `{"status": "healthy", "message": "Service is up and running"}`
    _, err := w.Write([]byte(responseMessage))
    if err != nil {
        // エラーハンドリング
        fmt.Println("Error writing response:", err)
    }
}

