package services

import (
	"github.com/nparavinja/chiron-project/back-legs/db"
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

	// generate password

	err := AdminService.UserRepository.Insert(db.Doctor{Name: name, Username: username, Email: email, LicenseNo: licenseNo, Password: "testpassword"})
	if err != nil {
		// some error
		return nil, err
	}

	// send password to mail
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
