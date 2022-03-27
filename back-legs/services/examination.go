package services

import "github.com/nparavinja/chiron-project/back-legs/db"

type ExaminationService struct {
	ExaminationRepository *db.ExaminationRepository
}
