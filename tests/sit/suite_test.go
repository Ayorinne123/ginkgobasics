package sit

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSuiteSIT(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "acceptance test suite")
}

var _ = BeforeSuite(func() {
	fmt.Println("BeforeSuite")
})

var _ = AfterSuite(func() {
	fmt.Println("AfterSuite")
})
