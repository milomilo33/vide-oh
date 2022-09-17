package main

import (
	"fmt"

	"support-service/database"
	"support-service/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	// Initialize Database
	database.Connect("postgresql://localhost/vide-oh-messages?user=postgres&password=root")
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(":8084")
}

func CORS() gin.HandlerFunc {
	// TO allow CORS
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORS())
	api := router.Group("/api/messages").Use(middleware.Auth())
	{
		api.GET("/video-stream/:name")
	}
	return router
}
