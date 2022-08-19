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

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./static")
	router.MaxMultipartMemory = 100 << 20
	api := router.Group("/api")
	{
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
