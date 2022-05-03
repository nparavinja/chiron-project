package db

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	Name         string
	JMBG         string
	Username     string
	Password     string
	Email        string
	IsAdmin      bool
	LicenseNo    string
	Examinations []Examination
}

type Patient struct {
	gorm.Model   `json:"-"`
	Name         string        `json:"name"`
	JMBG         string        `json:"-"`
	Username     string        `json:"username"`
	Password     string        `json:"-"`
	Email        string        `json:"email"`
	IsAdmin      bool          `json:"-"`
	PIN          string        `json:"-"`
	Examinations []Examination `json:"examinations"`
}

type Examination struct {
	gorm.Model     `json:"-"`
	TimestampStart time.Time `json:"tsStart"`
	TimestampEnd   time.Time `json:"tsEnd"`
	Status         rune      `json:"status"`
	Report         Report    `json:"report"`
	PatientID      uint      `json:"patient"`
	DoctorID       uint      `json:"doctor"`
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
