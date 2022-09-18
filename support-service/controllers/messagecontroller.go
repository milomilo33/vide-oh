package controllers

import (
	"errors"
	"net/http"
	"support-service/database"
	"support-service/models"
	"support-service/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func AddMessage(msg string, email string, jwtClaims utils.JWTClaim) (message *models.Message, err error) {
	message = &models.Message{
		Content:    msg,
		OwnerEmail: email,
		SentByUser: jwtClaims.Role == "RegisteredUser",
		Date:       time.Now(),
	}
	record := database.Instance.Save(&message)

	if record.Error != nil {
		err = errors.New("DB error")
	}

	return message, err
}

func GetAllMessagesForUser(c *gin.Context) {
	email := c.Param("email")
	_, claims := utils.GetTokenClaims(c)
	if claims.Role != "SupportUser" && !(claims.Role == "RegisteredUser" && claims.Email == email) {
		c.JSON(401, gin.H{"error": "unauthorized role"})
		c.Abort()
		return
	}

	var messages []models.Message

	if err := database.Instance.Where("owner_email = ?", email).Find(&messages).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, messages)
}

func GetAllUserEmailsWithMessages(c *gin.Context) {
	_, claims := utils.GetTokenClaims(c)
	if claims.Role != "SupportUser" {
		c.JSON(401, gin.H{"error": "unauthorized role"})
		c.Abort()
		return
	}

	var userEmails []string

	if err := database.Instance.Model(&models.Message{}).Distinct("owner_email").Find(&userEmails).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, userEmails)
}
