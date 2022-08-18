package main

import (
	"fmt"
	"user-service/controllers"
	"user-service/database"
	"user-service/middleware"
	"user-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")

	// Initialize Database
	database.Connect("postgresql://localhost/vide-oh-users?user=postgres&password=root")
	database.Migrate()

	// Initial Data
	user := &models.User{
		Name:     "Admin Adminsky",
		Email:    "admin@admin.com",
		Password: "123",
		Role:     models.Administrator,
		Blocked:  false,
	}
	user.HashPassword(user.Password)
	database.Instance.Save(&user)

	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/login", controllers.Login)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/ping", controllers.Ping)
			secured.GET("/user/all-registered")
			secured.GET("/block/:id", controllers.BlockUser)
		}
	}
	return router
}
