package database

import (
	"homework/app/models"

	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Class{},
		&models.Student{},
	)

	return err
}
