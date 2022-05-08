package db

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(pass string) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("root:%s@tcp(chiron-db:3306)/chiron?charset=utf8mb4&parseTime=True&loc=Local", pass)
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
	err := db.AutoMigrate(Doctor{}, Patient{}, Examination{}, Report{}, Therapy{}, Diagnosis{})
	if err != nil {
		log.Println("Error in auto migration")
		log.Println(err.Error())
		return false
	}
	return true
}
