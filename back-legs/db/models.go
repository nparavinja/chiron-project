package db

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Doctor struct {
	gorm.Model
	Name        string
	JMBG        string
	Username    string
	Password    string
	Email       string
	IsAdmin     bool
	LicenseNo   string
	Examination []Examination
}

type Patient struct {
	gorm.Model
	Name         string
	JMBG         string
	Username     string
	Password     string
	Email        string
	IsAdmin      bool
	PIN          string
	Examinations []Examination
}

type Examination struct {
	gorm.Model
	Timestamp time.Time
	status    rune
	Report    Report
	PatientID uint
	DoctorID  uint
}

type Report struct {
	gorm.Model
	Name          string `gorm:"type:text"`
	Comment       string `gorm:"type:text"`
	Therapies     []Therapy
	Diagnosis     Diagnosis
	ExaminationID uint
}

type Diagnosis struct {
	gorm.Model
	Name     string
	Comment  string `gorm:"type:text"`
	ReportID int
}

type Therapy struct {
	gorm.Model
	Comment           string `gorm:"type:text"`
	AdditionalComment string `gorm:"type:text"`
	ReportID          int
}

func GetAllModels() []interface{} {
	// brb
	return []interface{}{&Patient{}, &Doctor{}}
}
