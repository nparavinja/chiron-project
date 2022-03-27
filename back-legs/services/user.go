package services

import (
	"fmt"

	"github.com/nparavinja/chiron-project/back-legs/db"
)

type UserService struct {
	UserRepository *db.UserRepository
}

type UserActions interface {
	Login(data ...interface{}) (db.User, error)
	Register(data ...interface{}) (db.User, error)
}

func (UserService *UserService) Login(data map[string]interface{}) {
	fmt.Println("UserService Login called:")
	// data check here

	// if valid data, call userRepository.Select or smthing
}
func (UserService *UserService) Register(username string, password string, email string) (db.User, error) {
	fmt.Println("UserService Register called:")
	// data check here
	u := db.User{Username: username, Password: password, Email: email}
	fmt.Println(u)

	// if valid data, call userRepository.Select or smthing
	return u, nil

}
