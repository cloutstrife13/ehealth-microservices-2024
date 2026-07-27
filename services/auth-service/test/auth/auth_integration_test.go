package auth_test

import (
	"context"
	"time"

	"github.com/cloutstrife13/ehealth-microservices-2024/packages/helpers"
	"github.com/cloutstrife13/ehealth-microservices-2024/packages/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	gormPg "gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/testcontainers/testcontainers-go"
	testcontainersPg "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"

	utils "github.com/cloutstrife13/ehealth-microservices-2024/packages/helpers"
	a "github.com/cloutstrife13/ehealth-microservices-2024/services/auth-service/src/auth"
)

var _ = Describe("AuthService (integration)", Ordered, func() {
	var (
		ctx         context.Context
		pgContainer *testcontainersPg.PostgresContainer
		authService *a.AuthService
	)

	globalUser := models.User{
		Name:        "Ada",
		Surname:     "Lovelace",
		DateOfBirth: helpers.ParseDate("10/12/1815"),
		Email:       "ada@lovelace.com",
	}

	BeforeAll(func() {
		ctx = context.Background()

		pgContainer, _ = testcontainersPg.RunContainer(ctx,
			testcontainers.WithImage("postgres:15.3-alpine"),
			testcontainersPg.WithDatabase("auth"),
			testcontainersPg.WithUsername("postgres"),
			testcontainersPg.WithPassword("postgres"),
			testcontainers.WithWaitStrategy(
				wait.ForLog("database system is ready to accept connections").
					WithOccurrence(2).WithStartupTimeout(5*time.Second)),
		)

		connectionString, _ := pgContainer.ConnectionString(ctx)

		db, _ := gorm.Open(gormPg.Open(connectionString), &gorm.Config{})

		err := db.AutoMigrate(
			&models.User{},
			&models.VerificationStatus{},
		)

		if err != nil {
			Fail("Failed to auto-migrate database schemas: " + err.Error())
		}

		authService = &a.AuthService{
			Db: db,
		}
	})

	AfterAll(func() {
		pgContainer.Terminate(ctx)
	})

	It("creates a new user as part of a registration process", func() {
		authService.Register(&globalUser)
		isUuidGeneratedAfterCreate := utils.IsValidUUID(globalUser.ID)
		Expect(isUuidGeneratedAfterCreate).To(BeTrue())
		Expect(globalUser.VerificationStatus.Status).To(Equal(models.StatusUnverified))
	})
})
