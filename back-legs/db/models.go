package db

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
	Name     string
	JMBG     string
	Username string
	Password string
	Email    string
	IsAdmin  bool
}

type Doctor struct {
	User      User
	UserID    int
	LicenseNo string
}

type Patient struct {
	User   User
	UserID int
	PIN    string
}

func GetAllModels() []interface{} {
	// brb
	return []interface{}{&User{}, &Patient{}, &Doctor{}}
}
