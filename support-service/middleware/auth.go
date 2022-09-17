package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestURL := "http://localhost:8081/api/users/secured/ping"
		req, err := http.NewRequest(http.MethodGet, requestURL, nil)
		if err != nil {
			fmt.Printf("auth client: could not create request: %s\n", err)
			context.JSON(401, gin.H{"error": "auth client: could not create request: " + err.Error()})
			context.Abort()
			return
		}
		req.Header.Add("Authorization", context.GetHeader("Authorization"))

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("auth client: error making http request: %s\n", err)
			context.JSON(401, gin.H{"error": "auth client: error making http request: " + err.Error()})
			context.Abort()
			return
		}

		// fmt.Printf("client: got response!\n")
		// fmt.Printf("client: status code: %d\n", res.StatusCode)

		if res.StatusCode != http.StatusOK {
			context.JSON(401, gin.H{"error": "token is not authorized"})
			context.Abort()
			return
		}

		context.Next()
	}
}
