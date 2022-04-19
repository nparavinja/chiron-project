package db

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}
type ExaminationRepository struct {
	DB *gorm.DB
}

type Repository interface {
	Select(map[string]interface{}) (*User, error)
	Insert(map[string]interface{}) (*User, error)
	Update(map[string]interface{}) (*User, error)
	Delete(map[string]interface{}) (*User, error)
}

func (r *UserRepository) Select(user any, searchType string, params ...interface{}) (*User, error) {
	// or map[string]interface{}
	// queries here
	// user service
	switch searchType {
	case "register":
		// parameters := make
		// implement logi for username, email if exist
	case "login":

	default:

	}

	// var p Patient
	// db.Joins("User").Find(&p, 1)
	// result := db.Joins("User").First(&p, "username = ?", "oljasolja")

	// // Struct0
	// db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	// // SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

	// // Map
	// db.Where({"name": "jinzhu", "age": 20}).Find(&users)
	// // SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// fmt.Println(result.QueryFields)

	// db.First(&product, "code = ?", "D42") // find product with code D42

	// // Delete - delete product
	// db.Delete(&product, 1)

	return nil, nil
}
func (r *UserRepository) Insert(user any) error {
	switch user.(type) {
	case Patient:
		p, ok := user.(Patient)
		if !ok {
			return errors.New("Error while type into b")

		}
		tx := r.DB.Create(&p)
		if tx.Error != nil {
			return tx.Error
		}
	case Doctor:
		d, ok := user.(Doctor)
		if !ok {
			return errors.New("Error while type into b")
		}
		tx := r.DB.Create(&d)
		if tx.Error != nil {
			return tx.Error
		}
	default:
		return errors.New("Error while type into b")
	}

	return nil
}

func (r *UserRepository) Update(params ...interface{}) (*User, error) {
	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	return nil, nil
}
func (r *UserRepository) Delete(params ...interface{}) (*User, error) {

	return nil, nil
}
