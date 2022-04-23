package services

import (
	"fmt"

	"github.com/nparavinja/chiron-project/back-legs/db"
)

type DoctorService struct {
	UserRepository *db.UserRepository
}

func (DoctorService *DoctorService) Login(username string, password string) {
	fmt.Println("DoctorService Login called:")
	// data check here

	// if valid data, call userRepository.Select or smthing
}
func (DoctorService *DoctorService) Register(name string, username string, password string, email string, jmbg string) (db.Patient, error) { // data check here
	p := db.Patient{Username: username, Password: password, Email: email, Name: name, JMBG: jmbg}
	fmt.Println(p)

	// if valid data, call userRepository.Select or smthing
	return p, nil

}
