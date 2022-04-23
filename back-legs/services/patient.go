package services

import (
	"errors"

	"github.com/nparavinja/chiron-project/back-legs/db"
	crypto "github.com/nparavinja/chiron-project/back-legs/encryption"
)

type PatientService struct {
	UserRepository *db.UserRepository
}

type PatientResponse struct {
	Success     bool     `json:"success"`
	PatientData []string `json:"data"`
	Jwt         string   `json:"jwt,omitempty"`
}

func (PatientService *PatientService) Login(username string, password string) (any, error) {
	result, err := PatientService.UserRepository.Select("login", username, password)
	if err != nil {
		// some error
		return nil, err
	}
	var response PatientResponse
	patient, ok := result.(db.Patient)
	if !ok {
		return nil, err
	}
	response.Success = true
	response.PatientData = append(response.PatientData, patient.Username, patient.Email, patient.JMBG, patient.PIN)
	// generate jwt
	response.Jwt = crypto.CreateJWT(username)

	return response, nil
}
func (PatientService *PatientService) Register(name string, username string, password string, email string, jmbg string, pin string) (any, error) { // data check here
	var p db.Patient
	found, err := PatientService.UserRepository.Select("register", username, email)
	if err != nil {
		return nil, err
	}
	if !found.(bool) {
		pw, err := crypto.EncryptText(password)
		if err != nil {
			return nil, errors.New("unexpectedError")
		}
		p = db.Patient{Username: username, Password: pw, Email: email, Name: name, JMBG: jmbg, PIN: pin}
		PatientService.UserRepository.Insert(p)
		var response PatientResponse
		response.Success = true
		response.PatientData = append(response.PatientData, p.Username, p.Email)
		return response, nil
	} else {
		return nil, errors.New("User already exists.")
	}
}

func (PatientService *PatientService) Update(username string, password string) (any, error) {
	result, err := PatientService.UserRepository.Update(username, password)
	if err != nil {
		// some error
		return nil, err
	}
	var response PatientResponse
	patient, ok := result.(db.Patient)
	if !ok {
		return nil, err
	}
	response.Success = true
	response.PatientData = append(response.PatientData, patient.Username, patient.Email, patient.JMBG, patient.PIN)
	// generate jwt
	response.Jwt = crypto.CreateJWT(username)

	return response, nil
}
