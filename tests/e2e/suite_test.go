package e2e

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSuiteE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "acceptance test suite")
}

var _ = BeforeSuite(func() {
	fmt.Println("Before Suite")
})
