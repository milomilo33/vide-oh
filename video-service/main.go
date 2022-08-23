package main

import (
	"fmt"

	"video-service/controllers"
	"video-service/database"
	"video-service/middleware"
	"video-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	// Initialize Database
	database.Connect("postgresql://localhost/vide-oh-videos?user=postgres&password=root")
	database.Migrate()

	// Initial Data
	video := &models.Video{
		Filename:    "someuniquefilename",
		OwnerEmail:  "user@user.com",
		Title:       "user's example video",
		Description: "you'll see nothing special here",
	}
	database.Instance.Save(&video)

	// Initialize Router
	router := initRouter()
	router.Run(":8082")
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
	// router.Static("/static", "./static")
	router.MaxMultipartMemory = 100 << 20
	api := router.Group("/api/videos")
	{
		api.Static("/static", "./static")
		api.GET("/video-stream/:name", controllers.StreamVideo)
		api.GET("/report-video/:id", controllers.ReportVideo)
		api.GET("/search-videos", controllers.SearchVideos)
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/ping")
			secured.GET("/all-reported-videos", controllers.GetAllReportedVideos)
			secured.POST("/upload-video", controllers.UploadVideo)
			secured.GET("/delete-video/:id", controllers.DeleteVideo)
		}
	}
	return router
}
