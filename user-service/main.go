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

	user2 := &models.User{
		Name:     "User Usersky",
		Email:    "user@user.com",
		Password: "123",
		Role:     models.RegisteredUser,
		Blocked:  false,
	}
	user2.HashPassword(user2.Password)
	database.Instance.Save(&user2)

	user3 := &models.User{
		Name:     "User2 Usersky2",
		Email:    "user2@user2.com",
		Password: "123",
		Role:     models.RegisteredUser,
		Blocked:  false,
	}
	user3.HashPassword(user3.Password)
	database.Instance.Save(&user3)

	// Initialize Router
	router := initRouter()
	router.Run(":8081")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/users")
	{
		api.POST("/login", controllers.Login)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/ping", controllers.Ping)
			secured.GET("/user/all-registered", controllers.GetAllRegisteredUsers) // only admin
			secured.GET("/block/:email", controllers.BlockUser)                    // only admin
			secured.GET("/user/:id", controllers.GetUserById)
			secured.GET("/user/current", controllers.GetCurrentUser)
			secured.GET("/user/change-name", controllers.ChangeName)
		}
	}
	return router
}
