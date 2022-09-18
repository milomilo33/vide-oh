package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"support-service/controllers"
	"support-service/database"
	"support-service/middleware"
	"support-service/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
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

type SocketMessage struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func initRouter() *gin.Engine {
	router := gin.Default()
	m := melody.New()
	router.Use(CORS())
	api := router.Group("/api/messages")
	{
		api.GET("/:email/ws", func(c *gin.Context) {
			token, _ := c.GetQuery("token")
			if !middleware.AuthenticateToken(token) {
				c.JSON(401, gin.H{"error": "unauthorized"})
				c.Abort()
				return
			}

			email := c.Param("email")
			_, claims := utils.GetTokenClaimsFromTokenString(token)
			if claims.Role != "SupportUser" && !(claims.Role == "RegisteredUser" && claims.Email == email) {
				c.JSON(401, gin.H{"error": "unauthorized role"})
				c.Abort()
				return
			}

			m.HandleRequest(c.Writer, c.Request)
		})

		m.HandleMessage(func(s *melody.Session, msg []byte) {
			splitPath := strings.Split(s.Request.URL.Path, "/")
			email := splitPath[len(splitPath)-2]
			var socketMessage SocketMessage
			err := json.Unmarshal(msg, &socketMessage)
			if err != nil {
				fmt.Println("error: ", err)
				return
			}
			fmt.Println(socketMessage)
			_, claims := utils.GetTokenClaimsFromTokenString(socketMessage.Token)

			message, err := controllers.AddMessage(socketMessage.Message, email, claims)
			if err != nil {
				fmt.Println("error: ", err)
				return
			}

			messageJSON, err := json.Marshal(message)
			if err != nil {
				fmt.Println("error: ", err)
				return
			}

			m.BroadcastFilter(messageJSON, func(q *melody.Session) bool {
				return q.Request.URL.Path == s.Request.URL.Path
			})
		})

		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/:email/all", controllers.GetAllMessagesForUser)
			secured.GET("/user-emails", controllers.GetAllUserEmailsWithMessages)
		}
	}
	return router
}
