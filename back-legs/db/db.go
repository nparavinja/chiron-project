package db

import (
	"errors"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(connectionString string) (*gorm.DB, error) {
	gormDb, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database")
		return nil, err
	}
	migrateErr := Migrate(gormDb)
	if !migrateErr {
		log.Println("Failed to migrate models")
		return nil, errors.New("Failed to migrate models")
	}
	return gormDb, nil
}

func Migrate(db *gorm.DB) bool {
	err := db.AutoMigrate(Doctor{}, User{}, Patient{})
	if err != nil {
		log.Println("Error in auto migration")
		log.Println(err.Error())
		return false
	}
	return true
}
