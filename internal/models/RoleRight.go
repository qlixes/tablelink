package models

import "gorm.io/gorm"

type RoleRight struct {
	gorm.Model

	ID      int64
	RoleID  int64
	Route   string
	Section string
	Path    string
	RCreate bool
	RRead   bool
	RUpdate bool
	RDelete bool
}
