package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID         int64
	RoleID     int64
	Email      string
	Password   string
	Name       string
	LastAccess time.Time
}
