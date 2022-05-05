package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nparavinja/chiron-project/back-legs/db"
)

type ExaminationService struct {
	ExaminationRepository *db.ExaminationRepository
}

type ExaminationResponse struct {
	Success bool  `json:"success"`
	Data    []any `json:"data,omitempty"`
}

const (
	Pending = iota
	Approved
	InProgress
	Done
)

type Result struct {
}

func (ExaminationService *ExaminationService) SetupAppointment(patientID string, doctorID string, timestampStart string, timestampEnd string) (any, error) {
	patientUUID, err := uuid.Parse(patientID)
	if err != nil {
		return nil, err
	}
	doctorUUID, err := uuid.Parse(doctorID)
	if err != nil {
		return nil, err
	}
	// convert string to time.Time format
	// Parse the date string into Go's time object
	// The 1st param specifies the format,
	// 2nd is our date string
	st, err := time.Parse("2006:01:02 15:04", timestampStart)
	if err != nil {
		// panic(err)
		fmt.Println(err)
		return nil, err
	}
	et, err := time.Parse("2006:01:02 15:04", timestampEnd)
	if err != nil {
		// panic(err)
		fmt.Println(err)
		return nil, err
	}
	// logic for date comparing -> older, do they exist in db, check timestamps

	e := db.Examination{PatientID: patientUUID, DoctorID: doctorUUID, TimestampStart: st, TimestampEnd: et, Status: Pending}
	err = ExaminationService.ExaminationRepository.Insert(e)
	if err != nil {
		// some error
		return nil, err
	}
	var response ExaminationResponse
	response.Success = true
	return response, nil
}

func (ExaminationService *ExaminationService) GetExaminations(patientID string) (any, error) {
	result, err := ExaminationService.ExaminationRepository.Select("all-p", patientID)
	if err != nil {
		// some error
		return nil, err
	}
	var response ExaminationResponse
	var patient = result.(db.Patient)
	// format response here - perhaps a helper function
	response.Success = true
	response.Data = append(response.Data, patient.Examinations)

	return response, nil
}
