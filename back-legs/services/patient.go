package services

import (
	"fmt"

	"github.com/nparavinja/chiron-project/back-legs/db"
)

type PatientService struct {
	UserRepository *db.UserRepository
}

type UserActions interface {
	Login(data ...interface{}) (db.User, error)
	Register(data ...interface{}) (db.User, error)
}

func (PatientService *PatientService) Login(username string, password string) {
	fmt.Println("PatientService Login called:")
	// data check here

	// if valid data, call userRepository.Select or smthing
}
func (PatientService *PatientService) Register(name string, username string, password string, email string, jmbg string) (db.Patient, error) { // data check here
	p := db.Patient{User: db.User{Username: username, Password: password, Email: email, Name: name, JMBG: jmbg}, PIN: "12345"}
	fmt.Println(p)

	PatientService.UserRepository.Insert(p)
	// if valid data, call userRepository.Select or smthing
	return p, nil

}
