package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Content    string         `json:"content" gorm:"not null"`
	OwnerEmail string         `json:"ownerEmail" gorm:"not null"`
	Date       datatypes.Date `json:"date" gorm:"not null"`
	SentByUser bool           `json:"sentByUser" gorm:"not null"`
}
