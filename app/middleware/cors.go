package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)


func SetupCORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://www.storys-lab-fishing.com/",
			"https://storys-lab-fishing.com/",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	})
}
