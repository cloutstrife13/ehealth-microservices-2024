package patients_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"

	gormPg "gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/testcontainers/testcontainers-go"
	testcontainersPg "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	p "github.com/cloutstrife13/ehealth-microservices-2024/services/patient-service/src/patients"
	utils "github.com/cloutstrife13/ehealth-microservices-2024/services/patient-service/test/utils"
)

var _ = Describe("PatientService (integration)", Ordered, func() {
	var (
		ctx            context.Context
		pgContainer    *testcontainersPg.PostgresContainer
		patientService *p.PatientService
	)

	globalPatient := p.Patient{
		FirstName:   "Ada",
		LastName:    "Lovelace",
		DateOfBirth: "10/12/1815",
		InsuranceId: "123456789",
		InsurerId:   "987654321",
	}

	BeforeAll(func() {
		ctx = context.Background()

		pgContainer, _ = testcontainersPg.RunContainer(ctx,
			testcontainers.WithImage("postgres:15.3-alpine"),
			testcontainersPg.WithDatabase("patients"),
			testcontainersPg.WithUsername("postgres"),
			testcontainersPg.WithPassword("postgres"),
			testcontainers.WithWaitStrategy(
				wait.ForLog("database system is ready to accept connections").
					WithOccurrence(2).WithStartupTimeout(5*time.Second)),
		)

		connectionString, _ := pgContainer.ConnectionString(ctx)

		db, _ := gorm.Open(gormPg.Open(connectionString), &gorm.Config{})

		db.AutoMigrate(&p.Patient{})

		patientService = &p.PatientService{
			Db: db,
		}
	})

	AfterAll(func() {
		pgContainer.Terminate(ctx)
	})

	It("creates a new patient in the DB", func() {
		patientService.CreatePatient(&globalPatient)

		isUuidGeneratedAfterCreate := utils.IsValidUUID(globalPatient.ID.String())

		Expect(isUuidGeneratedAfterCreate).To(BeTrue())
	})

	It("reads patients from the DB", func() {
		patients := []p.Patient{}

		patientService.FindPatients(&patients)

		Expect(patients).To(HaveLen(1))
	})

	It("reads a patient from the DB", func() {
		patient := p.Patient{}

		patientService.FindPatient(&patient, globalPatient.ID)

		Expect(patient).To(MatchFields(IgnoreExtras, Fields{
			"FirstName":   Equal("Ada"),
			"LastName":    Equal("Lovelace"),
			"DateOfBirth": Equal("10/12/1815"),
			"InsuranceId": Equal("123456789"),
			"InsurerId":   Equal("987654321"),
		}))
	})
})
