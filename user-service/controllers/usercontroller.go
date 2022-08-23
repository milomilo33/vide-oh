package controllers

import (
	"net/http"
	"user-service/database"
	"user-service/models"
	"user-service/utils"

	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	user.Role = models.RegisteredUser
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email})
}

func BlockUser(context *gin.Context) {
	// RBAC
	_, claims := utils.GetTokenClaims(context)
	if claims.Role != models.Administrator.String() {
		context.JSON(401, gin.H{"error": "unauthorized role"})
		context.Abort()
		return
	}

	userEmail := context.Param("email")
	var user models.User

	if err := database.Instance.Where("email = ?", userEmail).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	user.Blocked = true

	database.Instance.Save(&user)

	utils.SendBlockedMail(user.Email)

	context.Status(http.StatusOK)
}

func GetAllRegisteredUsers(context *gin.Context) {
	_, claims := utils.GetTokenClaims(context)
	if claims.Role != models.Administrator.String() {
		context.JSON(401, gin.H{"error": "unauthorized role"})
		context.Abort()
		return
	}

	var users []models.User
	database.Instance.Where("role = ?", models.RegisteredUser).Find(&users)

	context.JSON(http.StatusOK, users)
}

func GetUserById(context *gin.Context) {
	// // RBAC
	// _, claims := utils.GetTokenClaims(context)
	// if claims.Role != models.Administrator.String() {
	// 	context.JSON(401, gin.H{"error": "unauthorized role"})
	// 	context.Abort()
	// 	return
	// }

	userId := context.Param("id")
	var user models.User

	if err := database.Instance.First(&user, userId).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, user)
}

func GetCurrentUser(context *gin.Context) {
	_, claims := utils.GetTokenClaims(context)

	var user models.User
	database.Instance.Where("email = ?", claims.Email).First(&user)

	context.JSON(http.StatusOK, user)
}

func ChangeName(context *gin.Context) {
	_, claims := utils.GetTokenClaims(context)

	var user models.User
	if err := database.Instance.Where("email = ?", claims.Email).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	name := context.Query("name")
	if strings.TrimSpace(name) == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "white space string"})
		context.Abort()
		return
	}

	user.Name = name

	database.Instance.Save(&user)

	context.Status(http.StatusOK)
}
