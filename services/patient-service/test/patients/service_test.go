package patients_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var _ = Describe("PatientService", Ordered, func() {
	var (
		ctx         context.Context
		pgContainer *postgres.PostgresContainer
		err         error
	)

	BeforeAll(func() {
		ctx = context.Background()

		pgContainer, err = postgres.RunContainer(ctx,
			testcontainers.WithImage("postgres:15.3-alpine"),
			postgres.WithDatabase("test-db"),
			postgres.WithUsername("postgres"),
			postgres.WithPassword("postgres"),
			testcontainers.WithWaitStrategy(
				wait.ForLog("database system is ready to accept connections").
					WithOccurrence(2).WithStartupTimeout(5*time.Second)),
		)

		if err = pgContainer.Terminate(ctx); err != nil {
		}
	})

	AfterAll(func() {
		pgContainer.Terminate(ctx)
	})

	It("fetches a list of patients", func() {
		Expect(true).To(BeTrue())
	})
})

// Context("When creating an invalid car", func() {

// It("should return the correct response", func() {
// 	    requestBody := `{"id": "701","title": "GM","color": "Transparent"}`
// 	    request, _ := http.NewRequest("POST", "http://localhost:8080/cars", bytes.NewBuffer([]byte(requestBody)))
// 	    request.Header.Set("Content-Type", "application/json")

// 	    client := &http.Client{}
// 	    response, err := client.Do(request)
// 	    Expect(err).NotTo(HaveOccurred())
// 	    Expect(response.StatusCode).To(Equal(http.StatusCreated))
// })

// 	It("should return an error in the response", func() {
// 		requestBody := ""
// 		request, _ := http.NewRequest("POST", "http://localhost:8080/cars", bytes.NewBuffer([]byte(requestBody)))
// 		request.Header.Set("Content-Type", "application/json")

// 		client := &http.Client{}
// 		response, err := client.Do(request)
// 		Expect(err).NotTo(HaveOccurred())
// 		Expect(response.StatusCode).To(Equal(http.StatusBadRequest))

// 		responseBody, _ := io.ReadAll(response.Body)
// 		Expect(string(responseBody)).To(ContainSubstring("Failed"))
// 	})
// })
