package patients

import "github.com/kataras/iris/v12"

var patients = []Patient{
	{ID: "1", FirstName: "Ada", LastName: "Lovelace", DateOfBirth: "10/12/1815"},
	{ID: "2", FirstName: "Lightning", LastName: "Quirks", DateOfBirth: "02/06/1955"},
}

func getPatients(ctx iris.Context) {
	ctx.JSON(patients)
}
