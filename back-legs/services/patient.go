package services

import (
	"errors"

	"github.com/nparavinja/chiron-project/back-legs/db"
	crypto "github.com/nparavinja/chiron-project/back-legs/encryption"
)

type PatientService struct {
	UserRepository *db.UserRepository
}

type UserActions interface {
	Login(data ...interface{}) (db.User, error)
	Register(data ...interface{}) (db.User, error)
}

func (PatientService *PatientService) Login(username string, password string) (any, error) {
	result, err := PatientService.UserRepository.Select("login", username, password)
	if err != nil {
		// some error
		return nil, err
	}
	response := result.(db.Patient)
	return response, nil
}
func (PatientService *PatientService) Register(name string, username string, password string, email string, jmbg string) (*db.Patient, error) { // data check here
	var p db.Patient
	found, err := PatientService.UserRepository.Select("register", username, email)
	if err != nil {
		return nil, err
	}
	if !found.(bool) {
		pw, err := crypto.EncryptText(password)
		if err != nil {
			return nil, errors.New("error")
		}
		p = db.Patient{User: db.User{Username: username, Password: pw, Email: email, Name: name, JMBG: jmbg}, PIN: "12345"}
		PatientService.UserRepository.Insert(p)
		return &p, nil

	} else {
		return nil, errors.New("User already exists.")
	}
}
