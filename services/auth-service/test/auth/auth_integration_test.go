package auth_test

import (
	"context"
	"fmt"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/cloutstrife13/ehealth-microservices-2024/packages/helpers"
	"github.com/cloutstrife13/ehealth-microservices-2024/packages/models"
	a "github.com/cloutstrife13/ehealth-microservices-2024/services/auth-service/src/auth"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	gormPg "gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/testcontainers/testcontainers-go"
	testcontainersPg "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var _ = Describe("AuthService (integration)", Ordered, func() {
	var (
		ctx               context.Context
		pgContainer       *testcontainersPg.PostgresContainer
		firebaseContainer testcontainers.Container
		authService       *a.AuthService
	)

	globalUser := models.User{
		Name:        "Ada",
		Surname:     "Lovelace",
		DateOfBirth: helpers.ParseDate("10/12/1815"),
		Email:       "ada@lovelace.com",
	}

	BeforeAll(func() {
		ctx = context.Background()

		// 1. Start PostgreSQL Container
		var err error
		pgContainer, err = testcontainersPg.RunContainer(ctx,
			testcontainers.WithImage("postgres:15.3-alpine"),
			testcontainersPg.WithDatabase("auth"),
			testcontainersPg.WithUsername("postgres"),
			testcontainersPg.WithPassword("postgres"),
			testcontainers.WithWaitStrategy(
				wait.ForLog("database system is ready to accept connections").
					WithOccurrence(2).WithStartupTimeout(10*time.Second)),
		)
		if err != nil {
			Fail("Failed to start Postgres container: " + err.Error())
		}

		connectionString, err := pgContainer.ConnectionString(ctx)
		if err != nil {
			Fail("Failed to get Postgres connection string: " + err.Error())
		}

		db, err := gorm.Open(gormPg.Open(connectionString), &gorm.Config{})
		if err != nil {
			Fail("Failed to connect to Postgres DB: " + err.Error())
		}

		// Run AutoMigrate in correct FK order
		err = db.AutoMigrate(
			&models.User{},
			&models.VerificationStatus{},
		)
		if err != nil {
			Fail("Failed to auto-migrate database schemas: " + err.Error())
		}

		// 2. Start Firebase Auth Emulator Container
		firebaseContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Image:        "andreysenov/firebase-tools:latest",
				ExposedPorts: []string{"9099/tcp"},
				Cmd: []string{
					"/bin/sh",
					"-c",
					`echo '{"emulators":{"auth":{"host":"0.0.0.0","port":9099}}}' > firebase.json && firebase emulators:start --only auth --project ehealth-local-project`,
				},
				WaitingFor: wait.ForListeningPort("9099/tcp").WithStartupTimeout(30 * time.Second),
			},
			Started: true,
		})
		if err != nil {
			Fail("Failed to start Firebase container: " + err.Error())
		}

		// Extract dynamic host & port mapped by Testcontainers
		fbHost, err := firebaseContainer.Host(ctx)
		if err != nil {
			Fail("Failed to get Firebase host: " + err.Error())
		}
		fbPort, err := firebaseContainer.MappedPort(ctx, "9099/tcp")
		if err != nil {
			Fail("Failed to get Firebase mapped port: " + err.Error())
		}

		// Point Firebase SDK to Testcontainers' mapped emulator port
		emulatorAddr := fmt.Sprintf("%s:%s", fbHost, fbPort.Port())
		os.Setenv("FIREBASE_AUTH_EMULATOR_HOST", emulatorAddr)

		// 3. Initialize Firebase Admin SDK Client
		app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: "ehealth-local-project"})
		if err != nil {
			Fail("Failed to initialize Firebase app: " + err.Error())
		}

		firebaseAuth, err := app.Auth(ctx)
		if err != nil {
			Fail("Failed to initialize Firebase Auth client: " + err.Error())
		}

		// 4. Construct AuthService with both dependencies wired
		authService = &a.AuthService{
			Db:           db,
			FirebaseAuth: firebaseAuth, // Adjust field name if named differently in service.go (e.g. AuthClient)
		}
	})

	AfterAll(func() {
		if pgContainer != nil {
			_ = pgContainer.Terminate(ctx)
		}
		if firebaseContainer != nil {
			_ = firebaseContainer.Terminate(ctx)
		}
	})

	It("creates a new user as part of a registration process", func() {
		err := authService.Register(&globalUser, "passwordTest@666")
		Expect(err).ToNot(HaveOccurred())
		Expect(globalUser.ID).ToNot(BeEmpty())
		Expect(len(globalUser.ID)).To(BeNumerically(">=", 20))
		Expect(globalUser.VerificationStatus.Status).To(Equal(models.StatusUnverified))
	})
})
