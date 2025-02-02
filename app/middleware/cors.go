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

			// 一般公開側
			"https://www.storys-lab-fishing.com",
			"https://storys-lab-fishing.com",

			// 管理側
			"https://admin-fishing.storys-lab.com",
			"https://www.admin-fishing.storys-lab.com",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	})
}
