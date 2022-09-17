package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRole int

const (
	Administrator  UserRole = 0
	RegisteredUser UserRole = 1
	SupportUser    UserRole = 2
)

func (e UserRole) String() string {
	switch e {
	case Administrator:
		return "Administrator"
	case RegisteredUser:
		return "RegisteredUser"
	case SupportUser:
		return "SupportUser"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type User struct {
	gorm.Model
	Name     string   `json:"name" gorm:"not null"`
	Email    string   `json:"email" gorm:"unique;not-null"`
	Password string   `json:"password" gorm:"not null"`
	Role     UserRole `json:"userRole" gorm:"not null"`
	Blocked  bool     `json:"blocked" gorm:"default:false"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
