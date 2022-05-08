package services

import (
	"github.com/nparavinja/chiron-project/back-legs/db"
	crypto "github.com/nparavinja/chiron-project/back-legs/encryption"
)

type DoctorService struct {
	UserRepository *db.UserRepository
}

type DoctorResponse struct {
	Success    bool     `json:"success"`
	DoctorData []string `json:"data"`
	Jwt        string   `json:"jwt,omitempty"`
}

func (DoctorService *DoctorService) Login(username string, password string) (any, error) {
	result, err := DoctorService.UserRepository.Select(db.Doctor{}, "login", username, password)
	if err != nil {
		// some error
		return nil, err
	}
	var response DoctorResponse
	doctor, ok := result.(db.Doctor)
	if !ok {
		return nil, err
	}
	response.Success = true
	response.DoctorData = append(response.DoctorData, doctor.Username, doctor.Name)
	// generate jwt
	response.Jwt = crypto.CreateJWT(username, doctor.ID.String())
	return response, nil
}
