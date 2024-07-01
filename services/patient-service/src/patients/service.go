package patients

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PatientService struct {
	Db *gorm.DB
}

func (service *PatientService) FindPatients(patients *[]Patient) {
	service.Db.Find(&patients)
}

func (service *PatientService) FindPatient(patient *Patient, ID uuid.UUID) {
	service.Db.First(&patient, ID)
}

func (service *PatientService) CreatePatient(patient *Patient) {
	service.Db.Create(&patient)
}
