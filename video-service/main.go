package main

import (
	"fmt"

	"video-service/controllers"
	"video-service/database"
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
	api := router.Group("/api")
	{
		api.GET("/video-stream/:name", controllers.StreamVideo)
	}
	return router
}
