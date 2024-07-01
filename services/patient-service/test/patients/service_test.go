package patients_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	gormPg "gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/testcontainers/testcontainers-go"
	testcontainersPg "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	p "github.com/cloutstrife13/ehealth-microservices-2024/services/patient-service/src/patients"
	utils "github.com/cloutstrife13/ehealth-microservices-2024/services/patient-service/test/utils"
)

var _ = Describe("PatientService", Ordered, func() {
	var (
		ctx            context.Context
		pgContainer    *testcontainersPg.PostgresContainer
		patientService *p.PatientService
	)

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

	It("creates a new patient in the database", func() {
		patient := p.Patient{
			FirstName:   "Ada",
			LastName:    "Lovelace",
			DateOfBirth: "10/12/1815",
			InsuranceId: "123456789",
			InsurerId:   "987654321",
		}

		patientService.CreatePatient(&patient)

		isUuidGeneratedAfterCreate := utils.IsValidUUID(patient.ID.String())

		Expect(isUuidGeneratedAfterCreate).To(BeTrue())
	})
})
