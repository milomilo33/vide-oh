package models

import (
	"gorm.io/gorm"
)

// type UserRole int

// const (
// 	Administrator  UserRole = 0
// 	RegisteredUser UserRole = 1
// )

// func (e UserRole) String() string {
// 	switch e {
// 	case Administrator:
// 		return "Administrator"
// 	case RegisteredUser:
// 		return "RegisteredUser"
// 	default:
// 		return fmt.Sprintf("%d", int(e))
// 	}
// }

type Video struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null"`
	Filename    string `json:"filename" gorm:"unique;not-null"`
	Description string `json:"description" gorm:"not null"`
	OwnerEmail  string `json:"ownerEmail" gorm:"not null"`
	Reported    bool   `json:"reported" gorm:"default:false"`
}
