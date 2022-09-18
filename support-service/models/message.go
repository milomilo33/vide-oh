package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Content    string    `json:"content" gorm:"not null"`
	OwnerEmail string    `json:"ownerEmail" gorm:"not null"`
	Date       time.Time `json:"date" gorm:"not null"`
	SentByUser bool      `json:"sentByUser" gorm:"not null"`
}
