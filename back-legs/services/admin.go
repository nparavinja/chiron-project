package services

import (
	"errors"
	"fmt"

	"github.com/nparavinja/chiron-project/back-legs/db"
	crypto "github.com/nparavinja/chiron-project/back-legs/encryption"
)

type AdminService struct {
	UserRepository *db.UserRepository
}

type AdminResponse struct {
	Success bool   `json:"success"`
	Data    []any  `json:"data"`
	Jwt     string `json:"jwt,omitempty"`
}

func (AdminService *AdminService) AddDoctor(name string, username string, email string, licenseNo string) (any, error) {
	// check for same username or email
	found, err := AdminService.UserRepository.Select(db.Doctor{}, "register", username, email)
	if err != nil {
		return nil, err
	}
	if found.(bool) {
		return nil, errors.New("Username/email already exist.")
	}
	// generate password
	firstPassword, firstPasswordEncrypted, err := crypto.GenerateRandomPassword()
	fmt.Println(firstPassword, firstPasswordEncrypted)

	err = AdminService.UserRepository.Insert(db.Doctor{Name: name, Username: username, Email: email, LicenseNo: licenseNo, Password: firstPasswordEncrypted})
	if err != nil {
		// some error
		return nil, err
	}

	// send firstPassword to mail

	return "Doctor added successfully.", nil
}
func (AdminService *AdminService) GetDoctors() (any, error) { // data check here
	result, err := AdminService.UserRepository.Select(db.Doctor{}, "all")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (AdminService *AdminService) UpdateDoctor(uuid string, params ...any) (any, error) {
	// result, err := AdminService.UserRepository.Update(username, password)
	// if err != nil {
	// 	// some error
	// 	return nil, err
	// }
	// var response PatientResponse
	// patient, ok := result.(db.Patient)
	// if !ok {
	// 	return nil, err
	// }
	// response.Success = true
	// response.PatientData = append(response.PatientData, patient.Username, patient.Email, patient.JMBG, patient.PIN)
	// // generate jwt
	// response.Jwt = crypto.CreateJWT(username)

	return nil, nil
}

func (AdminService *AdminService) RemoveDoctor(uuid string) (any, error) {
	err := AdminService.UserRepository.Delete(db.Doctor{}, uuid)
	if err != nil {
		// some error
		return nil, err
	}
	var response AdminResponse
	response.Success = true
	return response, nil
}
