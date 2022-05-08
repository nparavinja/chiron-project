package db

import (
	"time"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Doctor struct {
	ID           uuid.UUID `json:"id" gorm:"type:varchar(36)"`
	Name         string    `json:"name" gorm:"type:varchar(50)"`
	Username     string    `json:"username" gorm:"type:varchar(50)"`
	Password     string    `json:"-"`
	Email        string    `json:"email" gorm:"type:varchar(50)"`
	LicenseNo    string    `json:"licenseNo" gorm:"type:varchar(10)"`
	Examinations []Examination
}

type Patient struct {
	ID           uuid.UUID     `json:"id" gorm:"type:varchar(36)"`
	Name         string        `json:"name" gorm:"type:varchar(50)"`
	JMBG         string        `json:"-" gorm:"type:varchar(13)"`
	Username     string        `json:"username"  gorm:"type:varchar(50)"`
	Password     string        `json:"-"`
	Email        string        `json:"email"  gorm:"type:varchar(50)"`
	PIN          string        `json:"-"  gorm:"type:varchar(10)"`
	Examinations []Examination `json:"examinations"`
}

type Examination struct {
	ID             uuid.UUID `json:"id" gorm:"type:varchar(36)"`
	TimestampStart time.Time `json:"tsStart"`
	TimestampEnd   time.Time `json:"tsEnd"`
	Status         rune      `json:"status"`
	Report         Report    `json:"report"`
	PatientID      uuid.UUID `json:"patient" gorm:"type:varchar(36)"` // fk
	DoctorID       uuid.UUID `json:"doctor" gorm:"type:varchar(36)"`  // fk
}

type Report struct {
	ID            uint      `json:"id"`
	Name          string    `gorm:"type:text" json:"name"`
	Therapy       Therapy   `json:"therapy"`
	Diagnosis     Diagnosis `json:"diagnosis"`
	ExaminationID uuid.UUID `json:"-"`
}

type Diagnosis struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Comment  string `gorm:"type:text" json:"comment"`
	ReportID uint   `json:"-"`
}

type Therapy struct {
	ID                uint   `json:"id"`
	Comment           string `gorm:"type:text" json:"comment"`
	AdditionalComment string `gorm:"type:text" json:"additionalComment"`
	ReportID          uint   `json:"-"`
}

func GetAllModels() []interface{} {
	// brb
	return []interface{}{&Patient{}, &Doctor{}}
}

func (d *Doctor) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.New()
	return nil
}

func (p *Patient) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return nil
}

func (e *Examination) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	return nil
}
