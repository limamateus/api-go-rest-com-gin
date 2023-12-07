package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	Nome       string `json:"nome"`
	Email      string `json:"email"`
	Password   string `json:"senha"`
}
