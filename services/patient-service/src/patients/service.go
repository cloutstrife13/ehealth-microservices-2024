package patients

import (
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type PatientService struct {
	Db *gorm.DB
}

var patients = []Patient{
	{Id: "1", FirstName: "Ada", LastName: "Lovelace", DateOfBirth: "10/12/1815"},
	{Id: "2", FirstName: "Lightning", LastName: "Quirks", DateOfBirth: "02/06/1955"},
}

func (service PatientService) getPatients(ctx iris.Context) {
	ctx.JSON(patients)
}
