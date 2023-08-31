package database

import (
	_itemData "fakhry/cleanarch/features/item/data"
	_userData "fakhry/cleanarch/features/user/data"

	"gorm.io/gorm"
)

// db migration
func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_userData.User{})
	db.AutoMigrate(&_itemData.Item{})
	/*
		TODO 2:
		migrate struct item
	*/
}
