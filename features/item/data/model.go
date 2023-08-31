package data

import (
	_userData "fakhry/cleanarch/features/user/data"

	"gorm.io/gorm"
)

// struct item gorm model
type Item struct {
	gorm.Model
	Name        string
	UserID      uint
	Brand       string
	Description string
	Price       int
	Weight      int
	User        _userData.User
}
