package patients

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type PatientService struct {
	Db *gorm.DB
}

var patients = []PatientDto{
	{ID: "1", FirstName: "Ada", LastName: "Lovelace", DateOfBirth: "10/12/1815"},
	{ID: "2", FirstName: "Lightning", LastName: "Quirks", DateOfBirth: "02/06/1955"},
}

func (service *PatientService) getPatients(ctx iris.Context) {
	ctx.JSON(patients)
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
