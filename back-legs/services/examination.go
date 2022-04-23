package services

import (
	"github.com/nparavinja/chiron-project/back-legs/db"
)

type ExaminationService struct {
	ExaminationRepository *db.ExaminationRepository
}

func (ExaminationService *ExaminationService) SetupAppointment(username string) (any, error) {
	result, err := ExaminationService.ExaminationRepository.Select("", username)
	// check active examinations
	// check if slot is
	if err != nil {
		// some error
		return nil, err
	}

	return result, nil
}

func (ExaminationService *ExaminationService) GetExaminations(username string) (any, error) {
	result, err := ExaminationService.ExaminationRepository.Select("", username)
	if err != nil {
		// some error
		return nil, err
	}

	return result, nil
}
