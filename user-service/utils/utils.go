package utils

import (
	"user-service/auth"

	"github.com/gin-gonic/gin"
)

func GetTokenClaims(context *gin.Context) (err error, jwtClaims auth.JWTClaim) {
	tokenString := context.GetHeader("Authorization")
	err, jwtClaims = auth.ValidateToken(tokenString)
	return
}
