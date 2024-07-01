package patients

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	FirstName   string
	LastName    string
	DateOfBirth string
	InsuranceId string
	InsurerId   string
}

type PatientDto struct {
	gorm.Model
	ID          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	InsuranceId string `json:"insuranceId"`
	InsurerId   string `json:"insurerId"`
}
