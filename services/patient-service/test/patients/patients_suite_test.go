package patients_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPatients(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Patients Suite")
}
