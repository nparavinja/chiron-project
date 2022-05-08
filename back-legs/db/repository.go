package db

import (
	"errors"

	crypto "github.com/nparavinja/chiron-project/back-legs/encryption"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}
type ExaminationRepository struct {
	DB *gorm.DB
}

type Repository interface {
	Select(map[string]interface{}) (any, error)
	Insert(map[string]interface{}) (any, error)
	Update(map[string]interface{}) (any, error)
	Delete(map[string]interface{}) (any, error)
}

func (r *UserRepository) Select(userType any, searchType string, data ...any) (any, error) {
	// or map[string]interface{}
	// queries here
	// user service

	switch userType.(type) {
	case Doctor:
		switch searchType {
		case "all":
			var doctors []Doctor
			tx := r.DB.Find(&doctors)
			if tx.Error != nil {
				return nil, tx.Error
			}
			return doctors, nil
		case "register":
			var d Doctor
			found := true
			if result := r.DB.First(&d, "username = ? OR email = ?", data[0], data[1]); result.Error != nil {
				if errors.Is(result.Error, gorm.ErrRecordNotFound) {
					found = false
					return found, nil
				}
				return nil, result.Error
			}
			return found, nil
		case "not-added":
			var p Patient
			// get hashed password from db
			result := r.DB.First(&p, "username = ?", data[0])
			if result.Error != nil {
				return nil, result.Error
			}
			if crypto.Compare(data[1].(string), p.Password) {
				return p, nil
			}
			return nil, errors.New("Login error.")
		default:
		}

	case Patient:
		switch searchType {
		case "register":
			var p Patient
			found := true
			if result := r.DB.First(&p, "username = ? OR email = ?", data[0], data[1]); result.Error != nil {
				if errors.Is(result.Error, gorm.ErrRecordNotFound) {
					found = false
					return found, nil
				}
				return nil, result.Error
			}
			return found, nil
		case "login":
			var p Patient
			// get hashed password from db
			result := r.DB.First(&p, "username = ?", data[0])
			if result.Error != nil {
				return nil, result.Error
			}
			if crypto.Compare(data[1].(string), p.Password) {
				return p, nil
			}
			return nil, errors.New("Login error.")
		default:
		}
	default:

	}

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

func (r *UserRepository) Update(params ...interface{}) (any, error) {
	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	return nil, nil
}
func (r *UserRepository) Delete(user any, data ...interface{}) error {
	switch user.(type) {
	case Patient:
		var p Patient
		tx := r.DB.Delete(&p, data[0])
		if tx.Error != nil {
			return tx.Error
		}
	case Doctor:
		var d Doctor
		tx := r.DB.Where("id = ?", data[0]).Delete(&d)
		// check if something is really deleted
		if tx.Error != nil {
			return tx.Error
		}
		if tx.RowsAffected == 0 {
			return errors.New("No user found for ID.")
		}
	default:
		return errors.New("Error while type asserting")
	}
	return nil
}

func (e *ExaminationRepository) Select(searchType string, data ...any) (any, error) {
	switch searchType {
	case "all-p":
		var patient Patient
		// implement logic for one
		tx := e.DB.Preload("Examinations").Preload("Examinations.Report").Preload("Examinations.Report.Therapy").Preload("Examinations.Report.Diagnosis").Find(&patient, "id = ?", data[0])
		if tx.Error != nil {
			return nil, tx.Error
		}
		return patient, nil
	default:
		return nil, errors.New("unexpectedError")

	}
}

func (r *ExaminationRepository) Insert(ex Examination) error {
	tx := r.DB.Create(&ex)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *ExaminationRepository) Update(ex Examination) error {
	tx := r.DB.Model(&ex).Updates(ex)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
